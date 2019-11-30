package world

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/entity"
	"github.com/thomascking/project-nietzsche/state"
)

//Pause is a world
type Pause struct {
	DefaultWorld
}

//NewPauseWorld creates a new pause world
func NewPauseWorld(b pixel.Rect) *Pause {
	dw := NewWorld(b)
	p := &Pause{*dw}
	p.AddEntity(entity.NewButton("./images/resume.png",
		pixel.R(0, 130, 128, 258),
		func() {
			state.CurrState = state.GS
		}),
	)
	p.AddEntity(entity.NewButton("./images/exit.png",
		pixel.R(0, 0, 128, 128),
		func() {
			state.CurrState = state.ES
		}),
	)
	return p
}

//Update Pause
func (p *Pause) Update(w *pixelgl.Window, dt float64) {
	p.DefaultWorld.Update(w, dt)
}

//Draw Pause
func (p *Pause) Draw() {
	p.DefaultWorld.Draw()
}

//AddEntity adds an entity to the menu
func (p *Pause) AddEntity(e entity.Entity) {
	p.DefaultWorld.AddEntity(e)
}
