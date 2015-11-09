package forms_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"strings"
	"testing"
)

func TestIfWithoutCondition(t *testing.T) {
	form := test.NewList(Symbol("if"))
	errmsg := "missing condition"

	v, err := form.Eval(test.FakeEnv())
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if !strings.Contains(strings.ToLower(err.Error()), errmsg) {
		t.Fatalf(`expected err to match "%s", got: %s`, errmsg, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestIfWithoutPassValue(t *testing.T) {
	form := test.NewList(Symbol("if"), Int(1))
	errmsg := "missing body"

	v, err := form.Eval(test.FakeEnv())
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if !strings.Contains(strings.ToLower(err.Error()), errmsg) {
		t.Fatalf(`expected err to match "%s", got: %s`, errmsg, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestIfWithPassValue(t *testing.T) {
	pass := test.NewFakeValue(Int(1))
	fail := test.NewFakeValue(Int(2))
	form := test.NewList(Symbol("if"), Int(0), pass, fail)

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
	form := test.NewList(Symbol("if"), Nil, pass, fail)

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

func TestIfWithoutFailValue(t *testing.T) {
	pass := test.NewFakeValue(Int(1))
	form := test.NewList(Symbol("if"), Nil, pass)

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
	form := test.NewList(Symbol("if"), cond, pass, fail)

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

func TestIfWithErrorInPassValue(t *testing.T) {
	cond := test.NewFakeValue(Int(1))
	pass := test.NewFakeValue(Symbol("bad"))
	fail := test.NewFakeValue(Int(2))
	form := test.NewList(Symbol("if"), cond, pass, fail)

	v, err := form.Eval(test.FakeEnv())
	if err == nil {
		t.Fatalf(`expected err =! nil, got nil`)
	}

	if err != pass.Error {
		t.Fatalf(`expected err == pass.Error, got %s`, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestIfWithErrorInFailValue(t *testing.T) {
	cond := test.NewFakeValue(Nil)
	pass := test.NewFakeValue(Int(1))
	fail := test.NewFakeValue(Symbol("bad"))
	form := test.NewList(Symbol("if"), cond, pass, fail)

	v, err := form.Eval(test.FakeEnv())
	if err == nil {
		t.Fatalf(`expected err =! nil, got nil`)
	}

	if err != fail.Error {
		t.Fatalf(`expected err == fail.Error, got %s`, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestIfWithMultipleFailValues(t *testing.T) {
	pass := test.NewFakeValue(Int(1))
	f1 := test.NewFakeValue(Int(2))
	f2 := test.NewFakeValue(Int(3))
	f3 := test.NewFakeValue(Int(4))
	form := test.NewList(Symbol("if"), Nil, pass, f1, f2, f3)

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
