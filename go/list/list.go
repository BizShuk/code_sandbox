package list

import "fmt"

type List struct {
	Cur  *Listnode
	Head *Listnode
	Tail *Listnode
	Size int
}

func (l *List) push_back(n *Listnode) error {

	l.Size++

	if l.Tail == nil {
		l.Cur = n
		l.Head = n
	} else {
		l.Tail.Next = n
	}

	l.Tail = n

	return nil
}

func (l *List) Get(i int) *Listnode {
	cur := l.Head

	for i > 0 && cur != nil {
		cur = cur.Next
		i--
	}

	return cur
}

func (l *List) GetCurrent() *Listnode {
	return l.Cur
}

func (l *List) Next() *Listnode {
	if l.Cur.Next != nil {
		l.Cur = l.Cur.Next
	}

	return l.Cur
}

func (l *List) Show() {
	n := l.Head

	for n != nil {
		fmt.Print(n.Val, "(")
		if n.Aptr != nil {
			fmt.Print(n.Aptr.Val)
		}
		fmt.Print(")")
		if n.Next != nil {
			fmt.Print(" -> ")
		}
		n = n.Next
	}

	fmt.Println()
}

// simple copy for next
func (l *List) Copy() *List {
	new_list := &List{}

	var n *Listnode
	cur := l.Cur

	for cur != nil {
		n = cur.Copy()
		new_list.push_back(n)
		cur = cur.Next
	}

	return new_list
}

func (l *List) CopyArbitrary() *List {
	new_list := l.Copy()

	oA_ptr, o_cur, newA_ptr, new_cur := l.Head, l.Head, new_list.Head, new_list.Head

	for oA_ptr != nil {
		o_cur = l.Head
		new_cur = new_list.Head

		if oA_ptr.Aptr != nil {
			for o_cur != nil {
				if oA_ptr.Aptr == o_cur {
					newA_ptr.Aptr = new_cur
				}

				o_cur = o_cur.Next
				new_cur = new_cur.Next
			}
		}

		oA_ptr = oA_ptr.Next
		newA_ptr = newA_ptr.Next
	}
	return new_list
}

func (l *List) IsEnd() bool {
	if l.Tail == l.Cur {
		return true
	}

	return false
}

type Listnode struct {
	Val  interface{}
	Next *Listnode
	Aptr *Listnode // arbitrary pointer
}

func (n *Listnode) Copy() *Listnode {
	new_node := &Listnode{n.Val, nil, nil}
	return new_node
}
