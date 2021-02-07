package main

import (
	"image"
	"image/draw"

	"github.com/morgulbrut/wag/font"
	"github.com/morgulbrut/wag/img"
)

type imgSize image.Point

func main() {

	ft := font.ReadFont("32X32-FA.json")
	chars := img.LoadImage(ft.Font)

	bla := "HELLO WORLD"
	var letter image.Image

	canvas := image.NewRGBA(image.Rect(0, 0, len(bla)*ft.Tilesize.SizeX, ft.Tilesize.SizeY))

	for i, c := range bla {
		ch, _ := font.FindChar(string(c), ft)
		letter = img.GetSubimage(chars, ch.PosX, ch.PosY, ft.Tilesize.SizeX, ft.Tilesize.SizeY)

		draw.Draw(canvas,
			image.Rect(i*ft.Tilesize.SizeX, 0, (i+1)*ft.Tilesize.SizeX, ft.Tilesize.SizeY),
			letter,
			image.Point{},
			draw.Over)
	}

	img.WriteImage("line.png", canvas)
}
