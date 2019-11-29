package graphics

import "github.com/faiface/pixel"

// Graphic renderable graphic
type Graphic interface {
	Draw(target pixel.Target, m pixel.Matrix)
}
