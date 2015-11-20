package acceptance_test

import (
	"github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestList(t *testing.T) {
	AssertRunEqual(t, "(list)", runtime.EmptyList)
	AssertRunEqual(
		t, "(list 42 7)",
		runtime.
			EmptyList.
			Append(runtime.Int(42)).
			Append(runtime.Int(7)),
	)
}

func TestCons(t *testing.T) {
	AssertRunEqual(
		t, "(cons 42 '())",
		runtime.Cons(
			runtime.Int(42),
			runtime.EmptyList,
		),
	)

	AssertRunEqual(
		t, "(cons 42 '(7))",
		runtime.Cons(
			runtime.Int(42),
			runtime.EmptyList.Append(runtime.Int(7)),
		),
	)

	AssertRunEqual(
		t, `(cons 42 "abc")`,
		runtime.Cons(
			runtime.Int(42),
			runtime.String("abc"),
		),
	)

	AssertRunError(t, "(cons)")
	AssertRunError(t, "(cons 42)")
	AssertRunError(t, "(cons 42 7)")
}
