package core_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestNot(t *testing.T) {
	args := runtime.EmptyList
	v, err := core.Not(test.FakeEnv(), args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}

	args = runtime.EmptyList.Cons(runtime.Int(0))
	v, err = core.Not(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.False {
		t.Fatalf(`expected v == False, got %s`, v)
	}

	args = runtime.EmptyList.Cons(runtime.Int(42))
	v, err = core.Not(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.False {
		t.Fatalf(`expected v == False, got %s`, v)
	}

	args = runtime.EmptyList.Cons(runtime.True)
	v, err = core.Not(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.False {
		t.Fatalf(`expected v == False, got %s`, v)
	}

	args = runtime.EmptyList.Cons(runtime.False)
	v, err = core.Not(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}

	args = runtime.EmptyList.Cons(runtime.Nil)
	v, err = core.Not(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.True {
		t.Fatalf(`expected v == True, got %s`, v)
	}

	args = runtime.EmptyList.Cons(runtime.EmptyList)
	v, err = core.Not(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != runtime.False {
		t.Fatalf(`expected v == False, got %s`, v)
	}

	args = runtime.EmptyList.Cons(runtime.Int(0)).Cons(runtime.Int(1))
	v, err = core.Not(test.FakeEnv(), args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}
