package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"strings"
	"testing"
)

func TestFuncType(t *testing.T) {
	fn := &Func{
		Params: test.NewList(),
		Body:   test.NewList(),
		Env:    test.FakeEnv(),
	}

	if fn.Type() != FuncType {
		t.Fatalf(`expected fn.Type() == FuncType, got %s`, fn.Type())
	}
}

func TestFuncIsNotMacro(t *testing.T) {
	fn := &Func{
		Params: test.NewList(),
		Body:   test.NewList(),
		Env:    test.FakeEnv(),
	}

	if fn.IsMacro() {
		t.Fatalf(`expected fn.IsMacro() == false, got true`)
	}
}

func TestFuncEq(t *testing.T) {
	a := &Func{
		Params: test.NewList(),
		Body:   test.NewList(),
		Env:    test.FakeEnv(),
	}
	b := &Func{
		Params: test.NewList(),
		Body:   test.NewList(),
		Env:    test.FakeEnv(),
	}

	if a.Eq(b) { // operationally equivalent functions are not the same
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

func TestFuncEvalToSelf(t *testing.T) {
	fn := &Func{
		Params: test.NewList(),
		Body:   test.NewList(),
		Env:    test.FakeEnv(),
	}

	v, err := fn.Eval(test.FakeEnv())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != fn {
		t.Fatalf(`expected v == fn, got %s`, v)
	}
}

func TestFuncCallReturnsLastEvaluatedExpression(t *testing.T) {
	params := test.NewList()
	body := test.NewList(Int(7), Int(42))
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    test.FakeEnv(),
	}

	v, err := fn.Call(test.FakeEnv(), test.NewList())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestFuncCallUsesClosedEnvironment(t *testing.T) {
	params := test.NewList()
	body := test.NewList(Int(7), Symbol("foo"))
	env := test.FakeEnv()
	env.Def("foo", Int(99))
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    env,
	}

	v, err := fn.Call(test.FakeEnv(), test.NewList())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(99) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestFuncCallExtendsEnvironmentWithArgs(t *testing.T) {
	params := test.NewList(Symbol("x"))
	body := test.NewList(Int(7), Symbol("x"))
	env := test.FakeEnv()
	env.Def("foo", Int(99))
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := test.NewList(Int(21))

	v, err := fn.Call(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(21) {
		t.Fatalf(`expected v == Int(21), got %s`, v)
	}
}

func TestFuncCallValidatesTooFewArgs(t *testing.T) {
	params := test.NewList(Symbol("x"), Symbol("y"))
	body := test.NewList()
	env := test.FakeEnv()
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := test.NewList(Int(21))

	v, err := fn.Call(test.FakeEnv(), args)
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

func TestFuncCallValidatesTooManyArgs(t *testing.T) {
	params := test.NewList(Symbol("x"), Symbol("y"))
	body := test.NewList()
	env := test.FakeEnv()
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := test.NewList(Int(2), Int(9), Int(21))

	v, err := fn.Call(test.FakeEnv(), args)
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

func TestFuncCallWithVariadicArgs(t *testing.T) {
	params := test.NewList(Symbol("x"), Symbol("&"), Symbol("y"))
	body := test.NewList(Symbol("y"))
	env := test.FakeEnv()
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := test.NewList(Int(1), Int(2), Int(3), Int(4))

	v, err := fn.Call(test.FakeEnv(), args)
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

func TestFuncCallWithEmptyVariadicArgs(t *testing.T) {
	params := test.NewList(Symbol("x"), Symbol("&"), Symbol("y"))
	body := test.NewList(Symbol("y"))
	env := test.FakeEnv()
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := test.NewList(Int(1))

	v, err := fn.Call(test.FakeEnv(), args)
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

func TestFuncCallWithIgnoredVariadicArgs(t *testing.T) {
	params := test.NewList(Symbol("x"), Symbol("&"))
	body := test.NewList(Symbol("x"))
	env := test.FakeEnv()
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    env,
	}
	args := test.NewList(Int(1), Int(2), Int(3), Int(4))

	v, err := fn.Call(test.FakeEnv(), args)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(1) {
		t.Fatalf(`expected v == Int(1), got %s`, v)
	}
}

func TestFuncCallShortCirtcuitsOnError(t *testing.T) {
	v1 := test.NewFakeValue(Symbol("xx"))
	v2 := test.NewFakeValue(Symbol("yy"))

	params := test.NewList(Symbol("x"), Symbol("y"))
	body := test.NewList(v1, v2)
	fn := &Func{
		Params: params,
		Body:   body,
		Env:    test.FakeEnv(),
	}
	args := test.NewList(Int(9), Int(21))

	v, err := fn.Call(test.FakeEnv(), args)
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
