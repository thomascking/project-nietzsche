package ttscreen

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"os"
	"unicode"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

var (
	atlas *text.Atlas
)

//InitText initalizes the font face and the atlas for said font face
func InitText() {
	//TODO: Find a font for the game
	/* face, err := loadTTF("./", 12) //Some arb TTF font, I prefer Fira Mono Sans
	if err != nil {
		panic(err)
	} */
	atlas = text.NewAtlas(basicfont.Face7x13, text.ASCII, text.RangeTable(unicode.Latin))
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
func DrawText(t pixel.Target, s []string, v pixel.Vec, c color.RGBA) {
	txt := text.New(v, atlas)
	txt.Color = c
	txt.LineHeight = atlas.LineHeight() * 1.25
	for _, line := range s {
		txt.Dot.X -= txt.BoundsOf(line).W() / 2
		fmt.Fprintln(txt, line)
	}
	txt.Draw(t, pixel.IM.Scaled(txt.Orig, 2))
}
