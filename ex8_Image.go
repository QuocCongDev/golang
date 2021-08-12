package main

import (
  "golang.org/x/tour/pic"
  "image"
  "image/color"
)
type Image struct{
color uint8
width,height int	
}
func(img Image) Bounds() image.Rectangle{
	return image.Rect(0,0,img.width,img.height)

}
func(img Image) ColorModel() color.Model{
	return color.RGBAModel

}
func(img Image)At(x, y int) color.Color{
	return color.RGBA{img.color+uint8(x),img.color+uint8(y),0,0}
}

func main() {
  m := Image{20, 200, 100}
  pic.ShowImage(m)
}