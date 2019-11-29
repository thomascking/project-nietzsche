package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/state"
	"github.com/thomascking/project-nietzsche/ttscreen"
	"github.com/thomascking/project-nietzsche/world"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Nietzsche",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	worldMap := make(map[state.State]world.World)
	worldMap[state.GS] = world.NewWorld(pixel.R(0, 0, 1024, 768)) // <- Breaks Here
	worldMap[state.MS] = world.NewMenuWorld(pixel.R(0, 0, 1024, 768))
	worldMap[state.PS] = world.NewWorld(pixel.R(0, 0, 1024, 768))

	ttscreen.InitText()
	for !win.Closed() {
		currWorld := worldMap[state.CurrState]
		currWorld.Update(win)
		currWorld.Draw()
		currWorld.Render(win.Canvas())
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
