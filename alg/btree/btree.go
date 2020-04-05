package btree

type Interface interface {
	Less(other Interface) bool
}

// unbalanced btree
type BTree struct {
	*Node
}

func New() *BTree {
	return &BTree{}
}

func (tree *BTree) insert(v Interface) {
	if tree.Node == nil {
		tree.Node = &Node{
			parent: nil,
			right:  nil,
			left:   nil,
			Val:    v,
		}
		return
	}

	tree.Node.insert(v)
}

type Node struct {
	parent *Node
	right  *Node
	left   *Node
	Val    Interface
}

func (node *Node) insert(v Interface) {
	if node.Val.Less(v) {
		if node.right != nil {
			node.right.insert(v)
		} else {
			node.right = &Node{
				parent: node,
				right:  nil,
				left:   nil,
				Val:    v,
			}
		}
	} else {
		if node.left != nil {
			node.left.insert(v)
		} else {
			node.left = &Node{
				parent: node,
				right:  nil,
				left:   nil,
				Val:    v,
			}
		}
	}
}
