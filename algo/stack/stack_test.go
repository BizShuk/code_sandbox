package stack

import (
	"fmt"
	"testing"
)

func Teststack(t *testing.T) {
	s := &stack{}

	s.Push(3)
	s.Push(4)
	s.Push(5)
	fmt.Printf("%d\n", s.Pop())
	fmt.Printf("%d\n", s.Pop())
	fmt.Printf("%d\n", s.Pop())

}
