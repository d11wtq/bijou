package core_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestListReturnsVariadicArgs(t *testing.T) {
	args := (&runtime.List{}).Append(runtime.Int(42)).Append(runtime.Int(7))
	v, err := core.List(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if !v.Eq(args) {
		t.Fatalf(`expected v == args, got %s`, v)
	}
}

func TestConsReturnsAConsCell(t *testing.T) {
	args := (&runtime.List{}).
		Append(runtime.Int(7)).
		Append(runtime.Cons(runtime.Int(42), runtime.Nil))

	v, err := core.Cons(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	v2, ok := v.(*runtime.ConsCell)
	if ok == false {
		t.Fatalf(`expected v.(*ConsCell), but not a *ConsCell`)
	}

	if v2.Head() != runtime.Int(7) {
		t.Fatalf(`expected v2.Head() == Int(7), got %s`, v2.Head())
	}
	if v2.Tail().Head() != runtime.Int(42) {
		t.Fatalf(
			`expected v2.Tail().Head() == Int(42), got %s`,
			v2.Tail().Head(),
		)
	}
}
