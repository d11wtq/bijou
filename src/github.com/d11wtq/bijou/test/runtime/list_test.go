package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestListType(t *testing.T) {
	list := EmptyList
	if list.Type() != ListType {
		t.Fatalf(`expected list.Type() == ListType, got %s`, list.Type())
	}
}

func TestListEmpty(t *testing.T) {
	if !EmptyList.Empty() {
		t.Fatalf(`expected EmptyList.Empty(), got false`)
	}
}

func TestListPutPrependsNewHead(t *testing.T) {
	a := &List{
		&ConsCell{
			Int(42),
			&List{
				&ConsCell{
					Int(7),
					EmptyList,
				},
				nil,
			},
		},
		nil,
	}
	a.Last = a.Next.(*List)
	a.Next.(*List).Last = a.Next.(*List)

	b, err := a.Put(Int(7))

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	b2, ok := b.(*List)
	if ok == false {
		t.Fatalf(`expected b.(*List), got false`)
	}

	if b2.Data != Int(7) {
		t.Fatalf(`expected b2.Data == Int(7), got %s`, b2.Data)
	}

	if b2.Last != a.Last {
		t.Fatalf(`expected b2.Last == a.Last, got %s`, b2.Last)
	}

	if b2.Next.(*List).Data != Int(42) {
		t.Fatalf(
			`expected b2.Next.Data == Int(42), got %s`,
			b2.Next.(*List).Data,
		)
	}
}

func TestListAppendWithEmptyListIsPut(t *testing.T) {
	a := EmptyList
	b := a.Append(Int(7))

	if b.Data != Int(7) {
		t.Fatalf(`expected b.Data == Int(7), got %s`, b.Data)
	}

	if b.Last != b {
		t.Fatalf(`expected b.Last == b, got %s`, b.Last)
	}

	if b.Next.(*List) != EmptyList {
		t.Fatalf(`expected b.Next == EmptyList, got %s`, b.Next)
	}
}

func TestListAppendAddsANewValueAtTheEnd(t *testing.T) {
	a := &List{
		&ConsCell{
			Int(42),
			EmptyList,
		},
		nil,
	}
	a.Last = a

	b := a.Append(Int(7))

	if b.Data != Int(42) {
		t.Fatalf(`expected b.Data == Int(42), got %s`, b.Data)
	}

	if b.Last != b.Next.(*List) {
		t.Fatalf(`expected b.Last == b.Next, got %s`, b.Last)
	}

	if b.Last.Data != Int(7) {
		t.Fatalf(`expected b.Last.Data == Int(7), got %s`, b.Last.Data)
	}
}

func TestListEqWithEmptyLists(t *testing.T) {
	if !EmptyList.Eq(EmptyList) {
		t.Fatalf(`expected EmptyList.Eq(EmptyList), got false`)
	}
}

func TestListEqWithOneEmptyList(t *testing.T) {
	a := EmptyList
	b := EmptyList.Append(Int(42))

	if a.Eq(b) {
		t.Fatalf(`expected !a.Eq(b), got true`)
	}
	if b.Eq(a) {
		t.Fatalf(`expected !b.Eq(a), got true`)
	}
}

func TestListEqWithEquivalentLists(t *testing.T) {
	a := EmptyList.Append(Int(42)).Append(Int(7))
	b := EmptyList.Append(Int(42)).Append(Int(7))

	if !a.Eq(b) {
		t.Fatalf(`expected a.Eq(b), got false`)
	}
	if !b.Eq(a) {
		t.Fatalf(`expected b.Eq(a), got false`)
	}
}

func TestListEqWithEquivalentListsRecursive(t *testing.T) {
	a := EmptyList.
		Append(Int(7)).
		Append((&List{}).Append(Int(1))).
		Append(Int(42))
	b := EmptyList.
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
	a := EmptyList.Append(Int(7)).Append(Int(42))
	b := EmptyList.Append(Int(7))

	if a.Eq(b) {
		t.Fatalf(`expected !a.Eq(b), got true`)
	}
	if b.Eq(a) {
		t.Fatalf(`expected !b.Eq(a), got true`)
	}
}

func TestListEmptyListEvalItself(t *testing.T) {
	lst := EmptyList
	env := test.FakeEnv()

	v, err := lst.Eval(env)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if v != EmptyList {
		t.Fatalf(`expected v == EmptyList, got %s`, v)
	}
}

func TestListEmptyListHasNoTailOrHead(t *testing.T) {
	lst := EmptyList
	if lst.Head() != Nil {
		t.Fatalf(`expected lst.Head() == Nil, got %s`, lst.Head())
	}

	if lst.Tail() != EmptyList {
		t.Fatalf(
			`expected EmptyList.Tail() == EmptyList, got %s`,
			lst.Tail(),
		)
	}
}

func TestListTailIsAList(t *testing.T) {
	lst := EmptyList.Append(Int(42)).Append(Int(7))
	_, ok := lst.Tail().(*List)
	if ok == false {
		t.Fatalf(`expected lst.Tail().(*List), got false`)
	}
}
