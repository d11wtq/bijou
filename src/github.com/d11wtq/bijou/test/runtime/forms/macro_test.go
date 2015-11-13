package forms_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"strings"
	"testing"
)

func TestMacroReturnsAMacro(t *testing.T) {
	form := test.NewList(Symbol("macro"), test.NewList(Symbol("x")))
	env := test.FakeEnv()
	v, err := form.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v.Type() != MacroType {
		t.Fatalf(`expected v.Type() == MacroType, got %s`, v.Type())
	}
	mc := v.(*Macro)
	if mc.Params != form.Tail().Head() {
		t.Fatalf(`expected macro.Params == form.Tail().Head(), got %s`, mc.Params)
	}
	if mc.Body != form.Tail().Tail() {
		t.Fatalf(`expected macro.Body == form.Tail().Tail(), got %s`, mc.Body)
	}
	if mc.Env != env {
		t.Fatalf(`expected macro.Env == env, got %s`, mc.Env)
	}
}

func TestMacroValidatesParameterListPresence(t *testing.T) {
	form := test.NewList(Symbol("macro"))
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

func TestMacroValidatesParameterListType(t *testing.T) {
	form := test.NewList(Symbol("macro"), Int(1))
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

func TestMacroValidatesParameterTypes(t *testing.T) {
	form := test.NewList(Symbol("macro"), test.NewList(Int(1)))
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

func TestMacroValidatesVariadicIsLastParameter(t *testing.T) {
	form := test.NewList(
		Symbol("macro"),
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

func TestMacroAllowsEmptyVariadic(t *testing.T) {
	form := test.NewList(Symbol("macro"), test.NewList(Symbol("&")))

	env := test.FakeEnv()
	v, err := form.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v.Type() != MacroType {
		t.Fatalf(`expected v.Type() == MacroType, got %s`, v.Type())
	}
	mc := v.(*Macro)
	if mc.Params != form.Tail().Head() {
		t.Fatalf(`expected macro.Params == form.Tail().Head(), got %s`, mc.Params)
	}
}
