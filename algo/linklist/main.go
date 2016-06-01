package main

import (
	"./linklist"
	"fmt"
)

func main() {
	n1 := &linklist.Simplenode{Value: "t1"}
	n2 := &linklist.Simplenode{Value: "t2"}
	n3 := &linklist.Simplenode{Value: "t3"}
	n4 := &linklist.Simplenode{Value: "t4"}
	n5 := &linklist.Simplenode{Value: "t5"}
	n6 := &linklist.Simplenode{Value: "t6"}
	n7 := &linklist.Simplenode{Value: "t7"}

	simplelist := linklist.Simplelist{}
	simplelist.Printall()

	sn1, err := simplelist.Remove()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%v\n", sn1)
	simplelist.Printall()

	sn2, err := simplelist.Shift()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%v\n", sn2)
	simplelist.Printall()

	len := simplelist.Unshift(n7)
	simplelist.Printall()

	fmt.Printf("%v\n", simplelist.Append(n1))
	fmt.Printf("%v\n", simplelist.Append(n2))
	fmt.Printf("%v\n", simplelist.Append(n3))
	fmt.Printf("%v\n", simplelist.Append(n4))
	fmt.Printf("%v\n", simplelist.Append(n5))
	fmt.Printf("list len:%v\n", simplelist.Len())
	simplelist.Printall()

	i := simplelist.Head
	for i != nil {
		fmt.Printf("%s\n", i.Value)
		i = simplelist.Iterator()
	}

	simplelist.Printall()

	sn3, err := simplelist.Shift()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%v\n", sn3)
	simplelist.Printall()

	len = simplelist.Unshift(n6)
	fmt.Printf("%d\n", len)
	simplelist.Printall()

	sn4, err := simplelist.Remove()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Printf("%v\n", sn4)
	simplelist.Printall()

	simplelist.Reverse()
	simplelist.Printall()

	simplelist.Reverse()
	simplelist.Printall()
}
