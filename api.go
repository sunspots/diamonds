package diamonds

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/sunspots/diamonds/vector2"
)

const defaultURL = "http://diamonds.etimo.se/api/"

// API manages the requests to the server to
type API struct {
	token       string
	baseURL     string
	client      *http.Client
	boardNumber int
}

// NewAPI returns a new API instance
func NewAPI(token string) *API {
	tr := &http.Transport{
		MaxIdleConnsPerHost: 1024,
		TLSHandshakeTimeout: 0 * time.Second,
	}
	return &API{token, defaultURL, &http.Client{Transport: tr}, 1}
}

// Join the game for another round!
func (api *API) Join() error {
	data := bytes.NewBuffer([]byte{})
	fmt.Fprintf(data, `{"botToken": "%s"}`, api.token)

	res, err := api.client.Post(
		fmt.Sprintf("%sBoards/%d/join", api.baseURL, api.boardNumber),
		"application/json", data)
	if res.StatusCode == 409 {
		println("Already joined game!")
	} else if err != nil || res.StatusCode != 200 {
		println("join error ", err.Error(), res.Status)
	}
	res.Body.Close()
	return err
}

// GetBoard fetches and decodes a board from the server
func (api *API) GetBoard() (board Board, err error) {
	res, err := api.client.Get(fmt.Sprintf("%sBoards/%d", api.baseURL, api.boardNumber))
	if err != nil {
		return board, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&board)
	return board, err
}

// Move sends a request to move to the server
func (api *API) Move(dir vector2.Vector2) (status int, board Board, err error) {
	var direction string

	switch dir {
	case vector2.North():
		direction = "North"
	case vector2.East():
		direction = "East"
	case vector2.South():
		direction = "South"
	case vector2.West():
		direction = "West"
	default:
		direction = "South"
	}

	data := bytes.NewBuffer([]byte{})
	fmt.Fprintf(
		data,
		`{"botToken": "%s", "direction": "%s" }`,
		api.token,
		direction,
	)

	res, err := api.client.Post(fmt.Sprintf("%sBoards/%d/move", api.baseURL, api.boardNumber), "application/json", data)
	status = res.StatusCode

	if err != nil {
		return status, board, err
	}

	defer res.Body.Close()
	if res.StatusCode == StatusError {
		return status, board, errors.New(res.Status)
	}

	json.NewDecoder(res.Body).Decode(board)

	return status, board, err
}
