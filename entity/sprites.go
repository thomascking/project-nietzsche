package entity

import (
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
)

func loadSprite(path string, bounds pixel.Rect) (*pixel.Sprite, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	pic := pixel.PictureDataFromImage(img)

	return pixel.NewSprite(pic, bounds), nil
}
