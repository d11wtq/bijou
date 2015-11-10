package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestIntType(t *testing.T) {
	i := Int(42)
	if i.Type() != IntType {
		t.Fatalf(`expected i.Type() == IntType, got %s`, i.Type())
	}
}

func TestIntEvalToSelf(t *testing.T) {
	i := Int(42)
	env := FakeEnv()
	if v, err := i.Eval(env); err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	} else if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestIntEq(t *testing.T) {
	if !Int(42).Eq(Int(42)) {
		t.Fatalf(`expected Int(42).Eq(Int(42)), got false`)
	}
	if Int(7).Eq(Int(42)) {
		t.Fatalf(`expected !Int(7).Eq(Int(42)), got true`)
	}
}

func TestIntString(t *testing.T) {
	s := Int(42).String()
	if s != "42" {
		t.Fatalf(`expected s == "42", got %s`, s)
	}
}
