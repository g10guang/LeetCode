package bitmap

type bitmap struct {
	words []uint64
	cnt   int
}

func (b *bitmap) add(v int32) bool {
	index, bit := b.calIndex(v)
	for uint32(len(b.words)) <= index {
		b.words = append(b.words, 0)
	}
	if b.words[index]&(1<<bit) > 0 {
		return false
	}
	b.words[index] |= 1 << uint8(bit)
	return true
}

func (b *bitmap) isIn(v int32) bool {
	index, bit := b.calIndex(v)
	if uint32(len(b.words)) <= index {
		return false
	}
	return b.words[index]&(1<<bit) > 0
}

func (b *bitmap) remove(v int32) (ret bool) {
	index, bit := b.calIndex(v)
	if uint32(len(b.words)) <= index {
		return false
	}
	if b.words[index]&(1<<uint8(bit)) > 0 {
		b.words[index] &= ^(1 << bit)
		ret = true
	}
	// Recycle the storage.
	for last := len(b.words) - 1; last > 0 && b.words[last] == 0; last-- {
		b.words = b.words[:last]
	}
	return
}

// Union the bitmap
func (b *bitmap) union(o *bitmap) {
	for len(b.words) < len(o.words) {
		b.words = append(b.words, 0)
	}
	for i := 0; i < len(o.words); i++ {
		b.words[i] |= o.words[i]
	}
}

func (b *bitmap) intersection(o *bitmap) {
	for i := 0; i < len(o.words); i++ {
		b.words[i] &= o.words[i]
	}
}

func (b *bitmap) difference(o *bitmap) {
	for len(b.words) < len(o.words) {
		b.words = append(b.words, 0)
	}
	for i := 0; i < len(b.words); i++ {
		b.words[i] ^= o.words[i]
	}
}

func (b *bitmap) calIndex(v int32) (uint32, uint8) {
	t := uint32(v)
	return t / 64, uint8(t % 64)
}
