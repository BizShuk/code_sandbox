package main

import (
	"fmt"
)

var sem = make(chan int, 0)

func handler(index int, t chan int) {
	// with sequential
	for t_request := range t {
		fmt.Println(index, t_request)
	}
	/* without sequential
	for {
		t_request := <-t
		fmt.Println(index, t_request)
	}
	*/
}

func send_request(c chan int) {
	t := make(chan int, 100)

	for i := 0; i < 10; i++ {
		go handler(i, t)
	}
	for i := 0; ; i++ {
		t <- i
	}

	c <- 1
}

func main() {
	c := make(chan int)
	go send_request(c)
	<-c
}
