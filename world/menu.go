package world

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/entity"
	"github.com/thomascking/project-nietzsche/state"
	"github.com/thomascking/project-nietzsche/ttscreen"
	"golang.org/x/image/colornames"
)

//Menu is a world
type Menu struct {
	DefaultWorld
}

//NewMenuWorld creates a new menu world
func NewMenuWorld(b pixel.Rect) *Menu {
	dw := NewWorld(b)
	m := &Menu{*dw}
	centerVec := b.Center()
	m.AddEntity(entity.NewButton("./images/begin.png",
		pixel.R(centerVec.X-(128/2), centerVec.Y, centerVec.X+(128/2), centerVec.Y+128),
		func() {
			state.CurrState = state.GS
		}),
	)
	m.AddEntity(entity.NewButton("./images/exit.png",
		pixel.R(centerVec.X-(128/2), centerVec.Y-64, centerVec.X+(128/2), centerVec.Y+64),
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
	centerLine := m.Canvas.Bounds()
	ttscreen.DrawText(m.Canvas, []string{"A Game for All and None", "made by", "Thomas King and Taylor Lopez"}, pixel.V(centerLine.Center().X, centerLine.H()-128), colornames.Black, 1)
}

//AddEntity adds an entity to the menu
func (m *Menu) AddEntity(e entity.Entity) {
	m.DefaultWorld.AddEntity(e)
}
