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

func TestSymbolEvalUndefined(t *testing.T) {
	sym := Symbol("test")
	env := FakeEnv()
	if v, err := sym.Eval(env); err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	} else if !strings.Contains(strings.ToLower(err.Error()), "undefined") {
		t.Fatalf(`expected err to match "undefined", got: %s`, err)
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

func TestSymbolCastToNative(t *testing.T) {
	sym := Symbol("test")
	if string(sym) != "test" {
		t.Fatalf(`expected string(sym) == "test", got %s`, string(sym))
	}
}
