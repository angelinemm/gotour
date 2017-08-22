package main

import (
    "golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct{
    dx int
	dy int
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.dx, i.dy)
}

func (i Image) At(x, y int) color.Color {
    v := uint8(x*y+(x^y))
    return color.RGBA{v, v, 255, 255}
}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func main() {
	m := Image{180, 180}
	pic.ShowImage(m)
}