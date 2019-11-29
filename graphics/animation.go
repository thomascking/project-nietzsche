package graphics

import (
	"github.com/faiface/pixel"
)

// Animation an animation
type Animation struct {
	sprite       *pixel.Sprite
	animations   map[string][]pixel.Rect
	current      string
	currentIndex int
}

// NewAnimation create a new animation
func NewAnimation(path string, bounds pixel.Rect) *Animation {
	sprite, err := LoadSprite(path, bounds)
	if err != nil {
		panic(err)
	}
	return &Animation{
		sprite:       sprite,
		animations:   make(map[string][]pixel.Rect),
		current:      "",
		currentIndex: 0,
	}
}

// Draw draws the animation
func (a *Animation) Draw(t pixel.Target, m pixel.Matrix) {
	if a.current != "" {
		a.sprite.Set(a.sprite.Picture(), a.animations[a.current][a.currentIndex])
	}
	a.sprite.Draw(t, m)
}

// Update update the animation
func (a *Animation) Update() {
	// TODO: go to next frame
}

// AddAnimation add an animation
func (a *Animation) AddAnimation(name string, frames []pixel.Rect) {
	a.animations[name] = frames
}

// PlayAnimation start playing an animation
func (a *Animation) PlayAnimation(name string) {
	a.currentIndex = 0
	a.current = name
}
