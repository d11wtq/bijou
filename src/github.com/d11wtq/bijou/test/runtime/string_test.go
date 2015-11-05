package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestStringType(t *testing.T) {
	s := String("example")
	if s.Type() != StringType {
		t.Fatalf(`expected s.Type() == StringType, got %s`, s.Type())
	}
}

func TestStringEvalToSelf(t *testing.T) {
	s := String("example")
	env := FakeEnv()
	if v, err := s.Eval(env); err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	} else if v != String("example") {
		t.Fatalf(`expected v == String("example"), got %s`, v)
	}
}

func TestStringEq(t *testing.T) {
	if !String("foo").Eq(String("foo")) {
		t.Fatalf(`expected String("foo").Eq(String("foo")), got false`)
	}
	if String("foo").Eq(String("bar")) {
		t.Fatalf(`expected !String("foo").Eq(String("bar")), got true`)
	}
}
