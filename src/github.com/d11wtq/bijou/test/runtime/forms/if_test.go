package forms_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestIfWithPassValue(t *testing.T) {
	pass := test.NewFakeValue(Int(1))
	fail := test.NewFakeValue(Int(2))
	form := EmptyList.Cons(fail).Cons(pass).Cons(Int(0)).Cons(Symbol("if"))

	v, err := form.Eval(test.FakeEnv())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(1) {
		t.Fatalf(`expected v == Int(1), got %s`, v)
	}

	if fail.Evaluated {
		t.Fatalf(`expected fail.Evaluated == false, got true`)
	}
}

func TestIfWithFailValue(t *testing.T) {
	pass := test.NewFakeValue(Int(1))
	fail := test.NewFakeValue(Int(2))
	form := EmptyList.Cons(fail).Cons(pass).Cons(Nil).Cons(Symbol("if"))

	v, err := form.Eval(test.FakeEnv())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(2) {
		t.Fatalf(`expected v == Int(2), got %s`, v)
	}

	if pass.Evaluated {
		t.Fatalf(`expected pass.Evaluated == false, got true`)
	}
}

func TestIfWithAbsentFailValue(t *testing.T) {
	pass := test.NewFakeValue(Int(1))
	form := EmptyList.Cons(pass).Cons(Nil).Cons(Symbol("if"))

	v, err := form.Eval(test.FakeEnv())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Nil {
		t.Fatalf(`expected v == Nil, got %s`, v)
	}

	if pass.Evaluated {
		t.Fatalf(`expected pass.Evaluated == false, got true`)
	}
}

func TestIfWithErrorInCondition(t *testing.T) {
	cond := test.NewFakeValue(Symbol("bad"))
	pass := test.NewFakeValue(Int(1))
	fail := test.NewFakeValue(Int(2))
	form := EmptyList.Cons(fail).Cons(pass).Cons(cond).Cons(Symbol("if"))

	v, err := form.Eval(test.FakeEnv())
	if err == nil {
		t.Fatalf(`expected err =! nil, got nil`)
	}

	if err != cond.Error {
		t.Fatalf(`expected err == cond.Error, got %s`, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}

	if pass.Evaluated {
		t.Fatalf(`expected pass.Evaluated == false, got true`)
	}

	if fail.Evaluated {
		t.Fatalf(`expected fail.Evaluated == false, got true`)
	}
}

func TestIfWithMultipleFailValues(t *testing.T) {
	pass := test.NewFakeValue(Int(1))
	f1 := test.NewFakeValue(Int(2))
	f2 := test.NewFakeValue(Int(3))
	f3 := test.NewFakeValue(Int(4))
	form := EmptyList.
		Cons(f3).
		Cons(f2).
		Cons(f1).
		Cons(pass).
		Cons(Nil).
		Cons(Symbol("if"))

	v, err := form.Eval(test.FakeEnv())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(4) {
		t.Fatalf(`expected v == Int(4), got %s`, v)
	}

	if pass.Evaluated {
		t.Fatalf(`expected pass.Evaluated == false, got true`)
	}

	if !f1.Evaluated {
		t.Fatalf(`expected f1.Evaluated == true, got false`)
	}

	if !f2.Evaluated {
		t.Fatalf(`expected f2.Evaluated == true, got false`)
	}

	if !f3.Evaluated {
		t.Fatalf(`expected f3.Evaluated == true, got false`)
	}
}
