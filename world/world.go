package world

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/entity"
	"golang.org/x/image/colornames"
)

// World type is an interface to DefaultWorld
type World interface {
	AddEntity(e entity.Entity)
	// removeEntity()
	Update(w *pixelgl.Window)
	Draw()
	Render(c *pixelgl.Canvas)
}

// DefaultWorld the base world struct
type DefaultWorld struct {
	Canvas *pixelgl.Canvas
	Ents   []entity.Entity
}

// Update for Base-Type DefaultWorld
func (dw *DefaultWorld) Update(w *pixelgl.Window) {
	for _, e := range dw.Ents {
		e.Update(w)
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
	dw.Canvas.Draw(c, pixel.IM.Moved(c.Bounds().Center()))
}

//AddEntity adds entity to world's entity list
func (dw *DefaultWorld) AddEntity(e entity.Entity) {
	dw.Ents = append(dw.Ents, e)
}

// NewWorld initalizes the world
func NewWorld(b pixel.Rect) *DefaultWorld {
	w := &DefaultWorld{
		Canvas: pixelgl.NewCanvas(b),
		Ents:   []entity.Entity{},
	} //Dynamically Allocate a world
	return w
}
