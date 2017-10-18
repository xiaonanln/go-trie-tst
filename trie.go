package trietst

// Trie can be the root, and can be a sub-tree
type Trie struct {
	Val      interface{}
	children [256]*Trie
}

// Child returns the child subtree of the current tree
func (t *Trie) Child(c byte) *Trie {
	child := t.children[c]
	if child == nil {
		child = &Trie{}
		t.children[c] = child
	}
	return child
}

// Set sets the value of string in the current tree
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

// Get returns the value of string in the current tree
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

// Sub returns the subtree of the current tree with specified prefix
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

func (t *Trie) ForEach(f func(s string, val interface{})) {
	var prefix []byte
	t.forEach(f, prefix)
}

func (t *Trie) forEach(f func(s string, val interface{}), prefix []byte) {
	if t.Val != nil {
		f(string(prefix), t.Val)
	}

	for c, st := range t.children {
		if st != nil {
			st.forEach(f, append(prefix, byte(c)))
		}
	}
}
