package core_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestListReturnsVariadicArgs(t *testing.T) {
	args := runtime.EmptyList.Cons(runtime.Int(42)).Cons(runtime.Int(7))
	v, err := core.List(args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != args {
		t.Fatalf(`expected v == args, got %s`, v)
	}
}
