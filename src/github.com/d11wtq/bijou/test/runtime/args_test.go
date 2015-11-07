package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestReadArgsWithCorrectArity(t *testing.T) {
	var a, b Value

	args := EmptyList.Cons(Int(7)).Cons(Int(42))

	if err := ReadArgs(args, &a, &b); err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if a != Int(42) {
		t.Fatalf(`expected a == Int(42), got %s`, a)
	}

	if b != Int(7) {
		t.Fatalf(`expected b == Int(7), got %s`, b)
	}
}

func TestReadArgsWithBadArity(t *testing.T) {
	var a, b Value

	args1 := EmptyList.Cons(Int(7)).Cons(Int(42)).Cons(Int(13))
	args2 := EmptyList.Cons(Int(7))

	if err := ReadArgs(args1, &a, &b); err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	if err := ReadArgs(args2, &a, &b); err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
}
