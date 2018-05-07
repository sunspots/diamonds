package diamonds

import "github.com/sunspots/diamonds/vector2"

// GameObject represents objects in the world, really just teleporters...
type GameObject struct {
	Name                   string          `json:"name,omitempty"`
	Position               vector2.Vector2 `json:"position,omitempty"`
	IsBlocking             bool            `json:"isBlocking,omitempty"`
	LinkedTeleporterString string          `json:"linkedTeleporterString,omitempty"`
}
