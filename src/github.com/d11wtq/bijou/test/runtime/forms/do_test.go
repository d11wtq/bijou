package forms_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestDoReturnsLastEvaluatedExpression(t *testing.T) {
	form := EmptyList.Cons(Int(7)).Cons(Int(42)).Cons(Symbol("do"))
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

	form := EmptyList.Cons(Int(7)).Cons(foo).Cons(Symbol("do"))

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
	form := EmptyList.
		Cons(EmptyList.Cons(Int(42)).Cons(Symbol("do"))).
		Cons(Int(7)).
		Cons(Symbol("do"))

	v, err := form.Eval(test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestDoWithEmptyBody(t *testing.T) {
	form := EmptyList.Cons(Symbol("do"))

	v, err := form.Eval(test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestDoWithNilInTheBody(t *testing.T) {
	form := EmptyList.Cons(Nil).Cons(Symbol("do"))

	v, err := form.Eval(test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Nil {
		t.Fatalf(`expected v == Nil, got %s`, v)
	}
}
