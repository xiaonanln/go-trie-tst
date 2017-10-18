# go-trie-tst
[![travis](https://travis-ci.org/xiaonanln/go-trie-tst.svg?branch=master)](https://travis-ci.org/xiaonanln/go-trie-tst)
[![GoCover](http://gocover.io/_badge/github.com/xiaonanln/go-trie-tst)](https://gocover.io/github.com/xiaonanln/go-trie-tst)
[![GoReportCard](https://goreportcard.com/badge/github.com/xiaonanln/go-trie-tst)](https://goreportcard.com/report/github.com/xiaonanln/go-trie-tst)

Trie and Ternary Search Tree implemented in Golang

Trie outperforms TST slightly in CPU time, but costs much much more memory according to heap pprof.
So I think TST is more suitable in production.

## Performance
**CPU Profile:**
```
Trie    200000	      9964 ns/op
TST     200000	     10339 ns/op
```

**Heap Profile:**
```
(pprof) top50
1221.12MB of 1222.12MB total (99.92%)
Dropped 42 nodes (cum <= 6.11MB)
      flat  flat%   sum%        cum   cum%
 1193.62MB 97.67% 97.67%  1193.62MB 97.67%  github.com/xiaonanln/go-trie-tst.(*Trie).Set.func1
   27.50MB  2.25% 99.92%    27.50MB  2.25%  github.com/xiaonanln/go-trie-tst.(*TST).Child
         0     0% 99.92%    27.50MB  2.25%  github.com/xiaonanln/go-trie-tst.(*TST).Set
         0     0% 99.92%    27.50MB  2.25%  github.com/xiaonanln/go-trie-tst.(*TST).Set.func1
         0     0% 99.92%  1193.62MB 97.67%  github.com/xiaonanln/go-trie-tst.(*Trie).Set
         0     0% 99.92%  1221.12MB 99.92%  main.testRoutine
         0     0% 99.92%  1222.12MB   100%  runtime.goexit
```

## Usage 

```go
import "github.com/xiaonanln/go-trie-tst"

var tr trietst.TST // create a TST
tr.Set("", 0) // set "" to 0
tr.Set("abc", 3) // set "abc" to 3
tr.Get("") // == 0
tr.Get("a") // == nil
tr.Get("ab") // == nil
tr.Get("abc") // == 3

subtr = tr.Sub("ab") // get sub TST under "ab"
subtr.Get("") // == nil
subtr.Get("a") // == nil
subtr.Get("b") // == nil
subtr.Get("c") // == 3, because "abc" == 3
```
