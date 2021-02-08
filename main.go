package main

import (
	"image"

	"github.com/morgulbrut/wag/font"
	"github.com/morgulbrut/wag/img"
)

type imgSize image.Point

func main() {

	ft := font.ReadFont("32X32-FA.json")
	bla := "'MURICA FUCK YEAH!"
	text := font.WriteText(bla, ft)

	img.WriteImage("text.png", text)
}
