package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestBindMatchSuccessWithInts(t *testing.T) {
	pattern, value := Int(42), Int(42)
	env := NewScope(nil)
	err := Bind(pattern, value, env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
}

func TestBindMatchErrorWithInts(t *testing.T) {
	pattern, value := Int(42), Int(43)
	env := NewScope(nil)
	err := Bind(pattern, value, env)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
}

func TestBindMatchSuccessWithSymbolUnbound(t *testing.T) {
	pattern, value := Symbol("x"), Int(42)
	env := NewScope(nil)
	err := Bind(pattern, value, env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	def, ok := env.Get("x")
	if ok == false {
		t.Fatalf(`expected env.Get("x"), but not defined`)
	}
	if def != Int(42) {
		t.Fatalf(`expected def == Int(42), got %s`, def)
	}
}

func TestBindMatchSuccessWithSymbolBound(t *testing.T) {
	pattern, value := Symbol("x"), Int(42)
	env := NewScope(nil)
	env.Def("x", Int(42))

	err := Bind(pattern, value, env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	def, ok := env.Get("x")
	if ok == false {
		t.Fatalf(`expected env.Get("x"), but not defined`)
	}
	if def != Int(42) {
		t.Fatalf(`expected def == Int(42), got %s`, def)
	}
}

func TestBindMatchErrorWithSymbolBound(t *testing.T) {
	pattern, value := Symbol("x"), Int(43)
	env := NewScope(nil)
	env.Def("x", Int(42))

	err := Bind(pattern, value, env)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	def, ok := env.Get("x")
	if ok == false {
		t.Fatalf(`expected env.Get("x"), but not defined`)
	}
	if def != Int(42) {
		t.Fatalf(`expected def == Int(42), got %s`, def)
	}
}
