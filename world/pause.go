package world

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/entity"
)

//Pause is a world
type Pause struct {
	DefaultWorld
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
