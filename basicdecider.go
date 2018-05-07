package diamonds

import (
	"github.com/sunspots/diamonds/vector2"
)

// BasicDecider implements a very simple decider to run your bot
type BasicDecider struct {
	botID      string
	lastTarget vector2.Vector2
}

// NewBasicDecider returns a new instance of BasicDecider
func NewBasicDecider(botID string) *BasicDecider {
	return &BasicDecider{botID, vector2.Zero()}
}

// Next calculates the next move direction by finding the closest diamond
func (d *BasicDecider) Next(board *Board, lastResult int) (direction vector2.Vector2) {

	// do something smart in case of conflict
	if lastResult == StatusConflict {

	}

	self := board.GetBot(d.botID)

	if self.Diamonds > 4 {
		println("D: Go home, you're drunk!")
		return DirectionTo(self.Position, self.Base)
	}

	target := board.ClosestDiamond(self.Position)

	return DirectionTo(self.Position, target)
}

// DirectionTo returns the direction from a position to another
// this implementation will always direct you in a wide right angle
// following the X axis first, then the Y axis
func DirectionTo(from, to vector2.Vector2) vector2.Vector2 {
	if from.X < to.X {
		return vector2.East()
	}
	if from.X > to.X {
		return vector2.West()
	}
	if from.Y < to.Y {
		return vector2.South()
	}
	return vector2.North()
}
