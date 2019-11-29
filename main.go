package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/thomascking/project-nietzsche/ttscreen"
)

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

	for !win.Closed() {
		ttscreen.DrawText(win, "Hello, World!", 100, 500)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
