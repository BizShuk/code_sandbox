package tree

import (
	"fmt"
	"testing"
)

func TestTreeHeight_byindex(t *testing.T) {
	cases := 0
	a := TreeHeight_byindex(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 1
	a = TreeHeight_byindex(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 2
	a = TreeHeight_byindex(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 25
	a = TreeHeight_byindex(cases)
	fmt.Println("case:", cases, "result:", a)
	cases = 4
	a = TreeHeight_byindex(cases)
	fmt.Println("case:", cases, "result:", a)
}
