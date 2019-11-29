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
func (p *Pause) Update(w *pixelgl.Window) {
	p.DefaultWorld.Update(w)
}

//Draw Pause
func (p *Pause) Draw(t *pixelgl.Canvas) {
	p.DefaultWorld.Draw()
}

//AddEntity adds an entity to the menu
func (p *Pause) AddEntity(e entity.Entity) {
	p.DefaultWorld.AddEntity(e)
}
