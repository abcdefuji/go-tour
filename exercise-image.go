package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct {
	width int
	height int
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0,0,i.width,i.height)
}

func (i Image) At(x,y int) color.Color {
	v := uint8(x*y)
	return color.RGBA{v,v,255,255}
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{133,50}
	pic.ShowImage(m)
}

