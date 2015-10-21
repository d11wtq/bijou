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
	if v, err := Nil.Eval(env); err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	} else if v != Nil {
		t.Fatalf(`expected v == Nil, got %s`, v)
	}
}
