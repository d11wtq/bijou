package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestListType(t *testing.T) {
	list := &List{}
	if list.Type() != ListType {
		t.Fatalf(`expected list.Type() == ListType, got %s`, list.Type())
	}
}

func TestListEqWithEmptyLists(t *testing.T) {
	if !(&List{}).Eq(&List{}) {
		t.Fatalf(`expected EmptyList.Eq(EmptyList), got false`)
	}
}

func TestListEqWithOneEmptyList(t *testing.T) {
	a := &List{}
	b := (&List{}).Append(Int(42))

	if a.Eq(b) {
		t.Fatalf(`expected !a.Eq(b), got true`)
	}
	if b.Eq(a) {
		t.Fatalf(`expected !b.Eq(a), got true`)
	}
}

func TestListEqWithEquivalentLists(t *testing.T) {
	a := (&List{}).Append(Int(7)).Append(Int(42))
	b := (&List{}).Append(Int(7)).Append(Int(42))

	if !a.Eq(b) {
		t.Fatalf(`expected a.Eq(b), got false`)
	}
	if !b.Eq(a) {
		t.Fatalf(`expected b.Eq(a), got false`)
	}
}

func TestListEqWithEquivalentListsRecursive(t *testing.T) {
	a := (&List{}).
		Append(Int(7)).
		Append((&List{}).Append(Int(1))).
		Append(Int(42))
	b := (&List{}).
		Append(Int(7)).
		Append((&List{}).Append(Int(1))).
		Append(Int(42))

	if !a.Eq(b) {
		t.Fatalf(`expected a.Eq(b), got false`)
	}
	if !b.Eq(a) {
		t.Fatalf(`expected b.Eq(a), got false`)
	}
}

func TestListEqWithDifferentLengths(t *testing.T) {
	a := (&List{}).Append(Int(7)).Append(Int(42))
	b := (&List{}).Append(Int(7))

	if a.Eq(b) {
		t.Fatalf(`expected !a.Eq(b), got true`)
	}
	if b.Eq(a) {
		t.Fatalf(`expected !b.Eq(a), got true`)
	}
}

func TestListEmptyListEvalItself(t *testing.T) {
	lst := &List{}
	env := FakeEnv()

	v, err := lst.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != lst {
		t.Fatalf(`expected v == lst, got %s`, v)
	}
}

func TestListEmptyListHasNoTailOrHead(t *testing.T) {
	lst := &List{}
	if lst.Head() != Nil {
		t.Fatalf(`expected lst.Head() == Nil, got %s`, lst.Head())
	}

	if !lst.Tail().Empty() {
		t.Fatalf(
			`expected EmptyList.Tail().Empty(), got %s`,
			lst.Tail(),
		)
	}
}

func TestListPut(t *testing.T) {
	lst, err := (&List{}).Append(Int(42)).Put(Int(7))

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if lst.Head() != Int(7) {
		t.Fatalf(`expected lst.Head() == Int(7), got %s`, lst.Head())
	}
	if lst.Tail().Head() != Int(42) {
		t.Fatalf(
			`expected lst.Tail().Head() == Int(42), got %s`,
			lst.Tail().Head(),
		)
	}
}

func TestListAppend(t *testing.T) {
	lst := (&List{}).Append(Int(42)).Append(Int(7))
	if lst.Head() != Int(42) {
		t.Fatalf(`expected lst.Head() == Int(42), got %s`, lst.Head())
	}
	if lst.Tail().Head() != Int(7) {
		t.Fatalf(
			`expected lst.Tail().Head() == Int(7), got %s`,
			lst.Tail().Head(),
		)
	}
}

func TestListEmpty(t *testing.T) {
	lst := &List{}

	if !lst.Empty() {
		t.Fatalf(`expected lst.Empty(), got false`)
	}

	if lst.Append(Int(42)).Empty() {
		t.Fatalf(`expected !lst.Append(Int(42)).Empty(), got true`)
	}
}
