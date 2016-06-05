package bit

import (
	"fmt"
)

func countBits(num int) []int {
	ret := []int{0}
	for i := 0; i <= num; i++ {
		if i == 0 {
			ret[i] = 0
			continue
		}
		ret = append(ret, ret[i>>1]+(i&1))
	}
	return ret

}

func BitSequence(num int) []int {
	var ret = make([]int, num+1)
	fmt.Println("")
	for i := 0; i <= num; i++ {
		ret[i] = ret[i>>1] + (i & 1)
	}
	return ret
}
