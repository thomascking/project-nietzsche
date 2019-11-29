package entity

import (
	"github.com/faiface/pixel/pixelgl"
)

// Entity type that is used to update and render entities in game
type Entity interface {
	Update(win *pixelgl.Window)
	Draw(win *pixelgl.Window)
}
