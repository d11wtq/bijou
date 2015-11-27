package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"strings"
	"testing"
)

func TestSymbolType(t *testing.T) {
	sym := Symbol("test")
	if sym.Type() != SymbolType {
		t.Fatalf(`expected sym.Type() == SymbolType, got %s`, sym.Type())
	}
}

func TestSymbolEvalUnbound(t *testing.T) {
	sym := Symbol("test")
	env := FakeEnv()
	if v, err := sym.Eval(env); err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	} else if !strings.Contains(strings.ToLower(err.Error()), "unbound") {
		t.Fatalf(`expected err to match "unbound", got: %s`, err)
	} else if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestSymbolEvalToValue(t *testing.T) {
	sym := Symbol("test")
	env := FakeEnv()
	env.Def("test", Int(42))
	if v, err := sym.Eval(env); err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	} else if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestSymbolBindUnbound(t *testing.T) {
	sym := Symbol("x")
	env := NewScope(nil)
	value := Int(42)
	err := sym.Bind(env, value)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	x, ok := env.Get("x")
	if ok == false {
		t.Fatalf(`expected env.Get("x"), but not bound`)
	}
	if x != Int(42) {
		t.Fatalf(`expected x == Int(42), got %s`, x)
	}
}

func TestSymbolBindSameValue(t *testing.T) {
	sym := Symbol("x")
	env := NewScope(nil)
	env.Def("x", Int(42))
	value := Int(42)
	err := sym.Bind(env, value)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	x, ok := env.Get("x")
	if ok == false {
		t.Fatalf(`expected env.Get("x"), but not bound`)
	}
	if x != Int(42) {
		t.Fatalf(`expected x == Int(42), got %s`, x)
	}
}

func TestSymbolBindWrongValue(t *testing.T) {
	sym := Symbol("x")
	env := NewScope(nil)
	env.Def("x", Int(42))
	value := Int(43)
	err := sym.Bind(env, value)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	x, _ := env.Get("x")
	if x != Int(42) {
		t.Fatalf(`expected x == Int(42), got %s`, x)
	}
}

func TestSymbolShadowParentValue(t *testing.T) {
	sym := Symbol("x")
	parent := NewScope(nil)
	parent.Def("x", Int(42))
	env := NewScope(parent)
	value := Int(43)
	err := sym.Bind(env, value)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	x, ok := env.Get("x")
	if ok == false {
		t.Fatalf(`expected env.Get("x"), but not bound`)
	}
	if x != Int(43) {
		t.Fatalf(`expected x == Int(43), got %s`, x)
	}
}

func TestSymbolEq(t *testing.T) {
	if !Symbol("a").Eq(Symbol("a")) {
		t.Fatalf(`expected Symbol("a").Eq(Symbol("a")), got false`)
	}
	if Symbol("a").Eq(Symbol("b")) {
		t.Fatalf(`expected !Symbol("a").Eq(Symbol("b")), got true`)
	}
}

func TestSymbolGt(t *testing.T) {
	if Symbol("foo").Gt(Symbol("foo")) {
		t.Fatalf(`expected !Symbol("foo").Gt(Symbol("foo")), got true`)
	}
	if !Symbol("foo").Gt(Symbol("bar")) {
		t.Fatalf(`expected Symbol("foo").Gt(Symbol("bar")), got false`)
	}
	if !Symbol("foo").Gt(Symbol("fo")) {
		t.Fatalf(`expected Symbol("foo").Gt(Symbol("fo")), got false`)
	}
	if Symbol("foo").Gt(EmptyList) {
		t.Fatalf(`expected !Symbol("foo").Gt(EmptyList), got true`)
	}
	if !Symbol("41").Gt(Int(42)) {
		t.Fatalf(`expected Symbol("41").Gt(Int(42)), got false`)
	}
}

func TestSymbolLt(t *testing.T) {
	if Symbol("foo").Lt(Symbol("foo")) {
		t.Fatalf(`expected !Symbol("foo").Lt(Symbol("foo")), got true`)
	}
	if !Symbol("bar").Lt(Symbol("foo")) {
		t.Fatalf(`expected Symbol("bar").Lt(Symbol("foo")), got false`)
	}
	if !Symbol("fo").Lt(Symbol("foo")) {
		t.Fatalf(`expected Symbol("fo").Lt(Symbol("foo")), got false`)
	}
	if !Symbol("foo").Lt(EmptyList) {
		t.Fatalf(`expected Symbol("foo").Lt(EmptyList), got false`)
	}
	if Symbol("41").Lt(Int(42)) {
		t.Fatalf(`expected !Symbol("41").Lt(Int(42)), got true`)
	}
}

func TestSymbolCastToNative(t *testing.T) {
	sym := Symbol("test")
	if string(sym) != "test" {
		t.Fatalf(`expected string(sym) == "test", got %s`, string(sym))
	}
}

func TestSymbolString(t *testing.T) {
	sym := Symbol("example")
	if sym.String() != "example" {
		t.Fatalf(`expected sym.String() == "example", got %s`, sym.String())
	}
}
