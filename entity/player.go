package entity

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	"github.com/thomascking/project-nietzsche/graphics"
)

const speed = 45

// Player the player character
type Player struct {
	animation *graphics.Animation
	position  pixel.Vec
	direction pixel.Vec
}

// NewPlayer create a new player
func NewPlayer(p pixel.Vec) *Player {
	a := graphics.NewAnimation("./images/walking.png", pixel.R(0, 64, 32, 96))
	a.AddAnimation("walk", []pixel.Rect{
		pixel.R(32, 64, 64, 96),
		pixel.R(64, 64, 96, 96),
		pixel.R(0, 32, 32, 64),
		pixel.R(32, 32, 64, 64),
		pixel.R(64, 32, 96, 64),
		pixel.R(0, 0, 32, 32),
		pixel.R(32, 0, 64, 32),
		pixel.R(64, 0, 96, 32),
	})
	a.AddAnimation("idle", []pixel.Rect{
		pixel.R(0, 64, 32, 96),
	})
	return &Player{a, p, pixel.V(1, 1)}
}

// Update update the player each loop
func (p *Player) Update(w *pixelgl.Window, dt float64) {
	p.animation.Update(dt)
	if w.Pressed(pixelgl.KeyRight) {
		p.animation.PlayAnimation("walk")
		p.direction = pixel.V(1, 1)
		p.position = p.position.Add(pixel.V(dt*speed, 0))
	} else if w.Pressed(pixelgl.KeyLeft) {
		p.animation.PlayAnimation("walk")
		p.direction = pixel.V(-1, 1)
		p.position = p.position.Add(pixel.V(-dt*speed, 0))
	} else {
		p.animation.PlayAnimation("idle")
	}
}

// Draw the player each loop
func (p *Player) Draw(win *pixelgl.Canvas) {
	p.animation.Draw(win, pixel.IM.Moved(p.position).ScaledXY(p.position, p.direction))
}
