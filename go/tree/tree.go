package tree

type Tree struct {
	root *TreeNode
}

func (t *Tree) GetHeight() {

}

func (t *Tree) Insert() {

}

func (t *Tree) Delete() {

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

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) MaxDepth() int {
	if t == nil {
		return 0
	}

	left_max := 0
	if t.Left != nil {
		left_max = t.Left.MaxDepth()
	}
	right_max := 0
	if t.Right != nil {
		right_max = t.Right.MaxDepth()
	}

	if left_max > right_max {
		return left_max + 1
	}
	return right_max + 1
}

func (t *TreeNode) Show() {

}
