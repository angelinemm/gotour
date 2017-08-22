package main

import (
    "golang.org/x/tour/pic"
	"image"
	"image/color"
)

// Image has a max x and a max y
type Image struct{
    dx int
	dy int
}

// Bounds returns the bounds of the image
func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.dx, i.dy)
}

// At gives the color of the image at point x, y
func (i Image) At(x, y int) color.Color {
    v := uint8(x*y+(x^y))
    return color.RGBA{v, v, 255, 255}
}

// ColorModel return RGBA color model
func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func main() {
	m := Image{180, 180}
	pic.ShowImage(m)
}
