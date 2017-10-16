package trie_tst

type TST struct {
	left  *TST
	right *TST
	eq    *TST
	eqkey byte
	val   interface{}
}

func (t *TST) Child(c byte) *TST {
	if t.eq == nil {
		t.eqkey = c
		t.eq = &TST{}
		return t.eq
	} else if c == t.eqkey {
		return t.eq
	} else if c < t.eqkey {
		if t.left == nil {
			t.left = &TST{}
		}
		return t.left.Child(c)
	} else { // c > t.eqkey
		if t.right == nil {
			t.right = &TST{}
		}
		return t.right.Child(c)
	}
}

func (t *TST) Set(s string, val interface{}) {
	sl := len(s)
	var set func(st *TST, idx int)
	set = func(st *TST, idx int) {
		if idx < sl {
			set(st.Child(s[idx]), idx+1)
		} else {
			st.val = val
		}
	}
	set(t, 0)
}

func (t *TST) Get(s string) (val interface{}) {
	sl := len(s)
	var get func(st *TST, idx int) (val interface{})
	get = func(st *TST, idx int) interface{} {
		if idx < sl {
			return get(st.Child(s[idx]), idx+1)
		} else {
			return st.val
		}
	}
	return get(t, 0)
}

func (t *TST) SubTree(s string) *TST {
	sl := len(s)
	var subtree func(st *TST, idx int) *TST
	subtree = func(st *TST, idx int) *TST {
		if idx < sl {
			return subtree(st.Child(s[idx]), idx+1)
		} else {
			return st
		}
	}
	return subtree(t, 0)
}
