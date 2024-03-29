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

func TestStringGt(t *testing.T) {
	if String("foo").Gt(String("foo")) {
		t.Fatalf(`expected !String("foo").Gt(String("foo")), got true`)
	}
	if !String("foo").Gt(String("bar")) {
		t.Fatalf(`expected String("foo").Gt(String("bar")), got false`)
	}
	if !String("foo").Gt(String("fo")) {
		t.Fatalf(`expected String("foo").Gt(String("fo")), got false`)
	}
	if String("foo").Gt(EmptyList) {
		t.Fatalf(`expected !String("foo").Gt(EmptyList), got true`)
	}
	if !String("41").Gt(Int(42)) {
		t.Fatalf(`expected String("41").Gt(Int(42)), got false`)
	}
}

func TestStringLt(t *testing.T) {
	if String("foo").Lt(String("foo")) {
		t.Fatalf(`expected !String("foo").Lt(String("foo")), got true`)
	}
	if !String("bar").Lt(String("foo")) {
		t.Fatalf(`expected String("bar").Lt(String("foo")), got false`)
	}
	if !String("fo").Lt(String("foo")) {
		t.Fatalf(`expected String("fo").Lt(String("foo")), got false`)
	}
	if !String("foo").Lt(EmptyList) {
		t.Fatalf(`expected String("foo").Lt(EmptyList), got false`)
	}
	if String("41").Lt(Int(42)) {
		t.Fatalf(`expected !String("41").Lt(Int(42)), got true`)
	}
}

func TestStringHead(t *testing.T) {
	if v := String("foo").Head(); v != Int('f') {
		t.Fatalf(`expected String("foo").Head() == Int('f'), got`, v)
	}

	if v := String("").Head(); v != Nil {
		t.Fatalf(`expected String("").Head() == Nil, got`, v)
	}
}

func TestStringTail(t *testing.T) {
	if v := String("foo").Tail(); v != String("oo") {
		t.Fatalf(`expected String("foo").Tail() == String("oo"), got`, v)
	}

	if v := String("").Tail(); v != String("") {
		t.Fatalf(`expected String("").Tail() == String(""), got`, v)
	}
}

func TestStringPut(t *testing.T) {
	v, err := String("foo").Put(Int(100))
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != String("food") {
		t.Fatalf(`expected v == String("food"), got`, v)
	}

	v, err = String("foo").Put(String("d"))
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got`, v)
	}
}

func TestStringEmpty(t *testing.T) {
	if v := String("foo").Empty(); v == true {
		t.Fatalf(`expected !String("foo").Empty(), got true`)
	}

	if v := String("").Empty(); v == false {
		t.Fatalf(`expected String("").Empty(), got false`)
	}
}

func TestStringString(t *testing.T) {
	v := String("hello \"world\" \\ \n \r \t")
	s := v.String()
	if s != `"hello \"world\" \\ \n \r \t"` {
		t.Fatalf(`expected s == `+"`"+`"hello \"world\" \\ \n \r \t"`+"` got %s", s)
	}
}
