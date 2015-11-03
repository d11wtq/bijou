package core_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"testing"
)

func example(args *runtime.List) (runtime.Value, error) {
	return args, nil
}

func TestGoFunc(t *testing.T) {
	fn := core.GoFunc(example)

	if fn.Type() != runtime.FuncType {
		t.Fatalf(`expected fn.Type() == FuncType, got %s`, fn.Type())
	}
	if _, ok := fn.(runtime.Callable); ok == false {
		t.Fatalf(`expected fn.(Callable), got false`)
	}

	args := runtime.EmptyList.Cons(runtime.Int(42))
	v2, err2 := fn.Call(args)
	if err2 != nil {
		t.Fatalf(`expected err2 == nil, got %s`, err2)
	}
	if v2 != args {
		t.Fatalf(`expected v2 == args, got %s`, v2)
	}
}

func TestReadArgsWithCorrectArity(t *testing.T) {
	var a, b runtime.Value

	args := runtime.EmptyList.Cons(runtime.Int(7)).Cons(runtime.Int(42))

	if err := core.ReadArgs(args, &a, &b); err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if a != runtime.Int(42) {
		t.Fatalf(`expected a == Int(42), got %s`, a)
	}

	if b != runtime.Int(7) {
		t.Fatalf(`expected b == Int(7), got %s`, b)
	}
}

func TestReadArgsWithBadArity(t *testing.T) {
	var a, b runtime.Value

	args1 := runtime.EmptyList.
		Cons(runtime.Int(7)).
		Cons(runtime.Int(42)).
		Cons(runtime.Int(13))

	args2 := runtime.EmptyList.
		Cons(runtime.Int(7))

	if err := core.ReadArgs(args1, &a, &b); err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	if err := core.ReadArgs(args2, &a, &b); err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
}
