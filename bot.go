package diamonds

import (
	"time"

	"github.com/sunspots/diamonds/vector2"
)

// Bot represents a bot on the game board
type Bot struct {
	Name                string          `json:"name,omitempty"`
	Base                vector2.Vector2 `json:"base,omitempty"`
	Position            vector2.Vector2 `json:"position,omitempty"`
	Diamonds            int             `json:"diamonds,omitempty"`
	TimeJoined          time.Time       `json:"timeJoined,omitempty"`
	MillisecondsLeft    int             `json:"millisecondsLeft,omitempty"`
	Score               int             `json:"score,omitempty"`
	BotID               string          `json:"botId,omitempty"`
	NextMoveAvailableAt time.Time       `json:"nextMoveAvailableAt,omitempty"`
}
