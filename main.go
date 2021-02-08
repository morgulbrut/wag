package main

import (
	"fmt"
	"image"
	"os"

	"github.com/akamensky/argparse"
	"github.com/morgulbrut/wag/font"
	"github.com/morgulbrut/wag/img"
)

type imgSize image.Point

func main() {

	parser := argparse.NewParser("wag", "Make text look awesome using bitmap fonts")

	fontname := parser.String("f", "font", &argparse.Options{Required: true, Help: "<font>.json you want to use"})
	input := parser.String("t", "text", &argparse.Options{Required: true, Help: "text you want to write"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	ft := font.ReadFont(*fontname)

	text := font.WriteLine(*input, ft)

	img.WriteImage("text.png", text)
}
