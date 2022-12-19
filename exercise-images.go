package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct{}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (m Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), uint8(x^y), 255} 
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
