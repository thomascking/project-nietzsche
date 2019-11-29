package world

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/entity"
	"github.com/thomascking/project-nietzsche/state"
)

//Game is a world
type Game struct {
	DefaultWorld
}

// NewGameWorld creates a game world
func NewGameWorld(b pixel.Rect) *Game {
	dw := NewWorld(b)
	g := &Game{*dw}
	g.AddEntity(entity.NewPlayer(pixel.V(100, 100)))
	return g
}

//Update Game
func (g *Game) Update(w *pixelgl.Window, dt float64) {
	g.DefaultWorld.Update(w, dt)
	if w.JustPressed(pixelgl.KeyEscape) {
		state.CurrState = state.PS
	}
}

//Draw Game
func (g *Game) Draw() {
	g.DefaultWorld.Draw()
}

//AddEntity adds an entity to the menu
func (g *Game) AddEntity(e entity.Entity) {
	g.DefaultWorld.AddEntity(e)
}
