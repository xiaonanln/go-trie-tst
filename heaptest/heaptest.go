package main

import (
	"net/http"
	_ "net/http/pprof"

	"time"

	"math/rand"

	"github.com/xiaonanln/go-trie-tst"
)

func testRoutine() {
	trie := &trietst.Trie{}
	tst := &trietst.TST{}
	triemo := &trietst.TrieMO{}

	counter := 0
	for {
		slen := 5 + rand.Intn(10)
		s := randString(slen)
		trie.Set(s, struct{}{})
		tst.Set(s, struct{}{})
		triemo.Set(s, struct{}{})

		_ = trie.Get(s)
		_ = tst.Get(s)
		_ = triemo.Get(s)

		counter += 1
		if counter%10000 == 0 {
			println(counter)
		}
		time.Sleep(time.Microsecond * 100)
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
