package trie_tst

type Trie struct {
	val      interface{}
	children [256]*Trie
}

func (t *Trie) Set(s string, val interface{}) {
	sl := len(s)
	var set func(st *Trie, idx int)
	set = func(st *Trie, idx int) {
		if idx < sl {
			c := s[idx]
			child := st.children[c]
			if child == nil {
				child = &Trie{}
				st.children[c] = child
			}
			set(child, idx+1)
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
			c := s[idx]
			child := st.children[c]
			if child == nil {
				child = &Trie{}
				st.children[c] = child
			}
			return get(child, idx+1)
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
			c := s[idx]
			child := st.children[c]
			if child == nil {
				child = &Trie{}
				st.children[c] = child
			}
			return subtree(child, idx+1)
		} else {
			return st
		}
	}
	return subtree(t, 0)
}
