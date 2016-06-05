package bit

import (
	"fmt"
	"testing"
)

func TestBitSequence(t *testing.T) {
	fmt.Println("")
	var a []int
	for i := 0; i < 30; i++ {
		a = countBits(i)
		fmt.Println("cases:", i, "result:", a)
	}
	cases := 123456
	a = BitSequence(cases)
	fmt.Println("cases:", cases, "result:", a)
}
