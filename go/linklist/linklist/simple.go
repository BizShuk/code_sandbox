package linklist

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
func NewSimpleList() *simplelist {
	new := &simplelist{}
	return new
}

type simplelist struct {
	p    *simplenode
	head *simplenode
	tail *simplenode
}

func (l simplelist) Len() int {
	count := 0
	p := l.head
	if p != nil {
		count++
		for p.next != nil {
			count++
			p = p.next
		}
	}
	return count
}

type simplenode struct {
	value string
	next  *simplenode
	prev  *simplenode
}
