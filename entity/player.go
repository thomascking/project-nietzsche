package entity

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Player the player character
type Player struct {
	DefaultEntity
}

// NewPlayer create a new player
func NewPlayer(location pixel.Matrix) *Player {
	sprite, err := loadSprite("./images/meme (1).png", pixel.R(0, 0, 400, 300))
	if err != nil {
		panic(err)
	}
	return &Player{
		DefaultEntity{
			sprite: sprite,
			matrix: location}}
}

// Update update the player each loop
func (p *Player) Update(win *pixelgl.Window) {
}

// Draw the player each loop
func (p *Player) Draw(win *pixelgl.Window) {
	p.sprite.Draw(win, p.matrix)
}
