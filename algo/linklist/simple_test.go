// test godoc
package linklist

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	n1 := &Simplenode{Value: "t1"}
	n2 := &Simplenode{Value: "t2"}
	n3 := &Simplenode{Value: "t3"}
	n4 := &Simplenode{Value: "t4"}
	n5 := &Simplenode{Value: "t5"}
	n6 := &Simplenode{Value: "t6"}
	n7 := &Simplenode{Value: "t7"}

	simplelist := Simplelist{}
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

	fmt.Println("Linklist reverse")
	simplelist.Reverse()
	simplelist.Printall()

	simplelist.Reverse()
	simplelist.Printall()

	fmt.Println("Linklist copy")
	newcopy := simplelist.Copy()
	newcopy.Printall()
	newcopy.P.Value = "99"
	simplelist.Printall()
	newcopy.Printall()

}
