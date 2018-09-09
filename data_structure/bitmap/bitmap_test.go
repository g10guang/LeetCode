package bitmap

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"time"
)

func addBitmap(t *testing.T, b *bitmap) (ret map[int32]interface{}) {
	if b == nil {
		b = &bitmap{}
	}
	ret = make(map[int32]interface{})
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var v int32
	for i := 0; i < 10000; i++ {
		for {
			v = r.Int31()
			if i < 1000 {
				v = -v
			}
			if _, ok := ret[v]; !ok {
				ret[v] = nil
				break
			}
			assert.False(t, b.add(v))
		}
		assert.True(t, b.add(v))
	}
	return ret
}

func TestBitmapIsIn(t *testing.T) {
	b := &bitmap{}
	inElem := addBitmap(t, b)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for k := range inElem {
		assert.True(t, b.isIn(k))
	}

	for i := 0; i < 5000; i++ {
		x := -r.Int31()
		if _, ok := inElem[x]; ok {
			assert.True(t, b.isIn(x))
		} else {
			assert.False(t, b.isIn(x))
		}
	}

	for i := 0; i < 5000; i++ {
		x := r.Int31()
		if _, ok := inElem[x]; ok {
			assert.True(t, b.isIn(x))
		} else {
			assert.False(t, b.isIn(x))
		}
	}
}

func TestBitmapUnion(t *testing.T) {
	b, s := &bitmap{}, &bitmap{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bs, ss := addBitmap(t, b), addBitmap(t, s)
	b.union(s)
	for i := 0; i < 10000; i++ {
		x := r.Int31()
		if i < 5000 {
			x = -x
		}
		isIn := false
		if _, ok := bs[x]; ok {
			isIn = true
		}
		if _, ok := ss[x]; ok {
			isIn = true
		}
		if isIn {
			assert.True(t, b.isIn(x))
		} else {
			assert.False(t, b.isIn(x))
		}
	}

}


func TestBitmapIntersection(t *testing.T) {
	b, s := &bitmap{}, &bitmap{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bs, ss := addBitmap(t, b), addBitmap(t, s)
	b.intersection(s)
	for i := 0; i < 10000; i++ {
		x := r.Int31()
		if i < 5000 {
			x = -x
		}
		_, isIn := bs[x]
		if _, ok := ss[x]; isIn && ok {
			assert.True(t, b.isIn(x))
		} else {
			assert.False(t, b.isIn(x))
		}
	}
}

func TestBitmapDifference(t *testing.T) {
	b, s := &bitmap{}, &bitmap{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bs, ss := addBitmap(t, b), addBitmap(t, s)
	b.difference(s)
	for i := 0; i < 10000; i++ {
		x := r.Int31()
		if i < 5000 {
			x = -x
		}
		_, isIn := bs[x]
		if _, ok := ss[x]; (ok && isIn) && (ok || isIn) {
			assert.True(t, b.isIn(x))
		} else {
			assert.False(t, b.isIn(x))
		}
	}
}