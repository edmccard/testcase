package testcase

import "testing"

type Testable interface {
	Equals(other Testable) bool
	Diff(other Testable) interface{}
}

type Branch func(Tree, Testable, Testable)

type Tree struct {
	Name     string
	T        *testing.T
	Branches []Branch
}

func (t Tree) Run(name string, init, exp Testable) bool {
	if len(t.Branches) == 0 {
		if !exp.Equals(init) {
			diff := exp.Diff(init)
			if diff != nil {
				t.T.Errorf("%s %s: %s", t.Name, name, diff)
			}
			return false
		}
	} else {
		next := Tree{t.Name + " " + name, t.T, t.Branches[1:]}
		t.Branches[0](next, init, exp)
	}
	return true
}

func (t Tree) Start(init Testable) bool {
	return t.Run("", init, nil)
}

func NewTree(t *testing.T, name string, branches ...Branch) Tree {
	return Tree{name, t, branches}
}
