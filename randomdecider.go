package diamonds

import (
	"math/rand"

	"github.com/sunspots/diamonds/vector2"
)

// RandomDecider is a super-simple decider that simply returns a random direction
type RandomDecider struct {
}

// Next returns the next random direction
func (*RandomDecider) Next(board *Board, _ int) vector2.Vector2 {
	switch rand.Int() % 4 {
	case 0:
		return vector2.North()
	case 1:
		return vector2.East()
	case 2:
		return vector2.South()
	default:
		return vector2.West()
	}
}
