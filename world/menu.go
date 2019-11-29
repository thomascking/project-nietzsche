package world

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/thomascking/project-nietzsche/entity"
	"github.com/thomascking/project-nietzsche/state"
	"github.com/thomascking/project-nietzsche/ttscreen"
)

//Menu is a world
type Menu struct {
	DefaultWorld
}

//NewMenuWorld creates a new menu world
func NewMenuWorld(b pixel.Rect) *Menu {
	dw := NewWorld(b)
	m := &Menu{*dw}
	txt := text.New(b.Center(), ttscreen.Atlas)
	s := "BEGIN!"
	r := txt.BoundsOf(s)
	m.AddEntity(entity.NewButton(s,
		r.Resized(r.Center(), pixel.V(r.W()+13, r.H()+13)),
		func() {
			state.CurrState = state.GS
		}),
	)
	s = "EXIT"
	r = txt.BoundsOf(s).Moved(pixel.V(0, r.H()+20))
	m.AddEntity(entity.NewButton(s,
		r.Resized(r.Center(), pixel.V(r.W()+13, r.H()+13)),
		func() {
			state.CurrState = state.ES
		}),
	)
	return m
}

//Update Menu
func (m *Menu) Update(w *pixelgl.Window, dt float64) {
	m.DefaultWorld.Update(w, dt)
}

//Draw Menu
func (m *Menu) Draw() {
	m.DefaultWorld.Draw()
}

//AddEntity adds an entity to the menu
func (m *Menu) AddEntity(e entity.Entity) {
	m.DefaultWorld.AddEntity(e)
}
