package core_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestListReturnsVariadicArgs(t *testing.T) {
	args := runtime.EmptyList.Cons(runtime.Int(42)).Cons(runtime.Int(7))
	v, err := core.List(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != args {
		t.Fatalf(`expected v == args, got %s`, v)
	}
}

func TestHeadReturnsHeadOfList(t *testing.T) {
	args := runtime.EmptyList.Cons(
		runtime.EmptyList.Cons(runtime.Int(42)).Cons(runtime.Int(7)),
	)
	v, err := core.Head(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.Int(7) {
		t.Fatalf(`expected v == Int(7), got %s`, v)
	}
}

func TestHeadValidatesArity(t *testing.T) {
	v, err := core.Head(test.FakeEnv(), runtime.EmptyList)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
	args := runtime.EmptyList.Cons(runtime.EmptyList).Cons(runtime.EmptyList)
	v, err = core.Head(test.FakeEnv(), args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestTailReturnsTailOfList(t *testing.T) {
	args := runtime.EmptyList.Cons(
		runtime.EmptyList.Cons(runtime.Int(42)).Cons(runtime.Int(7)),
	)
	v, err := core.Tail(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	v2 := v.(*runtime.List).Data
	if v2 != runtime.Int(42) {
		t.Fatalf(`expected v2 == Int(42), got %s`, v2)
	}
}

func TestConsReturnsANewList(t *testing.T) {
	args := runtime.EmptyList.
		Cons(runtime.EmptyList.Cons(runtime.Int(42))).
		Cons(runtime.Int(7))
	v, err := core.Cons(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	v2 := v.(*runtime.List)
	if v2.Data != runtime.Int(7) {
		t.Fatalf(`expected v2.Data == Int(7), got %s`, v2.Data)
	}
	if v2.Next.Data != runtime.Int(42) {
		t.Fatalf(`expected v2.Next.Data == Int(42), got %s`, v2.Next.Data)
	}
}
