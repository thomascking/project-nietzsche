package world

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/thomascking/project-nietzsche/entity"
	"github.com/thomascking/project-nietzsche/state"
	"github.com/thomascking/project-nietzsche/ttscreen"
)

//Pause is a world
type Pause struct {
	DefaultWorld
}

//NewPauseWorld creates a new pause world
func NewPauseWorld(b pixel.Rect) *Pause {
	dw := NewWorld(b)
	p := &Pause{*dw}
	txt := text.New(b.Center(), ttscreen.Atlas)
	s := "Resume"
	r := txt.BoundsOf(s)
	p.AddEntity(entity.NewButton(s,
		r.Resized(r.Center(), pixel.V(r.W()+13, r.H()+13)),
		func() {
			state.CurrState = state.GS
		}),
	)
	s = "EXIT"
	r = txt.BoundsOf(s).Moved(pixel.V(0, r.H()+20))
	p.AddEntity(entity.NewButton(s,
		r.Resized(r.Center(), pixel.V(r.W()+13, r.H()+13)),
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
