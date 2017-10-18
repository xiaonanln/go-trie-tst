package trietst

// TrieMO can be the root, and can be a sub-tree
type TrieMO struct {
	Val          interface{}
	lessChildren []lessChildrenItem
	moreChildren map[byte]*TrieMO
	//children map[byte]*TrieMO
}

type lessChildrenItem struct {
	key byte
	st  *TrieMO
}

// Child returns the child subtree of the current tree
func (t *TrieMO) Child(c byte) *TrieMO {
	if t.moreChildren == nil {
		// find child in lessChildren
		for _, child := range t.lessChildren {
			if child.key == c {
				return child.st
			}
		}

		if len(t.lessChildren) < 8 {
			child := &TrieMO{}
			t.lessChildren = append(t.lessChildren, lessChildrenItem{c, child})
			return child
		}

		// can not find children in lessChildren, we convert lessChildren to moreChildren
		t.moreChildren = map[byte]*TrieMO{}
		for _, child := range t.lessChildren {
			t.moreChildren[child.key] = child.st
		}
		t.lessChildren = nil
	}

	child := t.moreChildren[c]
	if child == nil {
		child = &TrieMO{}
		t.moreChildren[c] = child
	}
	return child
}

//func (t *TrieMO) Child(c byte) *TrieMO {
//	if t.children == nil {
//		t.children = map[byte]*TrieMO{}
//	}
//
//	child := t.children[c]
//	if child == nil {
//		child = &TrieMO{}
//		t.children[c] = child
//	}
//	return child
//}

// Set sets the value of string in the current tree
func (t *TrieMO) Set(s string, val interface{}) {
	t.set(s, val, 0)
}

func (t *TrieMO) set(s string, val interface{}, idx int) {
	if idx < len(s) {
		t.Child(s[idx]).set(s, val, idx+1)
	} else {
		t.Val = val
	}
}

// Get returns the value of string in the current tree
func (t *TrieMO) Get(s string) (val interface{}) {
	return t.get(s, 0)
}

func (t *TrieMO) get(s string, idx int) (val interface{}) {
	if idx < len(s) {
		return t.Child(s[idx]).get(s, idx+1)
	} else {
		return t.Val
	}
}

// Sub returns the subtree of the current tree with specified prefix
func (t *TrieMO) Sub(s string) *TrieMO {
	return t.sub(s, 0)
}

func (t *TrieMO) sub(s string, idx int) *TrieMO {
	if idx < len(s) {
		return t.Child(s[idx]).sub(s, idx+1)
	} else {
		return t
	}
}

func (t *TrieMO) ForEach(f func(s string, val interface{})) {
	var prefix []byte
	t.forEach(f, prefix)
}

func (t *TrieMO) forEach(f func(s string, val interface{}), prefix []byte) {
	if t.Val != nil {
		f(string(prefix), t.Val)
	}

	if t.moreChildren == nil {
		for _, lc := range t.lessChildren {
			lc.st.forEach(f, append(prefix, lc.key))
		}
	} else {
		for c, st := range t.moreChildren {
			if st != nil {
				st.forEach(f, append(prefix, c))
			}
		}
	}
}
