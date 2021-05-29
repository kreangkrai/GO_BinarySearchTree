package main

import (
	"testing"
)

var tests = []struct {
	input  int
	output bool
}{
	{6, true},
	{16, false},
	{3, true},
}

func TestSearch(t *testing.T) {
	tree := &Node{Key: 6, Left: &Node{Key: 3}}
	for i, test := range tests {
		res := tree.Search(test.input)
		if res != test.output {
			t.Errorf("%d : got %v,expected %v", i, res, test.output)
		}
	}
}
func BenchmarkSearch(b *testing.B) {
	tree := &Node{Key: 6}
	for i := 0; i < b.N; i++ {
		tree.Search(6)
	}
}
