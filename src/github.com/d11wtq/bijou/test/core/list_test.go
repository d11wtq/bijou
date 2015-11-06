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
