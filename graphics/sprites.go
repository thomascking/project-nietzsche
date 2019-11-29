package graphics

import (
	"image"
	"os"

	_ "image/png" // so we can load png files

	"github.com/faiface/pixel"
)

// LoadSprite loads a sprite
func LoadSprite(path string, bounds pixel.Rect) (*pixel.Sprite, error) {
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
