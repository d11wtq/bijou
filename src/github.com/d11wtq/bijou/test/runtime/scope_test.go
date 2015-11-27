package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestScopeGetUnbound(t *testing.T) {
	env := NewScope(nil)
	v, ok := env.Get("test")
	if ok == true {
		t.Fatalf(`expected ok == false, got true`)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestScopeResolveUnbound(t *testing.T) {
	env := NewScope(nil)
	v, ok := env.Resolve("test")
	if ok == true {
		t.Fatalf(`expected ok == false, got true`)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestScopeDefAndGet(t *testing.T) {
	env := NewScope(nil)
	err := env.Def("test", Symbol("example"))
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	v, ok := env.Get("test")
	if ok == false {
		t.Fatalf(`expected ok == true, got false`)
	}
	if v != Symbol("example") {
		t.Fatalf(`expected v == Symbol("example"), got %s`, v)
	}
}

func TestScopeDefAndResolve(t *testing.T) {
	env := NewScope(nil)
	err := env.Def("test", Symbol("example"))
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	v, ok := env.Resolve("test")
	if ok == false {
		t.Fatalf(`expected ok == true, got false`)
	}
	if v != Symbol("example") {
		t.Fatalf(`expected v == Symbol("example"), got %s`, v)
	}
}

func TestScopeDefAndGetNonRecursing(t *testing.T) {
	parent := NewScope(nil)
	parent.Def("test", Symbol("example"))
	env := NewScope(parent)
	v, ok := env.Get("test")
	if ok == true {
		t.Fatalf(`expected ok == false, got true`)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestScopeDefAndResolveViaParent(t *testing.T) {
	parent := NewScope(nil)
	parent.Def("test", Symbol("example"))
	env := NewScope(parent)
	v, ok := env.Resolve("test")
	if ok == false {
		t.Fatalf(`expected ok == true, got false`)
	}
	if v != Symbol("example") {
		t.Fatalf(`expected v == Symbol("example"), got %s`, v)
	}
}

func TestScopeDefAndResolveExtendMasking(t *testing.T) {
	parent := NewScope(nil)
	parent.Def("test", Symbol("example"))
	env := parent.Extend()
	err := env.Def("test", Symbol("other"))
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	v, ok := env.Resolve("test")
	if ok == false {
		t.Fatalf(`expected ok == true, got false`)
	}
	if v != Symbol("other") {
		t.Fatalf(`expected v == Symbol("other"), got %s`, v)
	}
}

func TestScopeDefAndResolveExtendImmutable(t *testing.T) {
	parent := NewScope(nil)
	parent.Def("test", Symbol("example"))
	env := parent.Extend()
	env.Def("test", Symbol("other"))

	v, ok := parent.Resolve("test")
	if ok == false {
		t.Fatalf(`expected ok == true, got false`)
	}
	if v != Symbol("example") {
		t.Fatalf(`expected v == Symbol("example"), got %s`, v)
	}
}

func TestScopeDefTwiceInTheSameScope(t *testing.T) {
	var err error
	env := NewScope(nil)
	err = env.Def("x", Int(42))
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	err = env.Def("x", Int(7))
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	err = env.Def("x", Int(42))
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
}
