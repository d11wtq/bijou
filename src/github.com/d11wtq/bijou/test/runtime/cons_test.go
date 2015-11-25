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
