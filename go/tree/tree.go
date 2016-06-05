package tree

type Btree struct {
	root *Node
}

func (t *Btree) height() {

}

func (t *Btree) insert() {

}

func (t *Btree) delete() {

}

// null = 0
// root = 1
func TreeHeight_byindex(n int) (h int) {
	h = 0
	if n <= 0 {
		return 0
	}
	for n/2 >= 1 {
		h++
		n /= 2
	}

	h++
	return h
}

type Node struct {
	Val   interface{}
	left  *Node
	right *Node
}
