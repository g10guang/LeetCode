package tree

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"time"
	"sort"
)

func TestBSTSize(t *testing.T) {
	var T bst = &tree{root: nil}
	assert.True(t, T.size() == 0)
	for i := 0; i < 10; i++ {
		assert.False(t, T.isIn(&intElem{i}))
		assert.True(t, T.put(&intElem{i}, i))
	}

	assert.True(t, T.size() == 10)
	assert.True(t, T.put(&intElem{10}, 100))
	assert.False(t, T.put(&intElem{1}, 1))
	assert.True(t, T.size() == 11)
}

func TestBSTisIn(t *testing.T) {
	var T bst = &tree{root: nil}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	elems := make(map[*intElem]interface{})
	for i := 0; i < 10000; i++ {
		x := &intElem{r.Int()}
		if ok := elems[x]; ok == true {
			assert.True(t, T.isIn(x))
			assert.False(t, T.put(x, x.v))
		} else {
			assert.False(t, T.isIn(x))
			assert.True(t, T.put(x, x.v))
			elems[x] = x
		}
	}
}

func TestBSTGetAndDelete(t *testing.T) {
	var T bst = &tree{root: nil}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	elems := make(map[*intElem]interface{})
	for i := 0; i < 4; i++ {
		x := &intElem{r.Intn(100)}
		elems[x] = nil
		T.put(x, x.v)
	}
	expectedSize := len(elems)
	for k := range elems {
		flag, v := T.get(k)
		assert.True(t, flag)
		vv, ok := v.(int)
		assert.True(t, ok)
		assert.Equal(t, k.v, vv)
		// delete
		flag, v = T.delete(k)
		expectedSize--
		assert.True(t, flag)
		assert.Equal(t, expectedSize, T.size())
	}
}

func TestMinMax(t *testing.T) {
	var T bst = &tree{root: nil}
	flag, _, _ := T.max()
	assert.False(t, flag)
	flag, _, _ = T.min()
	assert.False(t, flag)
	flag, _, _ = T.deleteMax()
	assert.False(t, flag)
	flag, _, _ = T.deleteMin()
	assert.False(t, flag)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	elems := make([]int, 0)
	elemMap := make(map[int]interface{})
	for i := 0; i < 1000; i++ {
		x := r.Intn(1000)
		if _, ok := elemMap[x]; !ok {
			elems = append(elems, x)
			elemMap[x] = nil
		}
		T.put(&intElem{x}, x)
	}
	assert.Equal(t, len(elems), T.size())
	sort.Ints(elems)
	flag, k, _ := T.max()
	assert.True(t, flag)
	assert.Equal(t, 0, k.compare(&intElem{elems[len(elems)-1]}))
	flag, k, _ = T.min()
	assert.True(t, flag)
	assert.Equal(t, 0, k.compare(&intElem{elems[0]}))

	// deleteMin
	flag, k, _ = T.deleteMin()
	assert.True(t, flag)
	assert.Equal(t, 0, k.compare(&intElem{elems[0]}))

	// deleteMax
	flag, k, _ = T.deleteMax()
	assert.True(t, flag)
	assert.Equal(t, 0, k.compare(&intElem{elems[len(elems)-1]}))

	assert.Equal(t, len(elems)-2, T.size())
}

func TestBSTDelete(t *testing.T) {
	var T bst = &tree{nil}
	T.put(&intElem{5}, 5)
	T.put(&intElem{4}, 4)
	T.put(&intElem{6}, 6)
	T.delete(&intElem{5})
	assert.Equal(t, 2, T.size())
}

func TestBSTFloorCeil(t *testing.T) {
	var T bst = &tree{nil}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	elems := make([]int, 0)
	for i := 0; i < 1000; i++ {
		x := r.Int()
		duplicate := false
		for _, v := range (elems) {
			if v == x {
				duplicate = true
				break
			}
		}
		if duplicate {
			continue
		}
		T.put(&intElem{x}, x)
		elems = append(elems, x)
	}
	assert.Equal(t, len(elems), T.size())
	isBST(t, T)
	sort.Ints(elems)
	// test floor
	index := r.Intn(len(elems)-1) + 1
	testFloor(elems[index], T, t, true, elems[index])
	testFloor(elems[0]-1, T, t, false, 0)
	if elems[1]-elems[0] > 1 {
		testFloor(elems[0]+1, T, t, true, elems[0])
	}
	// test ceil
	index = r.Intn(len(elems) - 1)
	testCeil(elems[index], T, t, true, elems[index])
	testCeil(elems[len(elems)-1]+1, T, t, false, 0)
	if elems[len(elems)-1]-elems[len(elems)-2] > 1 {
		testCeil(elems[len(elems)-1]-1, T, t, true, elems[len(elems)-1])
	}
}

func testCeil(index int, T bst, t *testing.T, expectedFlag bool, expectedK int) {
	flag, k, _ := T.ceil(&intElem{index})
	assert.Equal(t, expectedFlag, flag)
	if expectedFlag {
		switch v := k.(type) {
		case *intElem:
			assert.Equal(t, expectedK, v.v)
		default:
			panic("Type error")
		}
	}
}

func testFloor(key int, T bst, t *testing.T, expectedFlag bool, expectedK int) {
	flag, k, _ := T.floor(&intElem{key})
	assert.Equal(t, expectedFlag, flag)
	if expectedFlag {
		switch v := k.(type) {
		case *intElem:
			assert.Equal(t, expectedK, v.v)
		default:
			panic("Type error")
		}
	}
}

func isBST(t *testing.T, T bst) {
	tr, ok := T.(*tree)
	if !ok {
		panic("type error")
	}
	isBstNode(t, tr.root)
}

func isBstNode(t *testing.T, n *node) {
	if n == nil {
		return
	}
	if n.left != nil {
		assert.True(t, n.k.compare(n.left.k) >= 0)
		isBstNode(t, n.left)
	}
	if n.right != nil {
		assert.True(t, n.k.compare(n.right.k) <= 0)
		isBstNode(t, n.right)
	}
}
