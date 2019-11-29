package world

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/entity"
)

//Game is a world
type Game struct {
	DefaultWorld
}

//Update Game
func (g *Game) Update(w *pixelgl.Window) {
	g.DefaultWorld.Update(w)
}

//Draw Game
func (g *Game) Draw(t *pixelgl.Canvas) {
	g.DefaultWorld.Draw()
}

//AddEntity adds an entity to the menu
func (g *Game) AddEntity(e entity.Entity) {
	g.DefaultWorld.AddEntity(e)
}
