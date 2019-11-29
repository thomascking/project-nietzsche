package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/entity"
	"github.com/thomascking/project-nietzsche/ttscreen"
	"golang.org/x/image/colornames"
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
		win.Clear(colornames.White)
		for _, e := range entities {
			e.Update(win)
		}

		for _, e := range entities {
			e.Draw(win)
		}

		t := pixel.Target(win.Canvas())
		px := win.Bounds().W()
		py := win.Bounds().H()
		ttscreen.DrawText(&t, "A Game for All and None", pixel.V((px/2)-160, 0), pixel.V(0, py-25), colornames.Black)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
