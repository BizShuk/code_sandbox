package btree

type Node struct {
	Value interface{}
	left  *Node
	right *Node
}
