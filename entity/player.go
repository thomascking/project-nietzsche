package entity

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"github.com/thomascking/project-nietzsche/graphics"
)

// Player the player character
type Player struct {
	animation *graphics.Animation
	position  pixel.Vec
}

// NewPlayer create a new player
func NewPlayer(p pixel.Vec) *Player {
	a := graphics.NewAnimation("./images/walking.png")
	return &Player{a, p}
}

// Update update the player each loop
func (p *Player) Update(w *pixelgl.Window) {
}

// Draw the player each loop
func (p *Player) Draw(win *pixelgl.Window) {
	p.animation.Draw(win, pixel.IM.Moved(p.position))
}
