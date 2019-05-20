package main

import (
	"bytes"
	"fmt"

	"github.com/brian-yu/learn/book/ch2/popcount"
)

// An IntSet is a set of small non-negative integers
// Its zero value represents the empty set
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// Len returns the number of elements
func (s *IntSet) Len() int {
	l := 0
	for _, word := range s.words {
		l += popcount.PopCount(word)
	}
	return l
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &^= 1 << bit
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = make([]uint64, s.Len())
}

// Elems returns a []int containing all elements of the set
func (s *IntSet) Elems() (elems []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for bit := 0; bit < 64; bit++ {
			if word&(1<<uint(bit)) != 0 {
				elems = append(elems, 64*i+bit)
			}
		}
	}
	return
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	t := IntSet{
		words: make([]uint64, len(s.words)),
	}
	copy(t.words, s.words)
	return &t
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String(), x.Len()) // "{1 9 144}"
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String(), y.Len()) // "{9 42}"
	x.UnionWith(&y)
	fmt.Println(x.String(), x.Len())  // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	x.Remove(9)
	fmt.Println(x.String(), x.Len())  // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	z := x.Copy()
	x.Remove(144)
	fmt.Println(z.String())
	fmt.Println(x.String())
	fmt.Println(z.Elems())

	z = &x
	x.Clear()
	fmt.Println(z.String())
	fmt.Println(x.String())
}
