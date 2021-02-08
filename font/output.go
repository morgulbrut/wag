package font

import (
	"image"
	"image/draw"

	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/wag/img"
)

func WriteLine(text string, ft BitmapFont) *image.RGBA {
	color256.PrintHiGreen("Building %s", text)
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

// func WriteText(text string, ft BitmapFont) *image.RGBA {

// 	maxW := 0
// 	lines := strings.Split(text, "#")
// 	for i, l := range lines {
// 		fmt.Println(i)
// 		fmt.Println(maxW)
// 		if len(l) > maxW {
// 			maxW = len(l)
// 		}
// 	}
// 	fmt.Println(maxW)

// 	canvas := image.NewRGBA(image.Rect(0, 0, maxW*ft.Tilesize.SizeX, len(lines)*ft.Tilesize.SizeY))
// 	for i, l := range lines {

// 		img := WriteLine(l, ft)
// 		draw.Draw(canvas,
// 			image.Rect(0, i*ft.Tilesize.SizeY, ft.Tilesize.SizeX, (i+1)*ft.Tilesize.SizeY),
// 			img,
// 			image.Point{},
// 			draw.Over)
// 	}

// 	return canvas
// }
