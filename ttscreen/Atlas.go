package ttscreen

import (
	"fmt"
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

//DrawText (string,int,int) draws a string at x,y implicit to a window
func DrawText(t *pixel.Target, s string, x int, y int) {
	fx, fy := float64(x), float64(y)
	txt := text.New(pixel.V(fx, fy), atlas)
	fmt.Fprintln(txt, s)
	txt.LineHeight = atlas.LineHeight() * 1.25
	txt.Draw(*t, pixel.IM.Scaled(txt.Orig, 2))
}
