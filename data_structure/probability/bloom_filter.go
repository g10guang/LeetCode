package probability

import (
	"crypto/sha512"
)

// 1 maybe in
// 2 must not in

const ByteSize = 64

type bloomFilter interface {
	// Add a new value into bloom filter bloomFilterStruct.
	add(v string)

	// Check a v is in bloom filter or not.
	isIn(v string) bool
}

type bloomFilterStruct struct {
	arr      [][]byte
	moveMark [16]uint8
}

// Init the bloom filter struct.
// 1. set storage arr length.
// 2. set the hash function
func (b *bloomFilterStruct) init(hashCount int) {
	if hashCount == 0 {
		hashCount = 10
	}
	b.arr = make([][]byte, hashCount)
	// Use sha512 hash function.
	for i := 0; i < hashCount; i++ {
		b.arr[i] = make([]byte, ByteSize)
	}

	for i := uint8(0); i < uint8(8); i++ {
		b.moveMark[i] = 1 << i
		b.moveMark[i+8] = ^(1 << i)
	}
}

func (b *bloomFilterStruct) hash(v []byte, randomBytes []byte) (ret []byte) {
	s := sha512.New()
	s.Write(v)
	ret = s.Sum(randomBytes)
	b.spare(ret)
	return
}

// Make hash function sparse
// Remove consecutive 1-bit in a byte.
func (b *bloomFilterStruct) spare(t []byte) {
	for i := 0; i < len(t); i++ {
		tmp := t[i]
		flag := tmp&b.moveMark[0] == b.moveMark[0]
		for j := uint8(1); j < uint8(7); j++ {
			if tmp&b.moveMark[j] == b.moveMark[j] {
				if flag {
					tmp |= b.moveMark[j+8]
					flag = false
				} else {
					flag = true
				}
			}
		}
		t[i] = tmp
	}
}

func (b *bloomFilterStruct) hashFunc(v string) (ret [][]byte) {
	ret = make([][]byte, len(b.arr))
	bytes := []byte(v)
	for i := 0; i < len(b.arr); i++ {
		ret[i] = b.hash(bytes, []byte{byte(i)})
	}
	return
}

func (b *bloomFilterStruct) add(v string) {
	hashBytes := b.hashFunc(v)
	b.merge(hashBytes)
}

func (b *bloomFilterStruct) merge(t [][]byte) {
	for i := 0; i < len(b.arr); i++ {
		for j := 0; j < ByteSize; j++ {
			b.arr[i][j] |= t[i][j]
		}
	}
}

func (b *bloomFilterStruct) isIn(v string) bool {
	hashBytes := b.hashFunc(v)
	for i := 0; i < len(b.arr); i++ {
		for j := 0; j < ByteSize; j++ {
			if hashBytes[i][j]&b.arr[i][j] != hashBytes[i][j] {
				return false
			}
		}
	}
	return true
}
