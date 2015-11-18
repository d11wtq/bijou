package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestCallReturnWithASimpleValue(t *testing.T) {
	fn := &Func{
		Env:    test.FakeEnv(),
		Params: EmptyList.Append(Symbol("x")),
		Body:   EmptyList.Append(Symbol("x")),
	}

	call := &Call{
		Fn:   fn,
		Args: EmptyList.Append(Int(42)),
		Env:  test.FakeEnv(),
	}

	v, err := call.Return()
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestCallReturnWithATailCall(t *testing.T) {
	fn := test.FakeFn(
		&Call{
			Fn:   test.FakeFn(Int(7)),
			Args: EmptyList,
			Env:  test.FakeEnv(),
		},
	)

	call := &Call{
		Fn:   fn,
		Args: EmptyList,
		Env:  test.FakeEnv(),
	}

	v, err := call.Return()
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != Int(7) {
		t.Fatalf(`expected v == Int(7), got %s`, v)
	}
}
