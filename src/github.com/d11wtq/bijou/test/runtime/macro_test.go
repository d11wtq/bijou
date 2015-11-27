package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestMacroType(t *testing.T) {
	mc := &Macro{
		Params: test.NewList(),
		Body:   test.NewList(),
		Env:    test.FakeEnv(),
	}

	if mc.Type() != MacroType {
		t.Fatalf(`expected mc.Type() == MacroType, got %s`, mc.Type())
	}
}

func TestMacroEvalToSelf(t *testing.T) {
	mc := &Macro{
		Params: test.NewList(),
		Body:   test.NewList(),
		Env:    test.FakeEnv(),
	}

	v, err := mc.Eval(test.FakeEnv())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != mc {
		t.Fatalf(`expected v == mc, got %s`, v)
	}
}

func TestMacroExpandReturnsLastEvaluatedExpression(t *testing.T) {
	params := test.NewList()
	body := test.NewList(Int(7), Int(42))
	mc := &Macro{
		Params: params,
		Body:   body,
		Env:    test.FakeEnv(),
	}

	v, err := mc.Expand(test.FakeEnv(), test.NewList())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestMacroExpandUsesClosedEnvironment(t *testing.T) {
	params := test.NewList()
	body := test.NewList(Int(7), Symbol("foo"))
	env := test.FakeEnv()
	env.Def("foo", Int(99))
	mc := &Macro{
		Params: params,
		Body:   body,
		Env:    env,
	}

	v, err := mc.Expand(test.FakeEnv(), test.NewList())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(99) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestMacroExpandExtendsEnvironmentWithArgs(t *testing.T) {
	params := test.NewList(Symbol("x"))
	body := test.NewList(Int(7), Symbol("x"))
	env := test.FakeEnv()
	env.Def("foo", Int(99))
	mc := &Macro{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := test.NewList(Int(21))

	v, err := mc.Expand(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(21) {
		t.Fatalf(`expected v == Int(21), got %s`, v)
	}
}

func TestMacroExpandValidatesTooFewArgs(t *testing.T) {
	params := test.NewList(Symbol("x"), Symbol("y"))
	body := test.NewList()
	env := test.FakeEnv()
	mc := &Macro{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := test.NewList(Int(21))

	v, err := mc.Expand(test.FakeEnv(), args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	if _, ok := err.(*ArgumentError); !ok {
		t.Fatalf(`expected err.(*ArgumentError), got %s`, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestMacroExpandValidatesTooManyArgs(t *testing.T) {
	params := test.NewList(Symbol("x"), Symbol("y"))
	body := test.NewList()
	env := test.FakeEnv()
	mc := &Macro{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := test.NewList(Int(2), Int(9), Int(21))

	v, err := mc.Expand(test.FakeEnv(), args)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	if _, ok := err.(*ArgumentError); !ok {
		t.Fatalf(`expected err.(*ArgumentError), got %s`, err)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestMacroExpandWithVariadicArgs(t *testing.T) {
	params := test.NewList(Symbol("x"), Symbol("&"), Symbol("y"))
	body := test.NewList(Symbol("y"))
	env := test.FakeEnv()
	mc := &Macro{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := test.NewList(Int(1), Int(2), Int(3), Int(4))

	v, err := mc.Expand(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	y, ok := v.(Sequence)
	if ok == false {
		t.Fatalf(`expected v.(Sequence), but not a Sequence`)
	}

	if y.Head() != Int(2) {
		t.Fatalf(`expected y.Head() == Int(2), got %s`, y.Head())
	}
}

func TestMacroExpandWithEmptyVariadicArgs(t *testing.T) {
	params := test.NewList(Symbol("x"), Symbol("&"), Symbol("y"))
	body := test.NewList(Symbol("y"))
	env := test.FakeEnv()
	mc := &Macro{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := test.NewList(Int(1))

	v, err := mc.Expand(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	y, ok := v.(Sequence)
	if ok == false {
		t.Fatalf(`expected v.(Sequence), but not a Sequence`)
	}

	if !y.Empty() {
		t.Fatalf(`expected y.Empty(), got false`)
	}
}

func TestMacroExpandWithIgnoredVariadicArgs(t *testing.T) {
	params := test.NewList(Symbol("x"), Symbol("&"))
	body := test.NewList(Symbol("x"))
	env := test.FakeEnv()
	mc := &Macro{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := test.NewList(Int(1), Int(2), Int(3), Int(4))

	v, err := mc.Expand(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(1) {
		t.Fatalf(`expected v == Int(1), got %s`, v)
	}
}

func TestMacroExpandShortCirtcuitsOnError(t *testing.T) {
	v1 := test.NewFakeValue(Symbol("xx"))
	v2 := test.NewFakeValue(Symbol("yy"))

	params := test.NewList(Symbol("x"), Symbol("y"))
	body := test.NewList(v1, v2)
	mc := &Macro{
		Params: params,
		Body:   body,
		Env:    test.FakeEnv(),
	}
	args := test.NewList(Int(9), Int(21))

	v, err := mc.Expand(test.FakeEnv(), args)
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

func TestMacroString(t *testing.T) {
	mc := &Macro{
		Params: EmptyList,
		Body:   EmptyList,
		Env:    test.FakeEnv(),
	}

	if mc.String() != "#<macro>" {
		t.Fatalf(`expected mc.String() == "#<macro>", got %s`, mc.String())
	}
}
