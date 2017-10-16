package trie_tst

type Trie struct {
	val      interface{}
	children [256]*Trie
}

func (t *Trie) Child(c byte) *Trie {
	child := t.children[c]
	if child == nil {
		child = &Trie{}
		t.children[c] = child
	}
	return child
}

func (t *Trie) Set(s string, val interface{}) {
	sl := len(s)
	var set func(st *Trie, idx int)
	set = func(st *Trie, idx int) {
		if idx < sl {
			set(st.Child(s[idx]), idx+1)
		} else {
			st.val = val
		}
	}
	set(t, 0)
}

func (t *Trie) Get(s string) (val interface{}) {
	sl := len(s)
	var get func(st *Trie, idx int) (val interface{})
	get = func(st *Trie, idx int) interface{} {
		if idx < sl {
			return get(st.Child(s[idx]), idx+1)
		} else {
			return st.val
		}
	}
	return get(t, 0)
}

func (t *Trie) SubTree(s string) *Trie {
	sl := len(s)
	var subtree func(st *Trie, idx int) *Trie
	subtree = func(st *Trie, idx int) *Trie {
		if idx < sl {
			return subtree(st.Child(s[idx]), idx+1)
		} else {
			return st
		}
	}
	return subtree(t, 0)
}
