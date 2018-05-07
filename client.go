package diamonds

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Client handles everything the bot needs to do to talk to the server
type Client struct {
	api  *API
	self Identity

	// decider is responsible for deciding what the next move is
	decider Decider
	// scores keeps a record of the scores from all previous rounds
	scores []int
	// The delayModifier should be something like 2*RTT + overhead
	// try keeping it as high as possible without getting too many drops
	delayModifier time.Duration

	// flag that tells the client to stop after the current round
	lastRound bool
}

// NewClient returns a new client object that handles one bots' game loop
func NewClient(id Identity, decider Decider, delayModifier time.Duration) *Client {
	return &Client{NewAPI(id.Token), id, decider, []int{}, delayModifier, false}
}

// Run runs the bot continually
func (client *Client) Run() {

	client.scores = []int{}

	lastScore := -1

	lastStatus := StatusOK
	drops := 0
	moveAttempts := 0
	startTime := time.Now()

	client.lastRound = false

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for {
			<-c
			// If we already set lastRound, force exit on second ^C
			if client.lastRound {
				client.printScores(drops, moveAttempts, startTime)
				os.Exit(1)
			}
			// Allow waiting for round to finish
			client.lastRound = true
			fmt.Println("\nWaiting for round completion ^C to force close.")
		}
	}()

	// Get an initial state fo the board to see if we're already playing
	b, _ := client.api.GetBoard()

	// Our bot is already on the board
	if b.HasBot(client.self.ID) {
		println("Resuming round")
	}

	disconnected := true
	for {
		board, err := client.api.GetBoard()
		if err != nil {
			println("Failed to get board")
		}

		if disconnected {
			if client.lastRound {
				return
			}
			if board.HasBot(client.self.ID) {
				disconnected = false
				continue
			}

			if err := client.api.Join(); err != nil {
				println("Join failed", err.Error())
			} else {
				println("Joined.")
				drops = 0
				moveAttempts = 0
				startTime = time.Now()
				disconnected = false
				board, err = client.api.GetBoard()
				if err != nil {
					println("Failed to get board")
				}
			}
		}
		lastStatus, _, err = client.api.Move(client.decider.Next(&board, lastStatus))
		moveAttempts++

		if lastStatus == StatusError {
			println("Server Error!")
			if err := client.api.Join(); err != nil {
				println("Join failed", err.Error())
			}
		}

		// Move was forbidden!
		if lastStatus == StatusForbidden {
			self := board.GetBot(client.self.ID)
			score := self.Score
			if score == 0 && lastScore == 0 && self.Name == "" {
				disconnected = true
				continue
			}
			// round reset?
			if score < lastScore || !board.HasBot(client.self.ID) && score != 0 {
				if score < lastScore {
					score = lastScore
				}

				client.scores = append(client.scores, score)

				println("Round completed:", score)
				client.printScores(drops, moveAttempts, startTime)

				lastScore = 0
				disconnected = true
				time.Sleep(time.Duration(self.MillisecondsLeft+10) * time.Millisecond * 2)
				continue
			}
			drops++
			continue
		}

		if lastStatus != StatusOK {
			if err != nil {
				println("Move err", err.Error())
			}
		} else {
			lastScore = board.GetBot(client.self.ID).Score
			delayBetween := time.Duration(board.MinimumDelayBetweenMoves) * time.Millisecond
			time.Sleep(delayBetween - client.delayModifier)
		}
	}
}

func (client *Client) printScores(drops, moveAttempts int, startTime time.Time) {
	dropRate := float64(drops) / float64(moveAttempts)
	totalTime := time.Now().Sub(startTime)
	successes := moveAttempts - drops
	spSec := float64(successes) / float64(totalTime.Seconds())

	fmt.Printf("Moves: %d, Drops %d, DropRate %.3f%%, Moves/sec %.3f\n", successes, drops, dropRate*100, spSec)
	fmt.Printf("Max: %d Avg: %d over %d rounds\n", max(client.scores), avg(client.scores), len(client.scores))
}
