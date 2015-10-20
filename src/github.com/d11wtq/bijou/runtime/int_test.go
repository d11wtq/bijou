package runtime

import (
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

func TestIntCastToNative(t *testing.T) {
	i := Int(42)
	if int(i) != 42 {
		t.Fatalf(`expected int(i) == 42, got %d`, int(i))
	}
}
