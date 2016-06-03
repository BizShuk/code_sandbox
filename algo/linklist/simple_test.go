// test godoc
package linklist

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	n1 := &ListNode{Val: "t1"}
	n2 := &ListNode{Val: "t2"}
	n3 := &ListNode{Val: "t3"}
	n4 := &ListNode{Val: "t4"}
	n5 := &ListNode{Val: "t5"}
	n6 := &ListNode{Val: "t6"}
	n7 := &ListNode{Val: "t7"}

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
		fmt.Printf("%s\n", i.Val)
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
	newcopy.P.Val = "99"
	simplelist.Printall()
	newcopy.Printall()

}
