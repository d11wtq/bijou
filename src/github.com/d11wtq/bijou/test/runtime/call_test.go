package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestFunctionCallWithSymbol(t *testing.T) {
	fn := &Func{
		Params: EmptyList,
		Body:   EmptyList.Cons(Int(42)),
		Env:    test.FakeEnv(),
	}
	env := test.FakeEnv()
	env.Def("example", fn)
	form := EmptyList.Cons(Symbol("example"))
	v, err := form.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestFunctionCallWithLambda(t *testing.T) {
	fn := &Func{
		Params: EmptyList,
		Body:   EmptyList.Cons(Int(42)),
		Env:    test.FakeEnv(),
	}
	form := EmptyList.Cons(fn)
	v, err := form.Eval(test.FakeEnv())
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestFunctionCallWithArguments(t *testing.T) {
	fn := &Func{
		Params: EmptyList.Cons(Symbol("a")),
		Body:   EmptyList.Cons(Symbol("a")),
		Env:    test.FakeEnv(),
	}
	env := test.FakeEnv()
	env.Def("example", fn)
	form := EmptyList.Cons(Int(1)).Cons(Symbol("example"))
	v, err := form.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Int(1) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}
