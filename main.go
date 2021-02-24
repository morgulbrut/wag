package main

import (
	"embed"
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/akamensky/argparse"
	"github.com/morgulbrut/wag/font"
	"github.com/morgulbrut/wag/img"
)

type imgSize image.Point

//go:embed fontdesc/*
//go:embed BitmapFonts/32X32-F*.png
var data embed.FS

func main() {

	testString, err := data.ReadFile("fontdesc/test.txt")
	if err != nil {
		panic(err)
	}

	parser := argparse.NewParser("wag", "Make text look awesome (questionable) using bitmap fonts")
	fontname := parser.String("f", "font", &argparse.Options{Default: "font5", Help: "<font>.json you want to use"})
	input := parser.String("t", "text", &argparse.Options{Default: string(testString), Help: "text you want to write"})
	spacing := parser.Int("s", "spacing", &argparse.Options{Default: 0, Help: "Optional spacing "})
	hexcolor := parser.String("c", "color", &argparse.Options{Default: "#000000FF", Help: "Colour to be made transparent"})
	output := parser.String("o", "output", &argparse.Options{Default: "text", Help: "Output file name (without ending)"})

	err = parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	ft := font.ReadFont(*fontname, data)
	text := font.WriteLine(*input, ft, *spacing, data)
	col, _ := ParseHexColor(*hexcolor)
	text = img.ColorToTransparent(text, col)
	img.WriteImage(*output+".png", text)
}

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 9:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 9, 7 or 4")

	}
	return
}
