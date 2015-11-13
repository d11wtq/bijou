package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestFunctionCallWithSymbol(t *testing.T) {
	fn := &Func{
		Params: test.NewList(),
		Body:   test.NewList(Int(42)),
		Env:    test.FakeEnv(),
	}
	env := test.FakeEnv()
	env.Def("example", fn)
	form := test.NewList(Symbol("example"))
	v, err := Eval(form, env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestFunctionCallWithLambda(t *testing.T) {
	fn := &Func{
		Params: test.NewList(),
		Body:   test.NewList(Int(42)),
		Env:    test.FakeEnv(),
	}
	form := test.NewList(fn)
	v, err := Eval(form, test.FakeEnv())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestFunctionCallWithArguments(t *testing.T) {
	fn := &Func{
		Params: test.NewList(Symbol("a")),
		Body:   test.NewList(Symbol("a")),
		Env:    test.FakeEnv(),
	}
	env := test.FakeEnv()
	env.Def("example", fn)
	env.Def("x", Int(42))
	form := test.NewList(Symbol("example"), Symbol("x"))
	v, err := Eval(form, env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}
