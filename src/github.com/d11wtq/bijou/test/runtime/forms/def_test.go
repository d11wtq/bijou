package forms_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"strings"
	"testing"
)

func TestDefWithNameAndValue(t *testing.T) {
	env := test.FakeEnv()
	form := test.NewList(Symbol("def"), Symbol("foo"), Int(42))

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
	form := test.NewList(Symbol("def"))
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
	form := test.NewList(Symbol("def"), Symbol("foo"))
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
	form := test.NewList(Symbol("def"), Int(42), Int(7))
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
	form := test.NewList(Symbol("def"), Symbol("x"), Int(42))
	env := test.FakeEnv()
	env.Def("x", Int(7))

	_, err := form.Eval(env)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
}

func TestDefWithDocString(t *testing.T) {
	form := test.NewList(
		Symbol("def"),
		Symbol("x"),
		String("doc string"),
		Int(42),
	)

	env := test.FakeEnv()

	v, err := form.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestDefWithValueViaCall(t *testing.T) {
	call := &Call{
		Fn:   test.FakeFn(Int(42)),
		Args: EmptyList,
		Env:  test.FakeEnv(),
	}

	form := test.NewList(Symbol("def"), Symbol("x"), call)

	v, err := form.Eval(test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}
