package entity

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/graphics"
	"github.com/thomascking/project-nietzsche/ttscreen"
	"golang.org/x/image/colornames"
)

// Button its a button
type Button struct {
	pressed bool
	s       string
	g       *pixel.Sprite
	b       pixel.Rect
	f       func()
}

// NewButton creates a button
func NewButton(text string, bounds pixel.Rect, f func()) *Button {
	sprite, err := graphics.LoadSprite("./images/black.png", bounds.Moved(pixel.V(-bounds.Min.X, -bounds.Min.Y)))
	if err != nil {
		panic(err)
	}
	return &Button{
		s:       text,
		g:       sprite,
		f:       f,
		b:       bounds,
		pressed: false,
	}
}

//Update updates the button
func (b *Button) Update(w *pixelgl.Window, _ float64) {
	if w.JustPressed(pixelgl.MouseButtonLeft) {
		pos := w.MousePosition()
		if b.b.Contains(pos) {
			b.pressed = true
		}
	}
	if w.JustReleased(pixelgl.MouseButtonLeft) && b.pressed {
		pos := w.MousePosition()
		if b.b.Contains(pos) {
			b.f()
		}
		b.pressed = false
	}
}

//Draw draws the button on the screen
func (b *Button) Draw(t *pixelgl.Canvas) {
	b.g.Draw(t, pixel.IM.Moved(b.b.Center()))
	ttscreen.DrawText(t, []string{b.s}, b.b.Center(), colornames.Red, 2)
}
