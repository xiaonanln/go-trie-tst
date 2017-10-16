package trie_tst

import (
	"testing"

	"github.com/bmizerany/assert"
)

type tree interface {
	Set(s string, val interface{})
	Get(s string) (val interface{})
}

func TestTrie(t *testing.T) {
	var tr Trie
	testTrie(t, &tr)
}

func testTrie(t *testing.T, tr tree) {
	tr.Set("", 0)
	tr.Set("abc", 3)

	assert.Equal(t, tr.Get(""), 0)
	assert.Equal(t, tr.Get("a"), nil)
	assert.Equal(t, tr.Get("ab"), nil)
	assert.Equal(t, tr.Get("abc"), 3)

	var subtr tree
	if trie, ok := tr.(*Trie); ok {
		subtr = trie.SubTree("ab")
	}

	assert.Equal(t, subtr.Get(""), nil)
	assert.Equal(t, subtr.Get("a"), nil)
	assert.Equal(t, subtr.Get("b"), nil)
	assert.Equal(t, subtr.Get("c"), 3)
}
