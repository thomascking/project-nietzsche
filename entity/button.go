package entity

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/graphics"
)

// Button its a button
type Button struct {
	pressed bool
	s       string
	g       *pixel.Sprite
	b       pixel.Rect
	frames  map[string]pixel.Rect
	f       func()
}

// NewButton creates a button
func NewButton(filename string, bounds pixel.Rect, f func()) *Button {
	modifiedPos := bounds.Moved(pixel.V(-bounds.Min.X, -bounds.Min.Y))
	sprite, err := graphics.LoadSprite(filename, modifiedPos)
	if err != nil {
		panic(err)
	}
	b := &Button{
		g:       sprite,
		f:       f,
		b:       bounds,
		frames:  make(map[string]pixel.Rect),
		pressed: false,
	}
	b.frames[""] = modifiedPos.Moved(pixel.V(0, 128))
	b.frames["p"] = modifiedPos
	b.g.Set(b.g.Picture(), b.frames[""])
	return b
}

//Draw draws the button on the screen
func (b *Button) Draw(t *pixelgl.Canvas) {
	b.g.Draw(t, pixel.IM.Moved(b.b.Center()))
}

//Update updates the button
func (b *Button) Update(w *pixelgl.Window, _ float64) {
	if w.JustPressed(pixelgl.MouseButtonLeft) {
		pos := w.MousePosition()
		if b.b.Contains(pos) {
			b.pressed = true
			b.g.Set(b.g.Picture(), b.frames["p"])
		}
	}
	if w.JustReleased(pixelgl.MouseButtonLeft) && b.pressed {
		b.g.Set(b.g.Picture(), b.frames[""])
		pos := w.MousePosition()
		if b.b.Contains(pos) {
			b.f()
		}
		b.pressed = false
	}
}
