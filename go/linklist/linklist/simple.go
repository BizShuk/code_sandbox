package linklist

import (
	"errors"
	"fmt"
)

/*
type linklist interface {
	len() int
	head() *node
	tail() *node
	insert(node) error
	delete(node) bool
	is_tail() bool
	iterator() *node
}

type node interface {
	next() *node
	prev() *node
}
*/

type Simplelist struct {
	P    *Simplenode
	Head *Simplenode
	Tail *Simplenode
	len  int
}

func (l *Simplelist) error(s string) string {
	return s
}

func (l *Simplelist) Len() int {
	count := 0
	for p := l.Head; p != nil; p = p.next {
		count++
	}
	return count
}

func (l *Simplelist) Iterator() *Simplenode {
	l.P = l.P.next
	return l.P
}

func (l *Simplelist) Printall() error {
	fmt.Printf("\nlist:")
	for i, c := 1, l.Head; c != nil; i, c = i+1, c.next {
		fmt.Printf("%d:%s ,", i, c.Value)
	}
	fmt.Printf("len:%d\n\n", l.len)
	return nil
}

func (l *Simplelist) Unshift(n *Simplenode) int {
	if l.Head == nil {
		l.Tail = n
		l.P = n
	}

	n.next = l.Head
	l.Head = n
	l.len++
	return l.len
}

func (l *Simplelist) Shift() (*Simplenode, error) {
	if l.Head == nil {
		return nil, errors.New("No Element could be removed")
	}

	shifted := l.Head
	l.Head = shifted.next
	l.len--
	return shifted, nil

}

func (l *Simplelist) Append(n *Simplenode) int {
	l.len++
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		l.P = n
		return l.len
	}

	l.Tail.next = n
	l.Tail = n
	return l.len

}

func (l *Simplelist) Remove() (*Simplenode, error) {
	if l.Head == nil {
		return nil, errors.New("Mo Element could be removed")
	}

	c := l.Head
	if c == l.Tail {
		l.Head = nil
		l.Tail = nil
		return c, nil
	}

	for c.next.next != nil {
		c = c.next
	}

	l.Tail = c
	c = c.next
	l.Tail.next = nil
	l.len--

	return c, nil

}

func (l *Simplelist) is_tail(n *Simplenode) bool {
	if l.Tail == n {
		return true
	}
	return false
}

type Simplenode struct {
	Value string
	next  *Simplenode
}
