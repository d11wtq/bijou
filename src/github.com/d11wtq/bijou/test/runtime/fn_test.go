package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"strings"
	"testing"
)

func TestFnType(t *testing.T) {
	fn := NewFn(EmptyList, EmptyList, test.FakeEnv())
	if fn.Type() != FnType {
		t.Fatalf(`expected fn.Type() == FnType, got %s`, fn.Type())
	}
}

func TestFnEvalToSelf(t *testing.T) {
	fn := NewFn(EmptyList, EmptyList, test.FakeEnv())

	v, err := fn.Eval(test.FakeEnv())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != fn {
		t.Fatalf(`expected v == fn, got %s`, v)
	}
}

func TestFnApplyReturnsLastEvaluatedExpression(t *testing.T) {
	params := EmptyList
	body := EmptyList.Cons(Int(42)).Cons(Int(7))
	fn := NewFn(params, body, test.FakeEnv())

	v, err := fn.Apply(EmptyList)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestFnApplyUsesClosedEnvironment(t *testing.T) {
	params := EmptyList
	body := EmptyList.Cons(Symbol("foo")).Cons(Int(7))
	env := test.FakeEnv()
	env.Def("foo", Int(99))
	fn := NewFn(params, body, env)

	v, err := fn.Apply(EmptyList)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(99) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestFnApplyExtendsEnvironmentWithArgs(t *testing.T) {
	params := EmptyList.Cons(Symbol("x"))
	body := EmptyList.Cons(Symbol("x")).Cons(Int(7))
	env := test.FakeEnv()
	env.Def("foo", Int(99))
	fn := NewFn(params, body, env)
	args := EmptyList.Cons(Int(21))

	v, err := fn.Apply(args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(21) {
		t.Fatalf(`expected v == Int(21), got %s`, v)
	}
}

func TestFnApplyValidatesTooFewArgs(t *testing.T) {
	params := EmptyList.Cons(Symbol("y")).Cons(Symbol("x"))
	body := EmptyList
	env := test.FakeEnv()
	fn := NewFn(params, body, env)
	args := EmptyList.Cons(Int(21))

	v, err := fn.Apply(args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	if _, ok := err.(*ArgumentError); !ok {
		t.Fatalf(`expected err.(*ArgumentError), got %s`, err)
	}

	match := "wanted 2, got 1"
	if !strings.Contains(strings.ToLower(err.Error()), match) {
		t.Fatalf(`expected err to match "%s", got %s`, match, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestFnApplyValidatesTooManyArgs(t *testing.T) {
	params := EmptyList.Cons(Symbol("y")).Cons(Symbol("x"))
	body := EmptyList
	env := test.FakeEnv()
	fn := NewFn(params, body, env)
	args := EmptyList.Cons(Int(21)).Cons(Int(9)).Cons(Int(2))

	v, err := fn.Apply(args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	if _, ok := err.(*ArgumentError); !ok {
		t.Fatalf(`expected err.(*ArgumentError), got %s`, err)
	}

	match := "wanted 2, got 3"
	if !strings.Contains(strings.ToLower(err.Error()), match) {
		t.Fatalf(`expected err to match "%s", got %s`, match, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}
