package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"strings"
	"testing"
)

func TestFuncType(t *testing.T) {
	fn := &Func{
		Params: EmptyList,
		Body:   EmptyList,
		Env:    test.FakeEnv(),
	}

	if fn.Type() != FuncType {
		t.Fatalf(`expected fn.Type() == FuncType, got %s`, fn.Type())
	}
}

func TestFuncEvalToSelf(t *testing.T) {
	fn := &Func{
		Params: EmptyList,
		Body:   EmptyList,
		Env:    test.FakeEnv(),
	}

	v, err := fn.Eval(test.FakeEnv())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != fn {
		t.Fatalf(`expected v == fn, got %s`, v)
	}
}

func TestFuncCallReturnsLastEvaluatedExpression(t *testing.T) {
	params := EmptyList
	body := EmptyList.Cons(Int(42)).Cons(Int(7))
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    test.FakeEnv(),
	}

	v, err := fn.Call(EmptyList)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestFuncCallUsesClosedEnvironment(t *testing.T) {
	params := EmptyList
	body := EmptyList.Cons(Symbol("foo")).Cons(Int(7))
	env := test.FakeEnv()
	env.Def("foo", Int(99))
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    env,
	}

	v, err := fn.Call(EmptyList)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(99) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestFuncCallExtendsEnvironmentWithArgs(t *testing.T) {
	params := EmptyList.Cons(Symbol("x"))
	body := EmptyList.Cons(Symbol("x")).Cons(Int(7))
	env := test.FakeEnv()
	env.Def("foo", Int(99))
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := EmptyList.Cons(Int(21))

	v, err := fn.Call(args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(21) {
		t.Fatalf(`expected v == Int(21), got %s`, v)
	}
}

func TestFuncCallValidatesTooFewArgs(t *testing.T) {
	params := EmptyList.Cons(Symbol("y")).Cons(Symbol("x"))
	body := EmptyList
	env := test.FakeEnv()
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := EmptyList.Cons(Int(21))

	v, err := fn.Call(args)
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

func TestFuncCallValidatesTooManyArgs(t *testing.T) {
	params := EmptyList.Cons(Symbol("y")).Cons(Symbol("x"))
	body := EmptyList
	env := test.FakeEnv()
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := EmptyList.Cons(Int(21)).Cons(Int(9)).Cons(Int(2))

	v, err := fn.Call(args)
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

func TestFuncCallShortCirtcuitsOnError(t *testing.T) {
	v1 := test.NewFakeValue(Symbol("xx"))
	v2 := test.NewFakeValue(Symbol("yy"))

	params := EmptyList.Cons(Symbol("y")).Cons(Symbol("x"))
	body := EmptyList.Cons(v2).Cons(v1)
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    test.FakeEnv(),
	}
	args := EmptyList.Cons(Int(21)).Cons(Int(9))

	v, err := fn.Call(args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if err != v1.Error {
		t.Fatalf(`expected err == v1.Error, got %s`, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}

	if v2.Evaluated {
		t.Fatalf(`expected v2.Evaluated == false, got true`)
	}
}
