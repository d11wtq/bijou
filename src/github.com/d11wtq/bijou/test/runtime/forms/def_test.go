package forms_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"strings"
	"testing"
)

func TestDefWithNameAndValue(t *testing.T) {
	env := test.FakeEnv()
	form := EmptyList.Cons(Int(42)).Cons(Symbol("foo")).Cons(Symbol("def"))

	v, err := form.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}

	v2, _ := env.Get("foo")
	if v2 != Int(42) {
		t.Fatalf(`expected env.Get("foo") == Int(42), got %s`, v2)
	}
}

func TestDefWithoutName(t *testing.T) {
	form := EmptyList.Cons(Symbol("def"))
	msg := "missing name"

	v, err := form.Eval(test.FakeEnv())
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if !strings.Contains(strings.ToLower(err.Error()), msg) {
		t.Fatalf(`expected err to match "%s", got: %s`, msg, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestDefWithoutValue(t *testing.T) {
	form := EmptyList.Cons(Symbol("foo")).Cons(Symbol("def"))
	msg := "missing value"

	v, err := form.Eval(test.FakeEnv())
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if !strings.Contains(strings.ToLower(err.Error()), msg) {
		t.Fatalf(`expected err to match "%s", got: %s`, msg, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestDefWithNonSymbolName(t *testing.T) {
	form := EmptyList.Cons(Int(7)).Cons(Int(42)).Cons(Symbol("def"))
	msg := "symbol"

	v, err := form.Eval(test.FakeEnv())
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if !strings.Contains(strings.ToLower(err.Error()), msg) {
		t.Fatalf(`expected err to match "%s", got: %s`, msg, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestDefWithDoubleDefine(t *testing.T) {
	form := EmptyList.Cons(Int(42)).Cons(Symbol("x")).Cons(Symbol("def"))
	env := test.FakeEnv()
	env.Def("x", Int(7))

	_, err := form.Eval(env)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
}
