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
	dw := NewWorld(b)
	m := &Menu{*dw}
	m.AddEntity(entity.NewButton("./images/begin.png",
		pixel.R(0, 130, 128, 258),
		func() {
			state.CurrState = state.GS
		}),
	)
	m.AddEntity(entity.NewButton("./images/exit.png",
		pixel.R(0, 0, 128, 128),
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
