package trietst

import (
	"testing"

	"math/rand"

	"github.com/bmizerany/assert"
)

type tree interface {
	Set(s string, val interface{})
	Get(s string) (val interface{})
}

func TestTrie(t *testing.T) {
	var tr Trie
	testSearchTree(t, &tr)
}

func TestTST(t *testing.T) {
	var tr TST
	testSearchTree(t, &tr)
}

func TestTrieMO(t *testing.T) {
	var tr TrieMO
	testSearchTree(t, &tr)
}

func testSearchTree(t *testing.T, tr tree) {
	tr.Set("", 0)
	tr.Set("abc", 3)

	assert.Equal(t, tr.Get(""), 0)
	assert.Equal(t, tr.Get("a"), nil)
	assert.Equal(t, tr.Get("ab"), nil)
	assert.Equal(t, tr.Get("abc"), 3)

	var subtr tree
	if trie, ok := tr.(*Trie); ok {
		subtr = trie.Sub("ab")
	} else if tst, ok := tr.(*TST); ok {
		subtr = tst.Sub("ab")
	} else if triemo, ok := tr.(*TrieMO); ok {
		subtr = triemo.Sub("ab")
	}

	assert.Equal(t, subtr.Get(""), nil)
	assert.Equal(t, subtr.Get("a"), nil)
	assert.Equal(t, subtr.Get("b"), nil)
	assert.Equal(t, subtr.Get("c"), 3)

	//testTreeMemoryAlloc(t, tr)
}

//func testTreeMemoryAlloc(t *testing.T, tr tree) {
//	buf := make([]byte, 32)
//	for i := 0; i < 100000; i++ {
//		slen := 5 + rand.Intn(6)
//		rand.Read(buf[:slen])
//		s := string(buf[:slen])
//		tr.Set(s, &[2]map[string]struct{}{})
//	}
//}
//

var (
	benchmarkStrings []string
)

func init() {
	for i := 0; i < 10; i++ {
		b := make([]byte, 10+i*10)
		rand.Read(b)
		benchmarkStrings = append(benchmarkStrings, string(b))
	}
}

func BenchmarkTrie(b *testing.B) {
	benchmarkTree(b, &Trie{})
}

func BenchmarkTST(b *testing.B) {
	benchmarkTree(b, &TST{})
}

func BenchmarkTrieMO(b *testing.B) {
	benchmarkTree(b, &TrieMO{})
}

func benchmarkTree(b *testing.B, tr tree) {
	for i := 0; i < b.N; i++ {
		for _, s := range benchmarkStrings {
			tr.Set(s, i)
			tr.Get(s)
		}
	}
}
