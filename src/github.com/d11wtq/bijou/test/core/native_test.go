package core_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func example(env runtime.Env, args *runtime.List) (runtime.Value, error) {
	return args, nil
}

func TestGoFunc(t *testing.T) {
	fn := core.GoFunc(example)

	if fn.Type() != runtime.FuncType {
		t.Fatalf(`expected fn.Type() == FuncType, got %s`, fn.Type())
	}
	if _, ok := fn.(runtime.Callable); ok == false {
		t.Fatalf(`expected fn.(Callable), got false`)
	}

	args := runtime.EmptyList.Cons(runtime.Int(42))
	v2, err2 := fn.Call(test.FakeEnv(), args)
	if err2 != nil {
		t.Fatalf(`expected err2 == nil, got %s`, err2)
	}
	if v2 != args {
		t.Fatalf(`expected v2 == args, got %s`, v2)
	}
}
