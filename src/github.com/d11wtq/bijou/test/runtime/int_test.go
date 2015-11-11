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

func TestIntGt(t *testing.T) {
	if !Int(42).Gt(Int(41)) {
		t.Fatalf(`expected Int(42).Gt(Int(41)), got false`)
	}
	if Int(42).Gt(Int(42)) {
		t.Fatalf(`expected !Int(42).Gt(Int(42)), got true`)
	}
	if Int(41).Gt(Int(42)) {
		t.Fatalf(`expected !Int(41).Gt(Int(42)), got true`)
	}
	if !Int(0).Gt(False) {
		t.Fatalf(`expected Int(0).Gt(False), got false`)
	}
	if !Int(0).Gt(Nil) {
		t.Fatalf(`expected Int(0).Gt(Nil), got false`)
	}
}

func TestIntLt(t *testing.T) {
	if !Int(41).Lt(Int(42)) {
		t.Fatalf(`expected Int(41).Lt(Int(42)), got false`)
	}
	if Int(42).Lt(Int(42)) {
		t.Fatalf(`expected !Int(42).Lt(Int(42)), got true`)
	}
	if Int(42).Lt(Int(42)) {
		t.Fatalf(`expected !Int(41).Lt(Int(42)), got true`)
	}
	if !Int(42).Lt(Symbol("a")) {
		t.Fatalf(`expected Int(42).Lt(Symbol("a")), got false`)
	}
	if !Int(42).Lt(String("")) {
		t.Fatalf(`expected Int(42).Lt(String("")), got false`)
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

func TestIntMulWithInt(t *testing.T) {
	a := Int(42)
	b := Int(2)

	c, err := a.Mul(b)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if c != Int(84) {
		t.Fatalf(`expected c == Int(84), got %s`, c)
	}
}

func TestIntMulWithAString(t *testing.T) {
	a := Int(42)
	b := String("foo")

	c, err := a.Mul(b)
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
