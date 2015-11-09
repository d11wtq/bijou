package core_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestEqWithoutArgs(t *testing.T) {
	args := &runtime.List{}
	v, err := core.Eq(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}

func TestEqWithOneArg(t *testing.T) {
	args := (&runtime.List{}).Append(runtime.Int(42))
	v, err := core.Eq(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}

func TestEqWithTwoEqualArgs(t *testing.T) {
	args := (&runtime.List{}).Append(runtime.Int(42)).Append(runtime.Int(42))
	v, err := core.Eq(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}

func TestEqWithTwoUnequalArgs(t *testing.T) {
	args := (&runtime.List{}).Append(runtime.Int(42)).Append(runtime.Int(7))

	v, err := core.Eq(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.False {
		t.Fatalf(`expected v == False, got %s`, v)
	}
}

func TestEqWithManyEqualArgs(t *testing.T) {
	args := (&runtime.List{}).
		Append(runtime.Int(42)).
		Append(runtime.Int(42)).
		Append(runtime.Int(42))

	v, err := core.Eq(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}

func TestEqWithManyUnequalArgs(t *testing.T) {
	args := (&runtime.List{}).
		Append(runtime.Int(7)).
		Append(runtime.Int(42)).
		Append(runtime.Int(42))

	v, err := core.Eq(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.False {
		t.Fatalf(`expected v == False, got %s`, v)
	}
}
