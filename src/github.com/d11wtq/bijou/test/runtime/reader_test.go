package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"strings"
	"testing"
)

func TestReadWithAnInt(t *testing.T) {
	v, s, err := Read("42")
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "" {
		t.Fatalf(`expected s == '', got %s`, s)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestReadWithAnIntFollowedByWhitespace(t *testing.T) {
	v, s, err := Read("42   ")
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "   " {
		t.Fatalf(`expected s == '   ', got %s`, s)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestReadWithAnIntPrecededByWhitespace(t *testing.T) {
	v, s, err := Read("   42")
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "" {
		t.Fatalf(`expected s == '', got %s`, s)
	}

	if v != Int(42) {
		t.Fatalf(`expected v == Int(42), got %s`, v)
	}
}

func TestReadWithAnIntNonsense(t *testing.T) {
	v, s, err := Read("42wat")
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	msg := "42wat"
	if !strings.Contains(err.Error(), msg) {
		t.Fatalf(`expected err to match "%s", got %s`, err)
	}
	if s != "42wat" {
		t.Fatalf(`expected s == '42wat', got %s`, s)
	}

	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestReadWithAnEmptyString(t *testing.T) {
	v, s, err := Read(`""`)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "" {
		t.Fatalf(`expected s == '', got %s`, s)
	}

	if v != String("") {
		t.Fatalf(`expected v == String(""), got %s`, v)
	}
}

func TestReadWithASimpleString(t *testing.T) {
	v, s, err := Read(`"hello, world!"`)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "" {
		t.Fatalf(`expected s == '', got %s`, s)
	}

	if v != String("hello, world!") {
		t.Fatalf(`expected v == String("hello, world!"), got %s`, v)
	}
}

func TestReadWithAStringSplitAcrossLines(t *testing.T) {
	v, s, err := Read("\"hello,\nnew line!\"")
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "" {
		t.Fatalf(`expected s == '', got %s`, s)
	}

	if v != String("hello,\nnew line!") {
		t.Fatalf(`expected v == String("hello,\nnew line!"), got %s`, v)
	}
}

func TestReadWithAStringContainingEscapedQuotes(t *testing.T) {
	v, s, err := Read(`"hello, \"beautiful\"!"`)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "" {
		t.Fatalf(`expected s == '', got %s`, s)
	}

	if v != String(`hello, "beautiful"!`) {
		t.Fatalf(`expected v == String('hello, "beautiful"!'), got %s`, v)
	}
}

func TestReadWithAStringContainingEscapedBackslashes(t *testing.T) {
	v, s, err := Read(`"hello, \\\"beautiful\"!"`)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "" {
		t.Fatalf(`expected s == '', got %s`, s)
	}

	if v != String(`hello, \"beautiful"!`) {
		t.Fatalf(`expected v == String('hello, \"beautiful"!'), got %s`, v)
	}
}

func TestReadWithAStringContainingLineFeedEscapes(t *testing.T) {
	v, s, err := Read(`"hello,\nnew line!"`)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "" {
		t.Fatalf(`expected s == '', got %s`, s)
	}

	if v != String("hello,\nnew line!") {
		t.Fatalf(`expected v == String("hello,\nnew line!"), got %s`, v)
	}
}

func TestReadWithAnEmptyList(t *testing.T) {
	v, s, err := Read("()")
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "" {
		t.Fatalf(`expected s == '', got %s`, s)
	}

	if v != EmptyList {
		t.Fatalf(`expected v == EmptyList, got %s`, v)
	}
}

func TestReadWithAnEmptyListFollowedByWhitespace(t *testing.T) {
	v, s, err := Read("()   ")
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "   " {
		t.Fatalf(`expected s == '', got %s`, s)
	}

	if v != EmptyList {
		t.Fatalf(`expected v == EmptyList, got %s`, v)
	}
}

func TestReadWithAnIntList(t *testing.T) {
	v, s, err := Read("(42 7)")
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "" {
		t.Fatalf(`expected s == '', got %s`, s)
	}

	lst, ok := v.(*List)
	if ok == false {
		t.Fatalf(`expected v.(*List), but not v.(*List)`)
	}

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

func TestReadWithANestedList(t *testing.T) {
	v, s, err := Read("((42) 7)")
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "" {
		t.Fatalf(`expected s == '', got %s`, s)
	}

	lst, ok := v.(*List)
	if ok == false {
		t.Fatalf(`expected v.(*List), but is not a *List`)
	}

	lst2, ok := lst.Head().(*List)
	if ok == false {
		t.Fatalf(`expected lst.Head().(*List), but is not a *List`)
	}

	if lst2.Head() != Int(42) {
		t.Fatalf(`expected lst2.Head() == Int(42), got %s`, lst2.Head())
	}

	if lst.Tail().Head() != Int(7) {
		t.Fatalf(
			`expected lst.Tail().Head() == Int(7), got %s`,
			lst.Tail().Head(),
		)
	}
}

func TestReadWithAnUnterminatedList(t *testing.T) {
	v, s, err := Read("(42 7")
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if s != "(42 7" {
		t.Fatalf(`expected s == '(42 7', got %s`, s)
	}
	if v != nil {
		t.Fatalf(`expected v == nil, got %s`, v)
	}
}

func TestReadWithASymbol(t *testing.T) {
	v, s, err := Read("xyz")
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "" {
		t.Fatalf(`expected s == '', got %s`, s)
	}
	if v != Symbol("xyz") {
		t.Fatalf(`expected v == Symbol("xyz"), got %s`, v)
	}
}

func TestReadWithASymbolFollowedByWhitespace(t *testing.T) {
	v, s, err := Read("xyz   ")
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != "   " {
		t.Fatalf(`expected s == '', got %s`, s)
	}
	if v != Symbol("xyz") {
		t.Fatalf(`expected v == Symbol("xyz"), got %s`, v)
	}
}

func TestReadWithASymbolFollowedByADelimiter(t *testing.T) {
	v, s, err := Read("xyz)")
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if s != ")" {
		t.Fatalf(`expected s == ')', got %s`, s)
	}
	if v != Symbol("xyz") {
		t.Fatalf(`expected v == Symbol("xyz"), got %s`, v)
	}
}

func TestReadSrcWithNoErrors(t *testing.T) {
	lst, err := ReadSrc(`  abc (x) 42  `)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if lst.Head() != Symbol("abc") {
		t.Fatalf(`expected lst.Head() == Symbol("abc"), got %s`, lst.Head())
	}

	lst2, ok := lst.Tail().Head().(*List)
	if ok == false {
		t.Fatalf(`expected lst.Tail().Head().(*List), but is not a *List`)
	}
	if lst2.Head() != Symbol("x") {
		t.Fatalf(`expected lst2.Head() == Symbol("x"), got %s`, lst2.Head())
	}
	if lst.Tail().Tail().Head() != Int(42) {
		t.Fatalf(
			`expected lst.Tail().Tail().Head() == Int(42), got %s`,
			lst.Tail().Tail().Head(),
		)
	}
}

func TestReadSrcWithErrors(t *testing.T) {
	lst, err := ReadSrc(`  abc (x  `)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if lst != nil {
		t.Fatalf(`expected lst == nil, got %s`, lst)
	}
}
