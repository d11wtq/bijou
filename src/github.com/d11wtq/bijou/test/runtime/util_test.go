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

	v = EmptyList.Append(Int(42))
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

func TestTypeName(t *testing.T) {
	if TypeName(NilType) != "nil" {
		t.Fatalf(
			`expected TypeName(NilType) == "nil", got %s`,
			TypeName(NilType),
		)
	}
	if TypeName(BooleanType) != "boolean" {
		t.Fatalf(
			`expected TypeName(BooleanType) == "boolean", got %s`,
			TypeName(BooleanType),
		)
	}
	if TypeName(IntType) != "integer" {
		t.Fatalf(
			`expected TypeName(IntType) == "integer", got %s`,
			TypeName(IntType),
		)
	}
	if TypeName(SymbolType) != "symbol" {
		t.Fatalf(
			`expected TypeName(SymbolType) == "symbol", got %s`,
			TypeName(SymbolType),
		)
	}
	if TypeName(SequenceType) != "sequence" {
		t.Fatalf(
			`expected TypeName(SequenceType) == "sequence", got %s`,
			TypeName(SequenceType),
		)
	}
	if TypeName(StringType) != "string" {
		t.Fatalf(
			`expected TypeName(StringType) == "string", got %s`,
			TypeName(StringType),
		)
	}
	if TypeName(ConsType) != "cons" {
		t.Fatalf(
			`expected TypeName(ConsType) == "cons", got %s`,
			TypeName(ConsType),
		)
	}
	if TypeName(ListType) != "list" {
		t.Fatalf(
			`expected TypeName(ListType) == "list", got %s`,
			TypeName(ListType),
		)
	}
	if TypeName(FuncType) != "function" {
		t.Fatalf(
			`expected TypeName(FuncType) == "function", got %s`,
			TypeName(FuncType),
		)
	}
	if TypeName(MacroType) != "macro" {
		t.Fatalf(
			`expected TypeName(MacroType) == "macro", got %s`,
			TypeName(MacroType),
		)
	}
}
