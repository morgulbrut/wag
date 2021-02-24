package font

import (
	"embed"
	"encoding/json"
)

type BitmapFont struct {
	Font  string `json:"font"`
	Chars []Char `json:"chars"`
}

type Char struct {
	Name  string `json:"name"`
	PosX  int    `json:"posX"`
	PosY  int    `json:"posY"`
	SizeX int    `json:"sizeX"`
	SizeY int    `json:"sizeY"`
}

func ReadFont(filename string, d embed.FS) BitmapFont {
	// file, err := ioutil.ReadFile(filename)

	fonts, err := d.ReadFile("fontdesc/fonts.json")
	if err != nil {
		panic(err)
	}

	fileMap := make(map[string]string)
	err = json.Unmarshal([]byte(fonts), &fileMap)
	if err != nil {
		panic(err)
	}

	fn := fileMap[filename]
	file, err := d.ReadFile(fn)
	if err != nil {
		panic(err)
	}
	data := BitmapFont{}
	err = json.Unmarshal([]byte(file), &data)
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
