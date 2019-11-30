package maps

import (
	"github.com/faiface/pixel"
	"github.com/thomascking/project-nietzsche/entity"
	"github.com/thomascking/project-nietzsche/world"
)

// NewTestMap creates a new test map
func NewTestMap(w world.World) {
	entities := []entity.Entity{
		newBG("./images/test-level-bg.png"),
		newFG("./images/test-level-fg.png"),
		entity.NewPlayer(pixel.V(10, 100)),
	}
	for _, e := range entities {
		w.AddEntity(e)
	}
}
