package maps

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/graphics"
	"github.com/thomascking/project-nietzsche/world"
)

// Map is a map
type Map func(w world.World)

// Maps the available maps
var Maps = [1]Map{NewTestMap}

type fgEntity struct {
	sprite *pixel.Sprite
}

func (fg *fgEntity) Update(w *pixelgl.Window, dt float64) {}
func (fg *fgEntity) Draw(t *pixelgl.Canvas) {
	fg.sprite.Draw(t, pixel.IM.Moved(fg.sprite.Frame().Max.Scaled(.5)))
}

type bgEntity struct {
	sprite *pixel.Sprite
}

func (bg *bgEntity) Update(w *pixelgl.Window, dt float64) {}
func (bg *bgEntity) Draw(t *pixelgl.Canvas) {
	bg.sprite.Draw(t, pixel.IM.Moved(bg.sprite.Frame().Max.Scaled(.5)))
}

func newFG(path string) *fgEntity {
	spr, _ := graphics.LoadSprite(path, pixel.R(0, 0, 640, 640))
	return &fgEntity{spr}
}

func newBG(path string) *bgEntity {
	spr, _ := graphics.LoadSprite(path, pixel.R(0, 0, 640, 640))
	return &bgEntity{spr}
}
