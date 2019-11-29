package world

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/entity"
	"github.com/thomascking/project-nietzsche/state"
)

//Menu is a world
type Menu struct {
	DefaultWorld
}

//NewMenuWorld creates a new menu world
func NewMenuWorld(b pixel.Rect) *Menu {
	buttonRect := pixel.R(0, 0, 128, 128)
	dw := NewWorld(b)
	m := &Menu{*dw}
	m.AddEntity(entity.NewButton(b.Center(), "BEGIN!", buttonRect, func() {
		state.CurrState = state.GS
	}))
	return m
}

//Update Menu
func (m *Menu) Update(w *pixelgl.Window) {
	m.DefaultWorld.Update(w)
}

//Draw Menu
func (m *Menu) Draw() {
	m.DefaultWorld.Draw()
}

//AddEntity adds an entity to the menu
func (m *Menu) AddEntity(e entity.Entity) {
	m.DefaultWorld.AddEntity(e)
}
