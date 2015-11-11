package core_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestAddWithoutArgs(t *testing.T) {
	args := runtime.EmptyList
	v, err := core.Add(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.Int(0) {
		t.Fatalf(`expected v == Int(0), got %s`, v)
	}
}

func TestAddWithASingleArg(t *testing.T) {
	args := runtime.EmptyList.Append(runtime.Int(42))
	v, err := core.Add(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestAddWithMultipleArgs(t *testing.T) {
	args := runtime.EmptyList.Append(runtime.Int(42)).Append(runtime.Int(5))
	v, err := core.Add(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.Int(47) {
		t.Fatalf(`expected v == Int(47), got %s`, v)
	}
}

func TestAddWithNonAdditiveData(t *testing.T) {
	args := runtime.EmptyList.
		Append(runtime.Int(42)).
		Append(runtime.Symbol("x"))
	v, err := core.Add(test.FakeEnv(), args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestSubWithoutArgs(t *testing.T) {
	args := runtime.EmptyList
	v, err := core.Sub(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.Int(0) {
		t.Fatalf(`expected v == Int(0), got %s`, v)
	}
}

func TestSubWithASingleArg(t *testing.T) {
	args := runtime.EmptyList.Append(runtime.Int(42))
	v, err := core.Sub(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != -runtime.Int(42) {
		t.Fatalf(`expected v == -Int(42), got %s`, v)
	}
}

func TestSubWithMultipleArgs(t *testing.T) {
	args := runtime.EmptyList.Append(runtime.Int(42)).Append(runtime.Int(5))
	v, err := core.Sub(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.Int(37) {
		t.Fatalf(`expected v == Int(37), got %s`, v)
	}
}

func TestSubWithNonSubtractiveData(t *testing.T) {
	args := runtime.EmptyList.
		Append(runtime.Int(42)).
		Append(runtime.Symbol("x"))
	v, err := core.Add(test.FakeEnv(), args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestDivWithoutArgs(t *testing.T) {
	args := runtime.EmptyList
	v, err := core.Div(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.Int(1) {
		t.Fatalf(`expected v == Int(1), got %s`, v)
	}
}

func TestDivWithASingleArg(t *testing.T) {
	args := runtime.EmptyList.Append(runtime.Int(2))
	v, err := core.Div(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.Int(0) {
		t.Fatalf(`expected v == Int(0), got %s`, v)
	}
}

func TestDivWithJustOne(t *testing.T) {
	args := runtime.EmptyList.Append(runtime.Int(1))
	v, err := core.Div(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.Int(1) {
		t.Fatalf(`expected v == Int(1), got %s`, v)
	}
}

func TestDivWithMultipleArgs(t *testing.T) {
	args := runtime.EmptyList.
		Append(runtime.Int(42)).
		Append(runtime.Int(7)).
		Append(runtime.Int(2))
	v, err := core.Div(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.Int(3) {
		t.Fatalf(`expected v == Int(3), got %s`, v)
	}
}

func TestDivWithDivideByZero(t *testing.T) {
	args := runtime.EmptyList.Append(runtime.Int(42)).Append(runtime.Int(0))
	v, err := core.Div(test.FakeEnv(), args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestDivWithNonDivisibleData(t *testing.T) {
	args := runtime.EmptyList.
		Append(runtime.Int(42)).
		Append(runtime.Symbol("x"))
	v, err := core.Div(test.FakeEnv(), args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}
