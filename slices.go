package main

import (
	"golang.org/x/tour/pic"
	"math"
)

func Pic(dx, dy int) [][]uint8 {
	// Skapar först en slice med längden dy
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ { // Fyller sedan a med slices av längden dx
		a[i] = make([]uint8, dx)
	}
	x := 3.0
	y := 5.0

	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			switch {
			case j%5 == 0:
				a[i][j] = uint8(x * y)
			case j%3 == 0:
				a[i][j] = uint8((x + y) / 2)
			case j%2 == 0:
				a[i][j] = uint8(math.Pow(x, y))
			default:
				a[i][j] = 50

			}
		}

	}

	return a
}

func main() {
	pic.Show(Pic)
}
