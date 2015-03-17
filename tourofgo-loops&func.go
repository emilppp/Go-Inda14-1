package main

import (
	"fmt"
	"math"
)

func main() {
	testSqrt(1)
	testSqrt(2)
	testSqrt(3)
	testSqrt(5)
	testSqrt(10)
}

func Sqrt(x float64) float64 {
	z := 2.0
	s := 0.0
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	s = z
	return s
}

func testSqrt(x float64) {
	fmt.Println("Using newtons method:", Sqrt(x), " Using math.Sqrt():", math.Sqrt(x))
}
