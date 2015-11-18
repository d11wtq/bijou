package forms_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestDoReturnsLastEvaluatedExpression(t *testing.T) {
	form := test.NewList(Symbol("do"), Int(42), Int(7))
	v, err := form.Eval(test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(7) {
		t.Fatalf(`expected v == Int(7), got %s`, v)
	}
}

func TestDoPropagatesErrors(t *testing.T) {
	env := test.FakeEnv()
	foo := Symbol("foo")
	_, expected := foo.Eval(env)

	form := test.NewList(Symbol("do"), foo, Int(7))

	v, err := form.Eval(env)

	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	if err.Error() != expected.Error() {
		t.Fatalf(`expected err == expected, got %s`, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestDoWithNestedDoInTailPosition(t *testing.T) {
	form := test.NewList(
		Symbol("do"),
		Int(7),
		test.NewList(Symbol("do"), Int(42)),
	)

	v, err := form.Eval(test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestDoWithEmptyBody(t *testing.T) {
	form := test.NewList(Symbol("do"))

	v, err := form.Eval(test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestDoWithNilInTheBody(t *testing.T) {
	form := test.NewList(Symbol("do"), Nil)

	v, err := form.Eval(test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Nil {
		t.Fatalf(`expected v == Nil, got %s`, v)
	}
}

func TestDoEmitsTailCalls(t *testing.T) {
	call := &Call{
		Fn:   test.FakeFn(Nil),
		Args: EmptyList,
		Env:  test.FakeEnv(),
	}
	form := test.NewList(Symbol("do"), Int(42), call)

	v, err := form.Eval(test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != call {
		t.Fatalf(`expected v == call, got %s`, v)
	}
}

func TestDoReturnsFromNonTailCalls(t *testing.T) {
	call := &Call{
		Fn:   test.FakeFn(Int(99)),
		Args: EmptyList,
		Env:  test.FakeEnv(),
	}

	form := test.NewList(
		Symbol("do"),
		test.NewList(Symbol("def"), Symbol("x"), call),
		Symbol("x"),
	)

	v, err := form.Eval(test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(99) {
		t.Fatalf(`expected v == Int(99), got %s`, v)
	}
}
