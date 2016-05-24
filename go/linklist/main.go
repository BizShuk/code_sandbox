package main

import (
	"./linklist"
	"fmt"
)

func main() {
	simplelist := linklist.NewSimpleList()

	fmt.Printf("%v\n", simplelist)
	len1 := simplelist.Len()
	fmt.Printf("%v\n", len1)
}
