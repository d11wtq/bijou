package core_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestEqWithoutArgs(t *testing.T) {
	args := runtime.EmptyList
	v, err := core.Eq(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}

func TestEqWithOneArg(t *testing.T) {
	args := runtime.EmptyList.Append(runtime.Int(42))
	v, err := core.Eq(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}

func TestEqWithTwoEqualArgs(t *testing.T) {
	args := runtime.EmptyList.
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

func TestEqWithTwoUnequalArgs(t *testing.T) {
	args := runtime.EmptyList.
		Append(runtime.Int(42)).
		Append(runtime.Int(7))

	v, err := core.Eq(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.False {
		t.Fatalf(`expected v == False, got %s`, v)
	}
}

func TestEqWithManyEqualArgs(t *testing.T) {
	args := runtime.EmptyList.
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
	args := runtime.EmptyList.
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

func TestGtWithoutArgs(t *testing.T) {
	args := runtime.EmptyList
	v, err := core.Gt(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}

func TestGtWithOneArg(t *testing.T) {
	args := runtime.EmptyList.Append(runtime.Int(42))
	v, err := core.Gt(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}

func TestGtWithTwoGreaterArgs(t *testing.T) {
	args := runtime.EmptyList.
		Append(runtime.Int(43)).
		Append(runtime.Int(42))
	v, err := core.Gt(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}

func TestGtWithTwoEqualArgs(t *testing.T) {
	args := runtime.EmptyList.
		Append(runtime.Int(42)).
		Append(runtime.Int(42))

	v, err := core.Gt(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.False {
		t.Fatalf(`expected v == False, got %s`, v)
	}
}

func TestGtWithManyGreaterArgs(t *testing.T) {
	args := runtime.EmptyList.
		Append(runtime.Int(44)).
		Append(runtime.Int(43)).
		Append(runtime.Int(42))

	v, err := core.Gt(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}

func TestGtWithSomeLowerArgs(t *testing.T) {
	args := runtime.EmptyList.
		Append(runtime.Int(44)).
		Append(runtime.Int(43)).
		Append(runtime.Int(42)).
		Append(runtime.Int(43)).
		Append(runtime.Int(41))

	v, err := core.Gt(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.False {
		t.Fatalf(`expected v == False, got %s`, v)
	}
}

func TestLtWithoutArgs(t *testing.T) {
	args := runtime.EmptyList
	v, err := core.Lt(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}

func TestLtWithOneArg(t *testing.T) {
	args := runtime.EmptyList.Append(runtime.Int(42))
	v, err := core.Lt(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}

func TestLtWithTwoLowerArgs(t *testing.T) {
	args := runtime.EmptyList.
		Append(runtime.Int(42)).
		Append(runtime.Int(43))
	v, err := core.Lt(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}

func TestLtWithTwoEqualArgs(t *testing.T) {
	args := runtime.EmptyList.
		Append(runtime.Int(42)).
		Append(runtime.Int(42))

	v, err := core.Lt(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.False {
		t.Fatalf(`expected v == False, got %s`, v)
	}
}

func TestLtWithManyLowerArgs(t *testing.T) {
	args := runtime.EmptyList.
		Append(runtime.Int(42)).
		Append(runtime.Int(43)).
		Append(runtime.Int(44))

	v, err := core.Lt(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}
}

func TestLtWithSomeGreaterArgs(t *testing.T) {
	args := runtime.EmptyList.
		Append(runtime.Int(41)).
		Append(runtime.Int(42)).
		Append(runtime.Int(43)).
		Append(runtime.Int(42)).
		Append(runtime.Int(44))

	v, err := core.Lt(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.False {
		t.Fatalf(`expected v == False, got %s`, v)
	}
}
