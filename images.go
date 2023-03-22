package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)
type Image struct{}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (i Image) At(x,y int) color.Color {
	return color.RGBA{0, 0, 255, 255}
}

// image -> image
func main() {
	i := Image{}
	pic.ShowImage(i)
}
