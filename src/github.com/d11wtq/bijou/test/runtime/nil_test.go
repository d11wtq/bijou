package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestNilType(t *testing.T) {
	if Nil.Type() != NilType {
		t.Fatalf(`expected Nil.Type() == NilType, got %s`, Nil.Type())
	}
}

func TestNilEvalToSelf(t *testing.T) {
	env := FakeEnv()
	v, err := Nil.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != Nil {
		t.Fatalf(`expected v == Nil, got %s`, v)
	}
}

func TestNilEq(t *testing.T) {
	if !Nil.Eq(Nil) {
		t.Fatalf(`expected Nil.Eq(Nil), got false`)
	}
	if Nil.Eq(Int(42)) {
		t.Fatalf(`expected !Nil.Eq(Int(42)), got true`)
	}
}
