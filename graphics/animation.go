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
	frameDur     float64
	time         float64
}

// NewAnimation create a new animation
func NewAnimation(path string, bounds pixel.Rect) *Animation {
	sprite, err := LoadSprite(path, bounds)
	if err != nil {
		panic(err)
	}
	a := &Animation{
		sprite:       sprite,
		animations:   make(map[string][]pixel.Rect),
		current:      "",
		currentIndex: 0,
		frameDur:     .1,
	}
	a.animations[""] = []pixel.Rect{bounds}
	return a
}

// Draw draws the animation
func (a *Animation) Draw(t pixel.Target, m pixel.Matrix) {
	a.sprite.Draw(t, m)
}

// Update update the animation
func (a *Animation) Update(dt float64) {
	// TODO: go to next frame
	a.time += dt
	if a.time > a.frameDur {
		a.time -= a.frameDur
		a.currentIndex++
		if a.currentIndex >= len(a.animations[a.current]) {
			a.currentIndex = 0
		}
		a.sprite.Set(a.sprite.Picture(), a.animations[a.current][a.currentIndex])
	}
}

// AddAnimation add an animation
func (a *Animation) AddAnimation(name string, frames []pixel.Rect) {
	a.animations[name] = frames
}

// PlayAnimation start playing an animation
func (a *Animation) PlayAnimation(name string) {
	if name != a.current {
		a.currentIndex = 0
		a.current = name
	}
}
