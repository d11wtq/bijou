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
	if !EmptyList.Eq(EmptyList) {
		t.Fatalf(`expected EmptyList.Eq(EmptyList), got false`)
	}
}

func TestListEqWithOneEmptyList(t *testing.T) {
	a := EmptyList
	b := EmptyList.Cons(Int(42))

	if a.Eq(b) {
		t.Fatalf(`expected !a.Eq(b), got true`)
	}
	if b.Eq(a) {
		t.Fatalf(`expected !b.Eq(a), got true`)
	}
}

func TestListEqWithEquivalentLists(t *testing.T) {
	a := EmptyList.Cons(Int(42)).Cons(Int(7))
	b := EmptyList.Cons(Int(42)).Cons(Int(7))

	if !a.Eq(b) {
		t.Fatalf(`expected a.Eq(b), got false`)
	}
	if !b.Eq(a) {
		t.Fatalf(`expected b.Eq(a), got false`)
	}
}

func TestListEqWithEquivalentListsRecursive(t *testing.T) {
	a := EmptyList.
		Cons(Int(42)).
		Cons(EmptyList.Cons(Int(1))).
		Cons(Int(7))
	b := EmptyList.
		Cons(Int(42)).
		Cons(EmptyList.Cons(Int(1))).
		Cons(Int(7))

	if !a.Eq(b) {
		t.Fatalf(`expected a.Eq(b), got false`)
	}
	if !b.Eq(a) {
		t.Fatalf(`expected b.Eq(a), got false`)
	}
}

func TestListEqWithDifferentLengths(t *testing.T) {
	a := EmptyList.Cons(Int(42)).Cons(Int(7))
	b := EmptyList.Cons(Int(7))

	if a.Eq(b) {
		t.Fatalf(`expected !a.Eq(b), got true`)
	}
	if b.Eq(a) {
		t.Fatalf(`expected !b.Eq(a), got true`)
	}
}

func TestListEmptyListEvalItself(t *testing.T) {
	var list *List = nil

	if EmptyList != list {
		t.Fatalf(`expected EmptyList == list, got %s`, EmptyList)
	}

	env := FakeEnv()

	if v, err := list.Eval(env); err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	} else if v != EmptyList {
		t.Fatalf(`expected v == EmptyList, got %s`, v)
	}
}

func TestListEmptyListHasNoTailOrHead(t *testing.T) {
	if EmptyList.Head() != Nil {
		t.Fatalf(`expected EmptyList.Head() == Nil, got %s`, EmptyList.Head())
	}

	if EmptyList.Tail() != EmptyList {
		t.Fatalf(
			`expected EmptyList.Tail() == EmptyList, got %s`,
			EmptyList.Tail(),
		)
	}
}

func TestListPut(t *testing.T) {
	list, err := EmptyList.Cons(Int(42)).Put(Int(7))

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if list.Head() != Int(7) {
		t.Fatalf(`expected list.Head() == Int(7), got %s`, list.Head())
	}
	if list.Tail().Head() != Int(42) {
		t.Fatalf(
			`expected list.Tail().Head() == Int(42), got %s`,
			list.Tail().Head(),
		)
	}
}

func TestListEmpty(t *testing.T) {
	if v := EmptyList.Empty(); v == false {
		t.Fatalf(`expected EmptyList.Empty(), got false`)
	}

	if v := EmptyList.Cons(Int(42)).Empty(); v == true {
		t.Fatalf(`expected !EmptyList.Cons(Int(42)).Empty(), got true`)
	}
}

func TestListConsAddsANewHead(t *testing.T) {
	list := EmptyList.Cons(Int(42)).Cons(Int(7))

	if list.Head() != Int(7) {
		t.Fatalf(`expected list.Head() == Int(7), got %s`, list.Head())
	} else if list.Tail().Head() != Int(42) {
		t.Fatalf(
			`expected list.Tail().Head() == Int(42), got %s`,
			list.Tail().Head(),
		)
	}
}

func TestListReverse(t *testing.T) {
	list := EmptyList.Cons(Int(42)).Cons(Int(7))
	rev := list.Reverse()

	if rev.Head() != Int(42) {
		t.Fatalf(`expected list.Head() == Int(42), got %s`, rev.Head())
	} else if rev.Tail().Head() != Int(7) {
		t.Fatalf(
			`expected rev.Tail().Head() == Int(7), got %s`,
			rev.Tail().Head(),
		)
	}
}
