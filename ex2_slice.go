package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	s1:= make([][]uint8,dy)
	for y:=0;y<dy;y++{
		s2 := make([]uint8,dx)
		for x:=0;x <dx;x++{
			s2[x] = uint8((x+y)/2)
			s2[x] = uint8(x*y)
		}
		s1[y] =s2
	}
	return s1
}

func main() {
	pic.Show(Pic)
}
