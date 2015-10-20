package runtime

import "testing"

func TestScopeGetUndefined(t *testing.T) {
	env := NewScope(nil)
	if v, ok := env.Get(Symbol("test")); ok {
		t.Fatalf(`expected ok == false, got true`)
	} else if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestScopeDefAndGet(t *testing.T) {
	env := NewScope(nil)
	env.Def(Symbol("test"), Symbol("example"))
	if v, ok := env.Get(Symbol("test")); ok == false {
		t.Fatalf(`expected ok == true, got false`)
	} else if v != Symbol("example") {
		t.Fatalf(`expected v == Symbol("example"), got %s`, v)
	}
}

func TestScopeDefAndGetViaParent(t *testing.T) {
	parent := NewScope(nil)
	parent.Def(Symbol("test"), Symbol("example"))
	env := NewScope(parent)
	if v, ok := env.Get(Symbol("test")); ok == false {
		t.Fatalf(`expected ok == true, got false`)
	} else if v != Symbol("example") {
		t.Fatalf(`expected v == Symbol("example"), got %s`, v)
	}
}

func TestScopeDefAndGetExtendMasking(t *testing.T) {
	parent := NewScope(nil)
	parent.Def(Symbol("test"), Symbol("example"))
	env := parent.Extend()
	env.Def(Symbol("test"), Symbol("other"))

	if v, ok := env.Get(Symbol("test")); ok == false {
		t.Fatalf(`expected ok == true, got false`)
	} else if v != Symbol("other") {
		t.Fatalf(`expected v == Symbol("other"), got %s`, v)
	}
}

func TestScopeDefAndGetExtendImmutable(t *testing.T) {
	parent := NewScope(nil)
	parent.Def(Symbol("test"), Symbol("example"))
	env := parent.Extend()
	env.Def(Symbol("test"), Symbol("other"))

	if v, ok := parent.Get(Symbol("test")); ok == false {
		t.Fatalf(`expected ok == true, got false`)
	} else if v != Symbol("example") {
		t.Fatalf(`expected v == Symbol("example"), got %s`, v)
	}
}
