package img

import (
	"image"
	"image/png"
	"os"

	"github.com/morgulbrut/color256"
)

func LoadImage(filename string) image.Image {
	color256.PrintHiCyan("Reading %s", filename)
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	color256.PrintHiCyan("  Height: %d", img.Bounds().Max.Y)
	color256.PrintHiCyan("  Width: %d", img.Bounds().Max.X)
	return img
}

func WriteImage(filename string, img image.Image) {
	color256.PrintHiCyan("Writing %s", filename)
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func GetSubimage(img image.Image, posX, posY, sizeX, sizeY int) image.Image {
	si := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(posX, posY, posX+sizeX, posY+sizeY))
	return si
}
