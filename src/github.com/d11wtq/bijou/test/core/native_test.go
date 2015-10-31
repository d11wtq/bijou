package core_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"testing"
)

func example(args *runtime.List) (runtime.Value, error) {
	return args, nil
}

func TestNativeFunc(t *testing.T) {
	fn := core.NativeFunc(example)

	if fn.Type() != runtime.FuncType {
		t.Fatalf(`expected fn.Type() == FuncType, got %s`, fn.Type())
	}

	args := runtime.EmptyList.Cons(runtime.Int(42))
	v2, err2 := fn.Call(args)
	if err2 != nil {
		t.Fatalf(`expected err2 == nil, got %s`, err2)
	}
	if v2 != args {
		t.Fatalf(`expected v2 == args, got %s`, v2)
	}
}
