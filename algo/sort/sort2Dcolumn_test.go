package sort

import (
	"fmt"
	"testing"
)

func TestSort2Dcolumn(t *testing.T) {
	//[[],[4],[1],[0],[4,1],[4,0],[1,0],[4,1,0]]
	a1 := Sort2Dcolumn([][]int{
		[]int{},
		[]int{4},
		[]int{1},
		[]int{0},
		[]int{1, 4},
		[]int{0, 4},
		[]int{0, 1},
		[]int{0, 1, 4},
	})
	fmt.Println(a1)
}
