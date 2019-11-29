package main

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/state"
	"github.com/thomascking/project-nietzsche/ttscreen"
	"github.com/thomascking/project-nietzsche/world"
)

func run() {
	/* Engine Initializations */
	cfg := pixelgl.WindowConfig{
		Title:  "Nietzsche",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	ttscreen.InitText()
	/* Engine Initializations Done */

	worldMap := make(map[state.State]world.World)
	worldMap[state.GS] = world.NewGameWorld(pixel.R(0, 0, 640, 480))
	worldMap[state.MS] = world.NewMenuWorld(pixel.R(0, 0, 1024, 768))
	worldMap[state.PS] = world.NewPauseWorld(pixel.R(0, 0, 1024, 768))

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last)
		s := dt.Seconds()
		last = time.Now()
		if state.CurrState == state.ES {
			win.Destroy()
			break
		}
		currWorld := worldMap[state.CurrState]
		currWorld.Update(win, s)
		currWorld.Draw()
		currWorld.Render(win.Canvas())
		win.Update()
		time.Sleep(time.Millisecond*17 - dt)
	}
}

func main() {
	pixelgl.Run(run)
}
