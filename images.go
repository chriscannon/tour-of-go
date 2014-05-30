package main

import (
	"image"
	"image/color"

	"code.google.com/p/go-tour/pic"
)

// Image is a struct that holds the height (H) and width (W) of an image.
type Image struct {
	H, W int
}

// ColorModel return the color model for an Image.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the rectangular bounds for an Image.
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.W, i.H)
}

// At returns the color of the pixel at x, y.
func (i Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{200, 200}
	pic.ShowImage(m)
}
