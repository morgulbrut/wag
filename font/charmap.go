package font

import (
	"encoding/json"
	"io/ioutil"
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
