package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"strings"
	"testing"
)

func TestMacroType(t *testing.T) {
	mc := &Macro{
		Func{
			Params: EmptyList,
			Body:   EmptyList,
			Env:    test.FakeEnv(),
		},
	}

	if mc.Type() != MacroType {
		t.Fatalf(`expected mc.Type() == MacroType, got %s`, mc.Type())
	}
}

func TestMacroIsMacro(t *testing.T) {
	mc := &Macro{
		Func{
			Params: EmptyList,
			Body:   EmptyList,
			Env:    test.FakeEnv(),
		},
	}

	if !mc.IsMacro() {
		t.Fatalf(`expected mc.IsMacro() == true, got false`)
	}
}

func TestMacroEq(t *testing.T) {
	a := &Macro{
		Func{
			Params: EmptyList,
			Body:   EmptyList,
			Env:    test.FakeEnv(),
		},
	}
	b := &Macro{
		Func{
			Params: EmptyList,
			Body:   EmptyList,
			Env:    test.FakeEnv(),
		},
	}

	if a.Eq(b) { // operationally equivalent macros are not the same
		t.Fatalf(`expected !a.Eq(b), got true`)
	}
	if b.Eq(a) {
		t.Fatalf(`expected !b.Eq(a), got true`)
	}

	if !a.Eq(a) {
		t.Fatalf(`expected a.Eq(a), got false`)
	}
	if !b.Eq(b) {
		t.Fatalf(`expected b.Eq(b), got false`)
	}
}

func TestMacroEvalToSelf(t *testing.T) {
	mc := &Macro{
		Func{
			Params: EmptyList,
			Body:   EmptyList,
			Env:    test.FakeEnv(),
		},
	}

	v, err := mc.Eval(test.FakeEnv())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != mc {
		t.Fatalf(`expected v == mc, got %s`, v)
	}
}

func TestMacroCallReturnsLastEvaluatedExpression(t *testing.T) {
	params := EmptyList
	body := EmptyList.Cons(Int(42)).Cons(Int(7))
	mc := &Macro{
		Func{
			Params: params,
			Body:   body,
			Env:    test.FakeEnv(),
		},
	}

	v, err := mc.Call(test.FakeEnv(), EmptyList)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestMacroCallUsesClosedEnvironment(t *testing.T) {
	params := EmptyList
	body := EmptyList.Cons(Symbol("foo")).Cons(Int(7))
	env := test.FakeEnv()
	env.Def("foo", Int(99))
	mc := &Macro{
		Func{
			Params: params,
			Body:   body,
			Env:    env,
		},
	}

	v, err := mc.Call(test.FakeEnv(), EmptyList)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(99) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestMacroCallExtendsEnvironmentWithArgs(t *testing.T) {
	params := EmptyList.Cons(Symbol("x"))
	body := EmptyList.Cons(Symbol("x")).Cons(Int(7))
	env := test.FakeEnv()
	env.Def("foo", Int(99))
	mc := &Macro{
		Func{
			Params: params,
			Body:   body,
			Env:    env,
		},
	}
	args := EmptyList.Cons(Int(21))

	v, err := mc.Call(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(21) {
		t.Fatalf(`expected v == Int(21), got %s`, v)
	}
}

func TestMacroCallValidatesTooFewArgs(t *testing.T) {
	params := EmptyList.Cons(Symbol("y")).Cons(Symbol("x"))
	body := EmptyList
	env := test.FakeEnv()
	mc := &Macro{
		Func{
			Params: params,
			Body:   body,
			Env:    env,
		},
	}
	args := EmptyList.Cons(Int(21))

	v, err := mc.Call(test.FakeEnv(), args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	if _, ok := err.(*ArgumentError); !ok {
		t.Fatalf(`expected err.(*ArgumentError), got %s`, err)
	}

	match := "wanted 2, got 1"
	if !strings.Contains(strings.ToLower(err.Error()), match) {
		t.Fatalf(`expected err to match "%s", got %s`, match, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestMacroCallValidatesTooManyArgs(t *testing.T) {
	params := EmptyList.Cons(Symbol("y")).Cons(Symbol("x"))
	body := EmptyList
	env := test.FakeEnv()
	mc := &Macro{
		Func{
			Params: params,
			Body:   body,
			Env:    env,
		},
	}
	args := EmptyList.Cons(Int(21)).Cons(Int(9)).Cons(Int(2))

	v, err := mc.Call(test.FakeEnv(), args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	if _, ok := err.(*ArgumentError); !ok {
		t.Fatalf(`expected err.(*ArgumentError), got %s`, err)
	}

	match := "wanted 2, got 3"
	if !strings.Contains(strings.ToLower(err.Error()), match) {
		t.Fatalf(`expected err to match "%s", got %s`, match, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestMacroCallShortCirtcuitsOnError(t *testing.T) {
	v1 := test.NewFakeValue(Symbol("xx"))
	v2 := test.NewFakeValue(Symbol("yy"))

	params := EmptyList.Cons(Symbol("y")).Cons(Symbol("x"))
	body := EmptyList.Cons(v2).Cons(v1)
	mc := &Macro{
		Func{
			Params: params,
			Body:   body,
			Env:    test.FakeEnv(),
		},
	}
	args := EmptyList.Cons(Int(21)).Cons(Int(9))

	v, err := mc.Call(test.FakeEnv(), args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if err != v1.Error {
		t.Fatalf(`expected err == v1.Error, got %s`, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}

	if v2.Evaluated {
		t.Fatalf(`expected v2.Evaluated == false, got true`)
	}
}
