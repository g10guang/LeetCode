package tree

import (
	"unicode/utf8"
)

// prefix tree.
// Application scenario:
// 1 the most like prefix
// 2 save storage
// 3 word spelling check


// note: utf8.DecodeRuneInString(word)
// if cannot extract a utf-8 rune in word then size == 0

type trie struct {
	//root map[rune]*point
	root *point
}

type point struct {
	children map[rune]*point
	// check is word or not.
	isWord bool
	// Maintain the reference count.
	// In remove function if cnt == 0 remove current point.
	cnt uint
}

type tries interface {
	// add a word into tries tree
	// if word has existed, return false
	// else return true
	add(word string) bool

	// check a word is whether in tries or not.
	isIn(word string) bool

	// remove a word from tries
	// return word in tries tree or not before remove.
	// return any remove statement run or not.
	remove(word string) bool

	// return the max prefix of the word.
	maxPrefix(word string) string

	init()
}

func (t *trie) init() {
	t.root = &point{children: make(map[rune]*point), isWord: false, cnt: 0}
}

func (t *trie) add(word string) bool {
	// here make sure t.root will not be the leaf.
	if len(word) == 0 {
		return false
	}
	return t.root.add(word)
}

func (t *trie) isIn(word string) bool {
	if len(word) == 0 {
		return false
	}
	return t.root.isIn(word)
}

func (t *trie) remove(word string) bool {
	if len(word) == 0 {
		return false
	}
	return t.root.remove(word)
}

func (t *trie) maxPrefix(word string) string {
	return word[:t.root.maxPrefix(word, 0)]
}

func (p *point) maxPrefix(word string, matchPrefixIndex int) int {
	if len(word) == 0 {
		// return the history match result
		return matchPrefixIndex
	}
	r, size := utf8.DecodeRuneInString(word)
	p.checkSize(size)
	if v, ok := p.children[r]; ok {
		return v.maxPrefix(word[size:], matchPrefixIndex+size)
	}
	return matchPrefixIndex
}

func (p *point) remove(word string) bool {
	if len(word) == 0 {
		if p.isWord {
			// remove word. if cnt == 0 remove current point statement will call in parent.
			p.isWord = false
			p.cnt--
			return true
		}
		return false
	}
	r, size := utf8.DecodeRuneInString(word)
	p.checkSize(size)
	if v, ok := p.children[r]; ok {
		if v.remove(word[size:]) {
			// reference count decrement.
			// maybe child will be remove.
			p.cnt--
			if v.cnt == 0 {
				// release this node.
				delete(p.children, r)
			}
			// remove successfully and return true.
			return true
		}
	}
	return false
}

func (p *point) isIn(word string) bool {
	if len(word) == 0 {
		return p.isWord
	}
	r, size := utf8.DecodeRuneInString(word)
	p.checkSize(size)
	if v, ok := p.children[r]; ok {
		// recursively find.
		return v.isIn(word[size:])
	}
	return false
}

func (p *point) add(word string) bool {
	if len(word) == 0 {
		// Add a new word successfully.
		if p.isWord {
			// return false the upper parent will not increment cnt.
			return false
		}
		p.isWord = true
		p.cnt += 1
		return true
	}
	r, size := utf8.DecodeRuneInString(word)
	p.checkSize(size)
	if _, ok := p.children[r]; !ok {
		// cnt will be update after the recursive call.
		newPoint := &point{children: make(map[rune]*point), isWord: false, cnt: 0}
		p.children[r] = newPoint
	}
	if p.children[r].add(word[size:]) {
		p.cnt++
		return true
	}
	return false
}

func (p *point) checkSize(size int) {
	if size == 0 {
		panic("The words you input to tries is not utf-8 encoded.")
	}
}

