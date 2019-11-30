package world

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/interfaces"
	"golang.org/x/image/colornames"
)

// DefaultWorld the base world struct
type DefaultWorld struct {
	Canvas *pixelgl.Canvas
	Ents   []interfaces.Entity
}

// Update for Base-Type DefaultWorld
func (dw *DefaultWorld) Update(w *pixelgl.Window, dt float64) {
	for _, e := range dw.Ents {
		e.Update(w, dt)
	}
}

// Draw for Base-Type DefaultWorld
func (dw *DefaultWorld) Draw() {
	dw.Canvas.Clear(colornames.White)
	for _, e := range dw.Ents {
		e.Draw(dw.Canvas)
	}
}

//Render renderes canvas to window
func (dw *DefaultWorld) Render(c *pixelgl.Canvas) {
	c.Clear(colornames.Black)
	c.SetMatrix(pixel.IM.Scaled(pixel.ZV,
		math.Min(
			c.Bounds().W()/dw.Canvas.Bounds().W(),
			c.Bounds().H()/dw.Canvas.Bounds().H(),
		),
	).Moved(c.Bounds().Center()))
	dw.Canvas.Draw(c, pixel.IM)
}

//AddEntity adds entity to world's entity list
func (dw *DefaultWorld) AddEntity(e interfaces.Entity) {
	dw.Ents = append(dw.Ents, e)
	e.SetWorld(dw)
}

// CheckCollision checks for collision at given vec
func (dw *DefaultWorld) CheckCollision(e interfaces.Entity, v pixel.Vec) []interfaces.Collision {
	// loop through entities
	// check intersection of bounds of the entities
	// change collisionCanvas to bounds of the intersection
	// clear to crazy color (magenta?)
	// draw e and other entity
	// check every pixel for not magenta
	return nil
}

// NewWorld initalizes the world
func NewWorld(b pixel.Rect) *DefaultWorld {
	w := &DefaultWorld{
		Canvas: pixelgl.NewCanvas(b),
		Ents:   []interfaces.Entity{},
	} //Dynamically Allocate a world
	return w
}
