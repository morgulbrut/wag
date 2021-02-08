package font

import (
	"encoding/json"
	"image"
	"image/draw"
	"io/ioutil"

	"github.com/morgulbrut/wag/img"
)

type BitmapFont struct {
	Font     string   `json:"font"`
	Tilesize Tilesize `json:"tilesize"`
	Chars    []Char   `json:"chars"`
}
type Tilesize struct {
	SizeX int `json:"sizeX"`
	SizeY int `json:"sizeY"`
}
type Char struct {
	Name string `json:"name"`
	PosX int    `json:"posX"`
	PosY int    `json:"posY"`
}

func ReadFont(filename string) BitmapFont {
	file, _ := ioutil.ReadFile(filename)
	data := BitmapFont{}
	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		panic(err)
	}
	return data
}

func FindChar(char string, font BitmapFont) (Char, int) {
	for i, ch := range font.Chars {
		if ch.Name == char {
			return ch, i
		}
	}
	return Char{}, -1
}

func WriteText(text string, ft BitmapFont) *image.RGBA {

	chars := img.LoadImage(ft.Font)

	canvas := image.NewRGBA(image.Rect(0, 0, len(text)*ft.Tilesize.SizeX, ft.Tilesize.SizeY))

	for i, c := range text {
		ch, _ := FindChar(string(c), ft)
		src := img.GetSubimage(chars, ch.PosX, ch.PosY, ft.Tilesize.SizeX, ft.Tilesize.SizeY)

		b := src.Bounds()
		letter := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(letter, letter.Bounds(), src, b.Min, draw.Src)

		draw.Draw(canvas,
			image.Rect(i*ft.Tilesize.SizeX, 0, (i+1)*ft.Tilesize.SizeX, ft.Tilesize.SizeY),
			letter,
			image.Point{},
			draw.Over)
	}
	return canvas
}
