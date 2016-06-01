package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Swap with out new variable\n")
	x, y := 4, 3

	fmt.Printf("XOR:\n")
	x, y = XOR(x, y)
	fmt.Printf("x:%d y:%d\n", x, y)

	fmt.Printf("+-:\n")
	x, y = op1(x, y)
	fmt.Printf("x:%d y:%d\n", x, y)

	fmt.Printf("*/:\n")
	x, y = op2(x, y)
	fmt.Printf("x:%d y:%d\n", x, y)
}

func XOR(x, y int) (x1, y1 int) {
	x1 = x ^ y
	y1 = x1 ^ y
	x1 = x1 ^ y1
	return
}

func op1(x, y int) (x1, y1 int) {
	x1 = x + y
	y1 = x1 - y
	x1 = x1 - y1
	return
}

func op2(x, y int) (x1, y1 int) {
	x1 = x * y
	y1 = x1 / y
	x1 = x1 / y1
	return
}
