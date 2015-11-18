package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestEvalDelegatesToValue(t *testing.T) {
	form := test.NewFakeValue(Int(42))
	v, err := Eval(form, test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}

	if form.Evaluated == false {
		t.Fatalf(`expected form.Evaluated, got false`)
	}
}

func TestEvalReturnsFromTailCalls(t *testing.T) {
	form := &Call{
		Fn:   test.FakeFn(Int(7)),
		Args: EmptyList,
		Env:  test.FakeEnv(),
	}
	v, err := Eval(form, test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(7) {
		t.Fatalf(`expected v == Int(7), got %s`, v)
	}
}

func TestEvalFormsReturnsLastResult(t *testing.T) {
	a, b := test.NewFakeValue(Int(42)), test.NewFakeValue(Int(29))
	forms := EmptyList.Append(a).Append(b)
	v, err := EvalForms(forms, test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(29) {
		t.Fatalf(`expected v == Int(29), got %s`, v)
	}

	if a.Evaluated == false {
		t.Fatalf(`expected a.Evaluated, got false`)
	}

	if b.Evaluated == false {
		t.Fatalf(`expected b.Evaluated, got false`)
	}
}

func TestEvalFormmShortCircuitsOnError(t *testing.T) {
	a, b := test.NewFakeValue(Symbol("bad")), test.NewFakeValue(Int(29))
	forms := EmptyList.Append(a).Append(b)
	v, err := EvalForms(forms, test.FakeEnv())

	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	if err != a.Error {
		t.Fatalf(`expected err == a.Error, got %s`, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}

	if a.Evaluated == false {
		t.Fatalf(`expected a.Evaluated, got false`)
	}

	if b.Evaluated == true {
		t.Fatalf(`expected !b.Evaluated, got true`)
	}
}

func TestEvalFormsReturnsFromTailCalls(t *testing.T) {
	forms := EmptyList.Append(&Call{
		Fn:   test.FakeFn(Int(7)),
		Args: EmptyList,
		Env:  test.FakeEnv(),
	})
	v, err := EvalForms(forms, test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(7) {
		t.Fatalf(`expected v == Int(7), got %s`, v)
	}
}
