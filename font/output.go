package font

import (
	"image"
	"image/draw"

	"github.com/morgulbrut/color256"
	"github.com/morgulbrut/wag/img"
)

func WriteLine(text string, ft BitmapFont, spacing int) *image.RGBA {
	color256.PrintHiGreen("Building %s", text)
	chars := img.LoadImage(ft.Font)

	width := 0

	for _, c := range text {
		ch, _ := FindChar(string(c), ft)
		width += ch.SizeX + spacing
	}

	canvas := image.NewRGBA(image.Rect(0, 0, width, ft.Chars[0].SizeY))

	xPos := 0
	for _, c := range text {
		ch, _ := FindChar(string(c), ft)
		src := img.GetSubimage(chars, ch.PosX, ch.PosY, ft.Chars[0].SizeX, ft.Chars[0].SizeY)

		b := src.Bounds()
		letter := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(letter, letter.Bounds(), src, b.Min, draw.Src)

		draw.Draw(canvas,
			image.Rect(xPos, 0, xPos+ch.SizeX+spacing, ft.Chars[0].SizeY),
			letter,
			image.Point{},
			draw.Over)
		xPos += ch.SizeX + spacing
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

// 	canvas := image.NewRGBA(image.Rect(0, 0, maxW*ft.Chars[0].SizeX, len(lines)*ft.Chars[0].SizeY))
// 	for i, l := range lines {

// 		img := WriteLine(l, ft)
// 		draw.Draw(canvas,
// 			image.Rect(0, i*ft.Chars[0].SizeY, ft.Chars[0].SizeX, (i+1)*ft.Chars[0].SizeY),
// 			img,
// 			image.Point{},
// 			draw.Over)
// 	}

// 	return canvas
// }
