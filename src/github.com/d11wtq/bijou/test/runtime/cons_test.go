package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestConsType(t *testing.T) {
	cons := Cons(Int(42), Nil)
	if cons.Type() != ConsType {
		t.Fatalf(`expected cons.Type() == ConsType, got %s`, cons.Type())
	}
}

func TestConsEqWithEmptyCons(t *testing.T) {
	if !Eq(EmptyCons, EmptyCons) {
		t.Fatalf(`expected Eq(EmptyCons, EmptyCons), got false`)
	}
}

func TestConsEqWithOneEmptyCons(t *testing.T) {
	a := EmptyCons
	b := Cons(Int(42), EmptyCons)

	if Eq(a, b) {
		t.Fatalf(`expected !Eq(a, b), got true`)
	}
	if Eq(b, a) {
		t.Fatalf(`expected !Eq(b, a), got true`)
	}
}

func TestConsEqWithEquivalentCons(t *testing.T) {
	a := Cons(Int(7), Cons(Int(42), Nil))
	b := Cons(Int(7), Cons(Int(42), Nil))

	if !Eq(a, b) {
		t.Fatalf(`expected Eq(a, b), got false`)
	}
	if !Eq(b, a) {
		t.Fatalf(`expected Eq(b, a), got false`)
	}
}

func TestConsEqWithEquivalentConsRecursive(t *testing.T) {
	a := Cons(Int(7), Cons(Cons(Int(1), Nil), Cons(Int(42), Nil)))
	b := Cons(Int(7), Cons(Cons(Int(1), Nil), Cons(Int(42), Nil)))

	if !Eq(a, b) {
		t.Fatalf(`expected Eq(a, b), got false`)
	}
	if !Eq(b, a) {
		t.Fatalf(`expected Eq(b, a), got false`)
	}
}

func TestConsEqWithDifferentLengths(t *testing.T) {
	a := Cons(Int(42), Cons(Int(7), Nil))
	b := Cons(Int(42), Nil)

	if Eq(a, b) {
		t.Fatalf(`expected !Eq(a, b), got true`)
	}
	if Eq(b, a) {
		t.Fatalf(`expected !Eq(b, a), got true`)
	}
}

func TestConsGtWithEmptyCons(t *testing.T) {
	if Gt(EmptyCons, EmptyCons) {
		t.Fatalf(`expected !Gt(EmptyCons, EmptyCons), got false`)
	}
}

func TestConsGtWithOneEmptyCons(t *testing.T) {
	a := EmptyCons
	b := Cons(Int(42), EmptyCons)

	if Gt(a, b) {
		t.Fatalf(`expected !Gt(a, b), got true`)
	}
	if !Gt(b, a) {
		t.Fatalf(`expected Gt(b, a), got false`)
	}
}

func TestConsGtWithEquivalentCons(t *testing.T) {
	a := Cons(Int(7), Cons(Int(42), Nil))
	b := Cons(Int(7), Cons(Int(42), Nil))

	if Gt(a, b) {
		t.Fatalf(`expected !Gt(a, b), got true`)
	}
	if Gt(b, a) {
		t.Fatalf(`expected !Gt(b, a), got true`)
	}
}

func TestConsGtWithDifferentConsRecursive(t *testing.T) {
	a := Cons(Int(7), Cons(Cons(Int(2), Nil), Cons(Int(42), Nil)))
	b := Cons(Int(7), Cons(Cons(Int(1), Nil), Cons(Int(42), Nil)))

	if !Gt(a, b) {
		t.Fatalf(`expected Gt(a, b), got false`)
	}
	if Gt(b, a) {
		t.Fatalf(`expected !Gt(b, a), got true`)
	}
}

func TestConsGtWithDifferentLengths(t *testing.T) {
	a := Cons(Int(42), Cons(Int(7), Nil))
	b := Cons(Int(42), Nil)

	if !Gt(a, b) {
		t.Fatalf(`expected Gt(a, b), got false`)
	}
	if Gt(b, a) {
		t.Fatalf(`expected !Gt(b, a), got true`)
	}
}

func TestConsLtWithEmptyCons(t *testing.T) {
	if Lt(EmptyCons, EmptyCons) {
		t.Fatalf(`expected !Lt(EmptyCons, EmptyCons), got false`)
	}
}

func TestConsLtWithOneEmptyCons(t *testing.T) {
	a := EmptyCons
	b := Cons(Int(42), EmptyCons)

	if !Lt(a, b) {
		t.Fatalf(`expected Lt(a, b), got false`)
	}
	if Lt(b, a) {
		t.Fatalf(`expected !Lt(b, a), got true`)
	}
}

func TestConsLtWithEquivalentCons(t *testing.T) {
	a := Cons(Int(7), Cons(Int(42), Nil))
	b := Cons(Int(7), Cons(Int(42), Nil))

	if Lt(a, b) {
		t.Fatalf(`expected !Lt(a, b), got true`)
	}
	if Lt(b, a) {
		t.Fatalf(`expected !Lt(b, a), got true`)
	}
}

func TestConsLtWithDifferentConsRecursive(t *testing.T) {
	a := Cons(Int(7), Cons(Cons(Int(2), Nil), Cons(Int(42), Nil)))
	b := Cons(Int(7), Cons(Cons(Int(1), Nil), Cons(Int(42), Nil)))

	if Lt(a, b) {
		t.Fatalf(`expected !Lt(a, b), got true`)
	}
	if !Lt(b, a) {
		t.Fatalf(`expected Lt(b, a), got false`)
	}
}

func TestConsLtWithDifferentLengths(t *testing.T) {
	a := Cons(Int(42), Cons(Int(7), Nil))
	b := Cons(Int(42), Nil)

	if Lt(a, b) {
		t.Fatalf(`expected !Lt(a, b), got true`)
	}
	if !Lt(b, a) {
		t.Fatalf(`expected Lt(b, a), got false`)
	}
}

func TestEmptyConsEvalItself(t *testing.T) {
	env := FakeEnv()

	v, err := EmptyCons.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != EmptyCons {
		t.Fatalf(`expected v == EmptyCons, got %s`, v)
	}
}

func TestConsBindWithNonList(t *testing.T) {
	env := NewScope(nil)
	cons := Cons(Int(42), EmptyCons)
	value := Int(42)
	err := cons.(*ConsCell).Bind(env, value)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
}

func TestConsBindWithListOfInts(t *testing.T) {
	env := NewScope(nil)
	cons := Cons(Int(7), Cons(Int(42), EmptyCons))
	value := Cons(Int(7), Cons(Int(42), EmptyCons))
	err := cons.(*ConsCell).Bind(env, value)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
}

func TestConsBindWithMismatchedListOfInts(t *testing.T) {
	env := NewScope(nil)
	cons := Cons(Int(7), EmptyCons)
	value := Cons(Int(7), Cons(Int(42), EmptyCons))
	err := cons.(*ConsCell).Bind(env, value)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}

	cons = Cons(Int(7), Cons(Int(42), Cons(Int(3), EmptyCons)))
	err = cons.(*ConsCell).Bind(env, value)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
}

func TestConsBindWithListOfUnboundSymbols(t *testing.T) {
	env := NewScope(nil)
	cons := Cons(Symbol("x"), Cons(Symbol("y"), EmptyCons))
	value := Cons(Int(7), Cons(Int(42), EmptyCons))
	err := cons.(*ConsCell).Bind(env, value)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	x, ok := env.Get("x")
	if ok == false {
		t.Fatalf(`expected env.Get("x"), but not bound`)
	}
	if x != Int(7) {
		t.Fatalf(`expected x == Int(7), but got %s`, x)
	}
	y, ok := env.Get("y")
	if ok == false {
		t.Fatalf(`expected env.Get("y"), but not bound`)
	}
	if y != Int(42) {
		t.Fatalf(`expected y == Int(42), but got %s`, y)
	}
}

func TestConsBindWithVariadicSymbol(t *testing.T) {
	env := NewScope(nil)
	cons := Cons(Symbol("x"), Cons(Symbol("&"), Cons(Symbol("y"), EmptyCons)))
	value := Cons(Int(7), Cons(Int(33), Cons(Int(42), EmptyCons)))
	err := cons.(*ConsCell).Bind(env, value)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	x, ok := env.Get("x")
	if ok == false {
		t.Fatalf(`expected env.Get("x"), but not bound`)
	}
	if x != Int(7) {
		t.Fatalf(`expected x == Int(7), but got %s`, x)
	}
	y, ok := env.Get("y")
	if ok == false {
		t.Fatalf(`expected env.Get("y"), but not bound`)
	}
	if !Eq(y, Cons(Int(33), Cons(Int(42), EmptyCons))) {
		t.Fatalf(`expected y == (33 42), but got %s`, y)
	}
}

func TestConsBindWithIgnoredVariadicSymbol(t *testing.T) {
	env := NewScope(nil)
	cons := Cons(Symbol("x"), Cons(Symbol("&"), EmptyCons))
	value := Cons(Int(7), Cons(Int(33), Cons(Int(42), EmptyCons)))
	err := cons.(*ConsCell).Bind(env, value)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	x, ok := env.Get("x")
	if ok == false {
		t.Fatalf(`expected env.Get("x"), but not bound`)
	}
	if x != Int(7) {
		t.Fatalf(`expected x == Int(7), but got %s`, x)
	}
}

func TestConsBindWithBadVariadicSymbols(t *testing.T) {
	env := NewScope(nil)
	cons := Cons(
		Symbol("a"),
		Cons(
			Symbol("&"),
			Cons(Symbol("x"), Cons(Symbol("y"), EmptyCons)),
		),
	)
	value := Cons(Int(7), Cons(Int(33), Cons(Int(42), EmptyCons)))
	err := cons.(*ConsCell).Bind(env, value)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
}

func TestEmptyConsHasNoTailOrHead(t *testing.T) {
	if EmptyCons.Head() != Nil {
		t.Fatalf(`expected EmptyCons.Head() == Nil, got %s`, EmptyCons.Head())
	}

	if EmptyCons.Tail() != EmptyCons {
		t.Fatalf(
			`expected EmptyCons.Tail() == EmptyCons, got %s`,
			EmptyCons.Tail(),
		)
	}
}

func TestConsPut(t *testing.T) {
	cons, err := Cons(Int(42), Nil).Put(Int(7))

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if cons.Head() != Int(7) {
		t.Fatalf(`expected cons.Head() == Int(7), got %s`, cons.Head())
	}
	if cons.Tail().Head() != Int(42) {
		t.Fatalf(
			`expected cons.Tail().Head() == Int(42), got %s`,
			cons.Tail().Head(),
		)
	}
}

func TestConsEmpty(t *testing.T) {
	v := EmptyCons.Empty()
	if v == false {
		t.Fatalf(`expected EmptyCons.Empty(), got false`)
	}

	v = Cons(Int(42), Nil).Empty()
	if v == true {
		t.Fatalf(`expected !EmptyList.Cons(Int(42)).Empty(), got true`)
	}
}

func TestConsString(t *testing.T) {
	v := Cons(Int(42), Cons(Int(7), EmptyCons))
	s := v.String()

	if s != "(42 7)" {
		t.Fatalf(`expected s == "(42 7)", got %s`, s)
	}
}
