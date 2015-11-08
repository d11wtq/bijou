package core_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestHeadReturnsHeadOfSequence(t *testing.T) {
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

	args = runtime.EmptyList.Cons(runtime.String("foo"))
	v, err = core.Head(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.Int('f') {
		t.Fatalf(`expected v == Int('f'), got %s`, v)
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

func TestTailReturnsTailOfSequence(t *testing.T) {
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

	args = runtime.EmptyList.Cons(runtime.String("foo"))
	v, err = core.Tail(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.String("oo") {
		t.Fatalf(`expected v == String("oo"), got %s`, v)
	}
}

func TestEmptySequence(t *testing.T) {
	args := runtime.EmptyList.Cons(
		runtime.EmptyList.Cons(runtime.Int(42)),
	)
	v, err := core.Empty(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.False {
		t.Fatalf(`expected v == False, got %s`, v)
	}

	args = runtime.EmptyList.Cons(runtime.EmptyList)
	v, err = core.Empty(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}
