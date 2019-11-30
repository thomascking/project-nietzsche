package interfaces

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Collision contains collision data
type Collision struct {
	Vec pixel.Vec
	Ent Entity
}

// World type is an interface to DefaultWorld
type World interface {
	AddEntity(e Entity)
	// removeEntity()
	Update(w *pixelgl.Window, dt float64)
	Draw()
	Render(c *pixelgl.Canvas)
	CheckCollision(e Entity, v pixel.Vec) []Collision
}

// Entity type that is used to update and render entities in game
type Entity interface {
	Update(w *pixelgl.Window, dt float64)
	Draw(t *pixelgl.Canvas)
	SetWorld(w World)
	GetType() string
}

// Graphic renderable graphic
type Graphic interface {
	Draw(target pixel.Target, m pixel.Matrix)
}
