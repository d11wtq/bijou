package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestReadArgsWithCorrectArity(t *testing.T) {
	var a, b Value

	args := test.NewList(Int(42), Int(7))

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

	args1 := test.NewList(Int(13), Int(42), Int(7))
	args2 := test.NewList(Int(7))

	if err := ReadArgs(args1, &a, &b); err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	if err := ReadArgs(args2, &a, &b); err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
}
