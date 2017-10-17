package main

import (
	"net/http"
	_ "net/http/pprof"

	"time"

	"math/rand"

	"fmt"

	"os"

	"github.com/xiaonanln/go-trie-tst"
)

func testRoutine() {
	trie := &trie_tst.Trie{}
	tst := &trie_tst.TST{}

	for {
		slen := 5 + rand.Intn(10)
		s := randString(slen)
		fmt.Fprintf(os.Stderr, "|%s", s)
		trie.Set(s, struct{}{})
		tst.Set(s, struct{}{})
		time.Sleep(time.Millisecond)
	}
}

func randString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = 'a' + byte(rand.Intn(26))
	}
	return string(b)
}

func main() {
	go testRoutine()
	http.ListenAndServe("localhost:8080", nil)
}
