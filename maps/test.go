package maps

import (
	"github.com/faiface/pixel"
	"github.com/thomascking/project-nietzsche/entity"
	"github.com/thomascking/project-nietzsche/interfaces"
)

// NewTestMap creates a new test map
func NewTestMap(w interfaces.World) {
	entities := []interfaces.Entity{
		newBG("./images/test-level-bg.png"),
		newFG("./images/test-level-fg.png"),
		entity.NewPlayer(pixel.V(10, 110)),
	}
	for _, e := range entities {
		w.AddEntity(e)
	}
}
