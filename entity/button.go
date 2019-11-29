package entity

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/state"
	"github.com/thomascking/project-nietzsche/ttscreen"
	"golang.org/x/image/colornames"
)

// Button its a button
type Button struct {
	s string
	v pixel.Vec
	g *pixel.Sprite
	b *pixel.Rect
	f func()
}

// NewButton creates a button
func NewButton(location pixel.Vec, text string, bounds pixel.Rect, f func()) *Button {
	sprite, err := loadSprite("./images/black.png", bounds)
	if err != nil {
		panic(err)
	}
	return &Button{
		s: text,
		v: location,
		g: sprite,
		f: f,
	}
}

//Update updates the button
func (b *Button) Update(w *pixelgl.Window) {
	if w.JustPressed(pixelgl.MouseButtonLeft) {
		b.f()
		fmt.Println(state.CurrState)
	}
}

//Draw draws the button on the screen
func (b *Button) Draw(t *pixelgl.Canvas) {
	b.g.Draw(t, pixel.IM.Moved(b.v))
	ttscreen.DrawText(t, []string{b.s}, b.v, colornames.White)
}
