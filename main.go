package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/entity"
	"github.com/thomascking/project-nietzsche/ttscreen"
)

var entities []entity.Entity

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Nietzsche",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := pixelgl.NewWindow(cfg)
	ttscreen.InitText()
	if err != nil {
		panic(err)
	}

	player := entity.NewPlayer(pixel.IM.Moved(win.Bounds().Center()))
	entities = append(entities, player)

	for !win.Closed() {

		for _, e := range entities {
			e.Update(win)
		}

		for _, e := range entities {
			e.Draw(win)
		}

		t := pixel.Target(win.Canvas())
		ttscreen.DrawText(&t, "Hello, World!", 100, 500)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
