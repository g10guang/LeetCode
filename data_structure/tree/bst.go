package tree

// binary search tree
type tree struct {
	root *node
}

type node struct {
	left  *node
	right *node
	k     comparable
	v     interface{}
}

type bst interface {
	// The number of node in the bst.
	size() int

	// delete key and return the value.
	// if key not exists, do nothing and return nil.
	delete(comparable) (bool, interface{})

	// put the new key-value pair into bst.
	// if key not exist, create a new node. return true
	// else cover the old value. return false
	put(comparable, interface{}) bool

	// if key exist return its value
	// else return nil.
	get(comparable) (bool, interface{})

	// if size() > 0 return min key-value pair in bst.
	// else return nil.
	min() (bool, comparable, interface{})

	// if size() > 0 return max key-value pair in bst.
	// else return nil.
	max() (bool, comparable, interface{})

	// return key-value pair which <= k.
	// if not exist return nil.
	floor(comparable) (bool, comparable, interface{})

	// return key-value pair which >= k.
	// if not exist return nil.
	ceil(comparable) (bool, comparable, interface{})

	// if size() > 0, delete the min key-value in bst.
	// else do nothing and return nil.
	deleteMin() (bool, comparable, interface{})

	// if size() > 0, delete the max key-value in bst.
	// else do nothing and return nil.
	deleteMax() (bool, comparable, interface{})

	// check k is in the tree or not.
	isIn(comparable) bool
}

type comparable interface {
	// == return -
	// > return 1
	// < return -1
	// panic("type mismatch")
	compare(comparable) int
}

type intElem struct {
	v int
}

func (o *intElem) compare(t comparable) int {
	x, ok := t.(*intElem)
	if !ok {
		panic("TypeError need intElem type.")
	}
	switch {
	case o.v > x.v:
		return 1
	case o.v < x.v:
		return -1
	default:
		return 0
	}
}

func (t *tree) size() int {
	return t.root.size()
}

func (t *tree) delete(k comparable) (bool, interface{}) {
	if t.root == nil {
		return false, nil
	}
	hasDel, delNode := t.root.delete(k)
	if hasDel && delNode == t.root {
		// t.root is to be removed.
		flag, substitute := t.root.left.deleteMax()
		if flag {
			substitute.copyChildrenLinkage(t.root)
			t.root = substitute
		} else {
			// t.root.left == nil
			t.root = t.root.right
		}
	}
	return hasDel, delNode.v
}

// note: to avoid cyclic linkage relation.
func (n *node) copyChildrenLinkage(src *node) {
	if src.left != n {
		n.left = src.left
	} else {
		n.left = nil
	}

	if src.right != n {
		n.right = src.right
	} else {
		n.right = nil
	}
}

func (t *tree) deleteMin() (hasDel bool, k comparable, v interface{}) {
	hasDel, delNode := t.root.deleteMin()
	if hasDel && delNode == t.root {
		// because t.root is the min element, t.root.left == nil
		t.root = delNode.right
	}
	if hasDel {
		k = delNode.k
		v = delNode.v
	}
	return hasDel, k, v
}

func (t *tree) deleteMax() (hasDel bool, k comparable, v interface{}) {
	hasDel, delNode := t.root.deleteMax()
	if hasDel && delNode == t.root {
		// because t.root is the max element, t.root.right == nil
		t.root = delNode.left
	}
	if hasDel {
		k = delNode.k
		v = delNode.v
	}
	return hasDel, k, v
}

// put a new key-value pair into tree.
func (t *tree) put(k comparable, v interface{}) bool {
	if t.root == nil {
		t.root = &node{k: k, v: v, left: nil, right: nil}
		return true
	} else {
		return t.root.put(k, v)
	}
}

func (t *tree) get(k comparable) (bool, interface{}) {
	return t.root.get(k)
}

func (t *tree) min() (bool, comparable, interface{}) {
	return t.root.min()
}

func (t *tree) max() (bool, comparable, interface{}) {
	return t.root.max()
}

func (t *tree) isIn(k comparable) bool {
	return t.root != nil && t.root.isIn(k)
}

func (t *tree) floor(k comparable) (bool, comparable, interface{}) {
	if t.root == nil {
		return false, nil, nil
	}
	return t.root.floor(k)
}

func (t *tree) ceil(k comparable) (bool, comparable, interface{}) {
	if t.root == nil {
		return false, nil, nil
	}
	return t.root.ceil(k)
}

func (n *node) ceil(k comparable) (bool, comparable, interface{}) {
	if n == nil {
		return false, nil, nil
	}
	switch n.k.compare(k) {
	case 1:
		hasFound, fk, fv := n.left.ceil(k)
		if hasFound {
			return hasFound, fk, fv
		}
		// return current node.
		return true, n.k, n.v
	case -1:
		return n.right.ceil(k)
	default:
		return true, n.k, n.v
	}
}

func (n *node) floor(k comparable) (bool, comparable, interface{}) {
	if n == nil {
		return false, nil, nil
	}
	switch n.k.compare(k) {
	case 1:
		return n.left.floor(k)
	case -1:
		hasFound, fk, fv := n.right.floor(k)
		if hasFound {
			return hasFound, fk, fv
		}
		return true, n.k, n.v
	case 0:
		return true, n.k, n.v
	default:
		panic("compare value error")
	}
}

func (n *node) isIn(k comparable) bool {
	if n == nil {
		return false
	}
	switch relate := n.k.compare(k); {
	case relate < 0:
		return n.left.isIn(k)
	case relate > 0:
		return n.right.isIn(k)
	default:
		return true
	}
}

func (n *node) max() (bool, comparable, interface{}) {
	if n == nil {
		return false, nil, nil
	}
	if n.right != nil {
		return n.right.max()
	}
	// current node is the max node in the tree
	return true, n.k, n.v
}

// find the left most node
func (n *node) min() (bool, comparable, interface{}) {
	if n == nil {
		return false, nil, nil
	}
	if n.left != nil {
		return n.left.min()
	}
	// current node is the min node in the tree
	return true, n.k, n.v
}

func (n *node) get(k comparable) (bool, interface{}) {
	if n == nil {
		return false, nil
	}
	switch n.k.compare(k) {
	case -1:
		return n.right.get(k)
	case 1:
		return n.left.get(k)
	default:
		return true, n.v
	}
}

func (n *node) put(k comparable, v interface{}) bool {
	switch n.k.compare(k) {
	case -1:
		if n.right == nil {
			n.right = &node{left: nil, right: nil, k: k, v: v}
			return true
		} else {
			return n.right.put(k, v)
		}
	case 1:
		if n.left == nil {
			n.left = &node{left: nil, right: nil, k: k, v: v}
			return true
		} else {
			return n.left.put(k, v)
		}
	default:
		// ==
		// cover the current node.
		n.v = v
		return false
	}
}

func (n *node) size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.size() + n.right.size()
}

// return the node the be deleted.
func (n *node) delete(k comparable) (hasDel bool, delNode *node) {
	if n == nil {
		return false, nil
	}
	switch n.k.compare(k) {
	case 1:
		// n.k > k
		hasDel, delNode = n.left.delete(k)
		if hasDel && delNode == n.left {
			// current node n's left child has been deleted.
			flag, substitute := delNode.right.deleteMin()
			if flag {
				// rebuild the relation.
				substitute.copyChildrenLinkage(delNode)
				n.left = substitute
			} else {
				// delNode.right == nil
				n.left = delNode.left
			}
			delNode.removeChildren()
		}

	case -1:
		// n.k < k
		hasDel, delNode = n.right.delete(k)
		if hasDel && delNode == n.right {
			flag, substitute := delNode.left.deleteMax()
			if flag {
				// rebuild the relation
				substitute.copyChildrenLinkage(delNode)
				n.right = substitute
			} else {
				// delNode.left == nil
				n.right = delNode.right
			}
			delNode.removeChildren()
		}

	case 0:
		// delete current node
		hasDel = true
		delNode = n

	default:
		panic("Compare relation error.")

	}
	return hasDel, delNode
}

func (n *node) deleteMin() (hasDel bool, delNode *node) {
	if n == nil {
		return false, nil
	}
	if n.left == nil {
		// current node is min node.
		return true, n
	}
	hasDel, delNode = n.left.deleteMin()
	if hasDel && delNode == n.left {
		// need to find the new left child for current node.
		// because delNode is the min node, delNode.left == nil
		n.left = delNode.right
		delNode.removeChildren()
	}
	return hasDel, delNode
}

func (n *node) deleteMax() (hasDel bool, delNode *node) {
	if n == nil {
		return false, nil
	}
	if n.right == nil {
		// current node is the max node.
		return true, n
	}
	hasDel, delNode = n.right.deleteMax()
	if hasDel && delNode == n.right {
		// because delNode is the max node, delNode.right == nil
		n.right = delNode.left
		delNode.removeChildren()
	}
	return hasDel, delNode
}

// delete the linkage of left and right children.
// make GC more efficient.
func (n *node) removeChildren() {
	n.left, n.right = nil, nil
}
