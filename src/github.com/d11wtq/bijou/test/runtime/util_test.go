package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestIsList(t *testing.T) {
	var v Value
	var ok bool
	var s Sequence

	v = Int(42)
	_, ok = IsList(v)
	if ok == true {
		t.Fatalf(`expected !IsList(v), got true`)
	}

	v = String("foo")
	_, ok = IsList(v)
	if ok == true {
		t.Fatalf(`expected !IsList(v), got true`)
	}

	v = Cons(Int(42), String("foo"))
	s, ok = IsList(v)
	if ok == false {
		t.Fatalf(`expected IsList(v), got false`)
	}
	_, ok = s.(*ConsCell)
	if ok == false {
		t.Fatalf(`expected s.(*ConsCell), got false`)
	}

	v = (&List{}).Append(Int(42))
	s, ok = IsList(v)
	if ok == false {
		t.Fatalf(`expected IsList(v), got false`)
	}
	_, ok = s.(*List)
	if ok == false {
		t.Fatalf(`expected s.(*List), got false`)
	}
}

func TestLength(t *testing.T) {
	s := Cons(Int(42), Cons(Int(7), Nil))
	if Length(s) != 2 {
		t.Fatalf(`expected Length(s) == 2, got %s`, Length(s))
	}
}
