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
