package forms_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"strings"
	"testing"
)

func TestFnReturnsAFunction(t *testing.T) {
	form := test.NewList(Symbol("fn"), test.NewList(Symbol("x")))
	env := test.FakeEnv()
	v, err := form.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v.Type() != FuncType {
		t.Fatalf(`expected v.Type() == FuncType, got %s`, v.Type())
	}
	fn := v.(*Func)
	if fn.Params != form.Tail().Head() {
		t.Fatalf(`expected fn.Params == form.Tail().Head(), got %s`, fn.Params)
	}
	if fn.Body != form.Tail().Tail() {
		t.Fatalf(`expected fn.Body == form.Tail().Tail(), got %s`, fn.Body)
	}
	if fn.Env != env {
		t.Fatalf(`expected fn.Env == env, got %s`, fn.Env)
	}
}

func TestFnValidatesParameterListPresence(t *testing.T) {
	form := test.NewList(Symbol("fn"))
	v, err := form.Eval(test.FakeEnv())
	errmsg := "missing param"
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if !strings.Contains(strings.ToLower(err.Error()), errmsg) {
		t.Fatalf(`expected err to match "%s", got %s`, errmsg, err)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestFnValidatesParameterListType(t *testing.T) {
	form := test.NewList(Symbol("fn"), Int(1))
	v, err := form.Eval(test.FakeEnv())
	errmsg := "invalid param"
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if !strings.Contains(strings.ToLower(err.Error()), errmsg) {
		t.Fatalf(`expected err to match "%s", got %s`, errmsg, err)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestFnValidatesParameterTypes(t *testing.T) {
	form := test.NewList(Symbol("fn"), test.NewList(Int(1)))
	v, err := form.Eval(test.FakeEnv())
	errmsg := "invalid param"
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if !strings.Contains(strings.ToLower(err.Error()), errmsg) {
		t.Fatalf(`expected err to match "%s", got %s`, errmsg, err)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestFnValidatesVariadicIsLastParameter(t *testing.T) {
	form := test.NewList(
		Symbol("fn"),
		test.NewList(Symbol("&"), Symbol("x"), Symbol("y")),
	)

	v, err := form.Eval(test.FakeEnv())
	errmsg := "variadic"
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if !strings.Contains(strings.ToLower(err.Error()), errmsg) {
		t.Fatalf(`expected err to match "%s", got %s`, errmsg, err)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestFnAllowsEmptyVariadic(t *testing.T) {
	form := test.NewList(Symbol("fn"), test.NewList(Symbol("&")))

	env := test.FakeEnv()
	v, err := form.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v.Type() != FuncType {
		t.Fatalf(`expected v.Type() == FuncType, got %s`, v.Type())
	}
	fn := v.(*Func)
	if fn.Params != form.Tail().Head() {
		t.Fatalf(`expected fn.Params == form.Tail().Head(), got %s`, fn.Params)
	}
}
