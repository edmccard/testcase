package testcase

import "testing"

type Testable interface {
	Equals(other Testable) bool
	Diff(other Testable) interface{}
}

type Branch func(Tree, Testable)

type Tree struct {
	Name     string
	T        *testing.T
	Branches []Branch
	Exp      Testable
}

func (t Tree) Run(name string, init Testable) bool {
	if len(t.Branches) == 0 {
		if !t.Exp.Equals(init) {
			diff := t.Exp.Diff(init)
			if diff != nil {
				t.T.Errorf("%s %s: %s", t.Name, name, diff)
			}
			return false
		}
	} else {
		next := Tree{t.Name + " " + name, t.T, t.Branches[1:], t.Exp}
		t.Branches[0](next, init)
	}
	return true
}
