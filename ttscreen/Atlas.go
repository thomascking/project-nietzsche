package ttscreen

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//Atlas for all text in the game
var (
	Atlas *text.Atlas
)

//InitText initalizes the font face and the atlas for said font face
func InitText() {
	//TODO: Find a font for the game
	face, err := loadTTF("./fonts/sansation/Sansation-Regular.ttf", 24) //Some arb TTF font, I prefer Fira Mono Sans
	if err != nil {
		panic(err)
	}
	Atlas = text.NewAtlas(face, text.ASCII)
}

func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

//DrawText (t *pixel.Target, s []string, pos pixel.Vec, textColor color.RGBA)
//draws the slice of strings at position pos with color textColor on the specified target T
//target can be a window, canvas, or generic triangle
//alignment 0 = left, 1 = center, 2 = right
func DrawText(t pixel.Target, s []string, v pixel.Vec, c color.RGBA, alignment int) {
	txt := text.New(v, Atlas)
	txt.Color = c
	for _, line := range s {
		currBounds := txt.BoundsOf(line)
		switch alignment {
		case 0:
			break
		case 1:
			txt.Dot.X -= currBounds.W() / 2
			break
		case 2:
			txt.Dot.X += currBounds.W() / 2
			break
		default:
			txt.Dot.X += 0
		}
		fmt.Fprintln(txt, line)
	}
	switch alignment {
	case 0:
		txt.Draw(t, pixel.IM)
	case 1:
		txt.Draw(t, pixel.IM)
	case 2:
		txt.Draw(t, pixel.IM)
	default:
		txt.Draw(t, pixel.IM)
	}
}
