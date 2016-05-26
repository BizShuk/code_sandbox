package main

import (
	"./sorts"
	"fmt"
	"log"
	"reflect"
	"time"
)

type sortingfunc func(src []int) (int, int)

func main() {
	times := 1
	log.Println("Start...")
	case1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	case2 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	case3 := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10}
	case4 := []int{1, 6, 2, 7, 3, 8, 4, 9, 5, 10}

	case5 := []int{1, 2, 3, 4, 5, 10, 9, 8, 7, 6}
	case6 := []int{5, 4, 3, 2, 1, 10, 9, 8, 7, 6}

	case7 := []int{9, 1, 6, 4, 10, 8, 7, 2, 3, 5}
	case8 := []int{3, 4, 5, 9, 10, 7, 6, 8, 1, 2}
	case9 := []int{5, 10, 2, 3, 1, 8, 9, 6, 4, 7}
	case0 := [][]int{case1, case2, case3, case4, case5, case6, case7, case8, case9}

	//call_sorting(sorts.Bubble_sort, case0, times)
	//call_sorting(sorts.Selection_sort, case0, times)
	//call_sorting(sorts.Insertion_sort, case0, times)
	call_sorting(sorts.Quick_sort, case0, times)

}

//
func call_sorting(T sortingfunc, cases [][]int, times int) {
	var ct, et int
	var t0, t1 time.Time
	var d time.Duration
	var cases_cmp []time.Duration

	fmt.Println("\n")
	for i := 0; i < 9; i++ {
		_case := make([]int, len(cases[i]), cap(cases[i]))
		t0 = time.Now()
		for j := 0; j < times; j++ {
			copy(_case, cases[i])
			ct, et = T(_case)
		}
		t1 = time.Now()
		d = t1.Sub(t0)
		cases_cmp = append(cases_cmp, d)

		fmt.Printf("Case %d: %v  sorted:%v ,", i+1, cases[i], _case)
		fmt.Printf("comparsion:%v , exchangation:%v run %v\n", ct, et, d)
	}

	result_stat := case_cmp(cases_cmp)
	log.Printf("Current function:%s , cases:%d , times:%d\n", reflect.TypeOf(T).Name(), len(cases), times)
	fmt.Printf("Worst: case%d %v %v\n", result_stat["worst"].i, cases[result_stat["worst"].i], result_stat["worst"].duration)
	fmt.Printf("Best: case%d %v %v\n", result_stat["best"].i, cases[result_stat["best"].i], result_stat["best"].duration)
	fmt.Printf("Avg duration:%v\n", result_stat["avg"].duration)

	fmt.Println("")
}

type case_stat struct {
	i        int
	duration time.Duration
}

func case_cmp(cases []time.Duration) map[string]case_stat {
	var total time.Duration
	worst, best, avg := case_stat{}, case_stat{}, case_stat{}
	for k, v := range cases {
		//fmt.Printf("%v %v\n", k, v)
		if worst.duration == 0 || best.duration < v {
			worst.i, worst.duration = k, v
		}

		if best.duration == 0 || best.duration > v {
			best.i, best.duration = k, v
		}
		total += v
	}

	avg.duration = time.Duration(float64(total) / float64(len(cases)))
	return map[string]case_stat{"worst": worst, "best": best, "avg": avg}
}
