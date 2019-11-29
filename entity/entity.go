package entity

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Entity type that is used to update and render entities in game
type Entity interface {
	Update(w *pixelgl.Window)
	Draw(t *pixelgl.Canvas)
}

// DefaultEntity normal entity to extend for Entities
type DefaultEntity struct {
	matrix pixel.Matrix
	sprite *pixel.Sprite
}
