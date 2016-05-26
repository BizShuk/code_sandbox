package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%f\n", Sqrt(1))
	fmt.Printf("%f\n", Sqrt(2))
	fmt.Printf("%f\n", Sqrt(4))
	fmt.Printf("%f\n", Sqrt(5))
	fmt.Printf("%f\n", Sqrt(9))
	fmt.Printf("%f\n", Sqrt(16))
}

func Sqrt(x float64) float64 {
	z := x

	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}
