package diamonds

import (
	"github.com/sunspots/diamonds/vector2"
)

// The Decider interface is used to choose the next direction vector for your bot
type Decider interface {
	Next(board *Board, status int) vector2.Vector2
}

// Status values for MoveResults
//
const (
	StatusConflict  = 409
	StatusError     = 500
	StatusOK        = 200
	StatusForbidden = 403
)
