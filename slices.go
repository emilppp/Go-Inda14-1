package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var z [][]uint8
	var c []uint8
	x := 2.0
	y := 3.0
	for i := 0; i < dx; i++ {
		c = append(c, uint8((x*y)/2))
		y++
		x++
	}

	for a := 0; a < dy; a++ {
		z = append(z, c)
	}

	return z
}

func main() {
	pic.Show(Pic)
}
