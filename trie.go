package trietst

type Trie struct {
	Val      interface{}
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
	t.set(s, val, 0)
}

func (t *Trie) set(s string, val interface{}, idx int) {
	if idx < len(s) {
		t.Child(s[idx]).set(s, val, idx+1)
	} else {
		t.Val = val
	}
}

func (t *Trie) Get(s string) (val interface{}) {
	return t.get(s, 0)
}

func (t *Trie) get(s string, idx int) (val interface{}) {
	if idx < len(s) {
		return t.Child(s[idx]).get(s, idx+1)
	} else {
		return t.Val
	}
}

func (t *Trie) Sub(s string) *Trie {
	return t.sub(s, 0)
}

func (t *Trie) sub(s string, idx int) *Trie {
	if idx < len(s) {
		return t.Child(s[idx]).sub(s, idx+1)
	} else {
		return t
	}
}
