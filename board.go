package diamonds

import (
	"github.com/sunspots/diamonds/vector2"
)

// Board contains all the information about the current game state.
type Board struct {
	ID                       string            `json:"id,omitempty"`
	Width                    int               `json:"width,omitempty"`
	Height                   int               `json:"height,omitempty"`
	Bots                     []Bot             `json:"bots,omitempty"`
	GameObjects              []GameObject      `json:"gameobjects,omitempty"`
	Diamonds                 []vector2.Vector2 `json:"diamonds,omitempty"`
	MinimumDelayBetweenMoves int               `json:"minimumDelayBetweenMoves,omitempty"`
}

// OnBoard checks if a position is within the boundaries of the board
func (board *Board) OnBoard(pos vector2.Vector2) bool {
	return pos.X >= 0 && pos.X < board.Width-1 && pos.Y >= 0 && pos.Y < board.Width-1
}

// ClosestDiamond gets the closest diamond using Manhattan distance
func (board *Board) ClosestDiamond(from vector2.Vector2) vector2.Vector2 {
	if len(board.Diamonds) == 0 {
		return vector2.Zero()
	}
	closest := board.Diamonds[0]
	for _, diamond := range board.Diamonds {
		if diamond.Distance(from) < closest.Distance(from) {
			closest = diamond
			break
		}
	}
	return closest
}

// IsOccupied indicates if there's a bot in the position
func (board *Board) IsOccupied(pos vector2.Vector2) bool {
	for _, bot := range board.Bots {
		if bot.Position == pos {
			return true
		}
	}
	return false
}

// IsPortal indicates if there's a portal in the position
func (board *Board) IsPortal(pos vector2.Vector2) bool {
	for _, obj := range board.GameObjects {
		if obj.Position == pos {
			return true
		}
	}
	return false
}

// IsDiamond indicates if there's a diamond in the position
func (board *Board) IsDiamond(pos vector2.Vector2) bool {
	for _, diamond := range board.Diamonds {
		if diamond == pos {
			return true
		}
	}
	return false
}

// PortalTo returns the target position of the portal
func (board *Board) PortalTo(from vector2.Vector2) (to vector2.Vector2) {
	portal := GameObject{}

	for _, obj := range board.GameObjects {
		if obj.Position == from {
			portal = obj
		}
	}

	for _, obj := range board.GameObjects {
		if obj.LinkedTeleporterString == portal.LinkedTeleporterString {
			return obj.Position
		}
	}
	return from
}

// HasBot indicates if the board has a bot with the given ID
func (board *Board) HasBot(botID string) bool {
	for _, bot := range board.Bots {
		if bot.BotID == botID {
			return true
		}
	}
	return false
}

// GetBot returns a bot with the given name, or the zero-value-bot (not useful)
func (board *Board) GetBot(botID string) Bot {
	for _, bot := range board.Bots {
		if bot.BotID == botID {
			return bot
		}
	}
	return Bot{}
}
