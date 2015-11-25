package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestApplyInvokesFnWithArgs(t *testing.T) {
	fn := test.FakeFn(Symbol("args"))
	res, err := Apply(fn, test.FakeEnv(), EmptyList.Append(Int(42)))
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if !Eq(res, EmptyList.Append(Int(42))) {
		t.Fatalf(`expected res == (42), got %s`, res)
	}
}

func TestApplyReturnsFromTailCalls(t *testing.T) {
	fn := test.FakeFn(
		&Call{
			Fn:   test.FakeFn(Int(7)),
			Args: EmptyList,
			Env:  test.FakeEnv(),
		},
	)
	res, err := Apply(fn, test.FakeEnv(), EmptyList)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if !Eq(res, Int(7)) {
		t.Fatalf(`expected res == 7, got %s`, res)
	}
}
