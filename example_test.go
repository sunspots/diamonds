package diamonds_test

import (
	"time"

	"github.com/sunspots/diamonds"
)

const botid = "foo"
const token = "bar"

// ExampleBot won't work without a proper bot Identity
func ExampleBot() {
	println("Starting bot...")
	diamonds.NewClient(
		// An identity is required for authentication
		// and finding your own bot on the board
		diamonds.Identity{
			ID:    botid,
			Name:  "Vidar",
			Email: "",
			Token: token,
		},
		// The decider will decide where your bot will go
		diamonds.NewBasicDecider(botid),
		// The delayModifier will reduce your wait time,
		// it should be around 2*RTT + overhead
		time.Millisecond*100,
	).Run()
}
