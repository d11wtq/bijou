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

func TestIntAddWithInt(t *testing.T) {
	a := Int(42)
	b := Int(7)

	c, err := a.Add(b)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if c != Int(49) {
		t.Fatalf(`expected c == Int(49), got %s`, c)
	}
}

func TestIntAddWithAString(t *testing.T) {
	a := Int(42)
	b := String("foo")

	c, err := a.Add(b)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	if c != nil {
		t.Fatalf(`expected c == nil, got %s`, c)
	}
}

func TestIntSubWithInt(t *testing.T) {
	a := Int(42)
	b := Int(7)

	c, err := a.Sub(b)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if c != Int(35) {
		t.Fatalf(`expected c == Int(35), got %s`, c)
	}
}

func TestIntSubWithAString(t *testing.T) {
	a := Int(42)
	b := String("foo")

	c, err := a.Sub(b)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	if c != nil {
		t.Fatalf(`expected c == nil, got %s`, c)
	}
}

func TestIntDivWithInt(t *testing.T) {
	a := Int(43)
	b := Int(7)

	c, err := a.Div(b)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if c != Int(6) {
		t.Fatalf(`expected c == Int(6), got %s`, c)
	}
}

func TestIntDivWithAString(t *testing.T) {
	a := Int(43)
	b := String("foo")

	c, err := a.Div(b)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	if c != nil {
		t.Fatalf(`expected c == nil, got %s`, c)
	}
}
