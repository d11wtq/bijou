package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestScopeGetUndefined(t *testing.T) {
	env := NewScope(nil)
	if v, ok := env.Get("test"); ok {
		t.Fatalf(`expected ok == false, got true`)
	} else if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestScopeDefAndGet(t *testing.T) {
	env := NewScope(nil)
	env.Def("test", Symbol("example"))
	if v, ok := env.Get("test"); ok == false {
		t.Fatalf(`expected ok == true, got false`)
	} else if v != Symbol("example") {
		t.Fatalf(`expected v == Symbol("example"), got %s`, v)
	}
}

func TestScopeDefAndGetViaParent(t *testing.T) {
	parent := NewScope(nil)
	parent.Def("test", Symbol("example"))
	env := NewScope(parent)
	if v, ok := env.Get("test"); ok == false {
		t.Fatalf(`expected ok == true, got false`)
	} else if v != Symbol("example") {
		t.Fatalf(`expected v == Symbol("example"), got %s`, v)
	}
}

func TestScopeDefAndGetExtendMasking(t *testing.T) {
	parent := NewScope(nil)
	parent.Def("test", Symbol("example"))
	env := parent.Extend()
	env.Def("test", Symbol("other"))

	if v, ok := env.Get("test"); ok == false {
		t.Fatalf(`expected ok == true, got false`)
	} else if v != Symbol("other") {
		t.Fatalf(`expected v == Symbol("other"), got %s`, v)
	}
}

func TestScopeDefAndGetExtendImmutable(t *testing.T) {
	parent := NewScope(nil)
	parent.Def("test", Symbol("example"))
	env := parent.Extend()
	env.Def("test", Symbol("other"))

	if v, ok := parent.Get("test"); ok == false {
		t.Fatalf(`expected ok == true, got false`)
	} else if v != Symbol("example") {
		t.Fatalf(`expected v == Symbol("example"), got %s`, v)
	}
}
