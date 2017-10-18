package trietst

// TST can be the root, and can be a sub-tree
type TST struct {
	left  *TST
	right *TST
	eq    *TST
	eqkey byte
	Val   interface{}
}

// Child returns the child subtree of the current tree
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

// Set sets the value of string in the current tree
func (t *TST) Set(s string, val interface{}) {
	t.set(s, val, 0)
}

func (t *TST) set(s string, val interface{}, idx int) {
	if idx < len(s) {
		t.Child(s[idx]).set(s, val, idx+1)
	} else {
		t.Val = val
	}
}

// Get returns the value of string in the current tree
func (t *TST) Get(s string) (val interface{}) {
	return t.get(s, 0)
}

func (t *TST) get(s string, idx int) (val interface{}) {
	if idx < len(s) {
		return t.Child(s[idx]).get(s, idx+1)
	} else {
		return t.Val
	}
}

// Sub returns the subtree of the current tree with specified prefix
func (t *TST) Sub(s string) *TST {
	return t.sub(s, 0)
}

func (t *TST) sub(s string, idx int) *TST {
	if idx < len(s) {
		return t.Child(s[idx]).sub(s, idx+1)
	} else {
		return t
	}
}

func (t *TST) ForEach(f func(s string, val interface{})) {
	var prefix []byte
	t.forEach(f, prefix)
}

func (t *TST) forEach(f func(s string, val interface{}), prefix []byte) {
	if t.Val != nil {
		f(string(prefix), t.Val)
	}

	if t.left != nil {
		t.left.forEach(f, prefix)
	}

	if t.eq != nil {
		t.eq.forEach(f, append(prefix, t.eqkey))
	}

	if t.right != nil {
		t.right.forEach(f, prefix)
	}
}
