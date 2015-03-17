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
	for i := 0; i < dy; i++ {
		for j := 0; j < dx; j++ {
			switch {
			case j%5 == 0:
				a[i][j] = uint8(i * j)
			case j%2 == 0:
				y := 2.0
				x := 10.0
				a[i][j] = uint8(math.Pow(x, y))
			default:
				a[i][j] = uint8((i + j) / 2)

			}
		}

	}

	return a
}

func main() {
	pic.Show(Pic)
}
