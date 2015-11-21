package acceptance_test

import (
	"github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestHead(t *testing.T) {
	AssertRunEqual(t, "(head '())", runtime.Nil)
	AssertRunEqual(t, "(head '(42 7 8))", runtime.Int(42))
	AssertRunError(t, "(head)")
	AssertRunError(t, "(head 42)")
	AssertRunError(t, "(head '() '())")
}

func TestTail(t *testing.T) {
	AssertRunEqual(t, "(tail '())", runtime.EmptyList)
	AssertRunEqual(
		t, "(tail '(42 7 8))",
		runtime.
			EmptyList.
			Append(runtime.Int(7)).
			Append(runtime.Int(8)),
	)
	AssertRunError(t, "(tail)")
	AssertRunError(t, "(tail 42)")
	AssertRunError(t, "(tail '() '())")
}

func TestPut(t *testing.T) {
	AssertRunEqual(
		t, "(put '() 42)",
		runtime.EmptyList.Append(runtime.Int(42)),
	)
	AssertRunEqual(
		t, "(put '(42 7) 2)",
		runtime.
			EmptyList.
			Append(runtime.Int(2)).
			Append(runtime.Int(42)).
			Append(runtime.Int(7)),
	)
	AssertRunError(t, "(put)")
	AssertRunError(t, "(put '())")
	AssertRunError(t, "(put 42)")
	AssertRunError(t, "(put '() 42 7)")
}

func TestEmpty(t *testing.T) {
	AssertRunEqual(t, "(empty? '())", runtime.True)
	AssertRunEqual(t, "(empty? '(42 7 8))", runtime.False)
	AssertRunEqual(t, `(empty? "")`, runtime.True)
	AssertRunError(t, "(empty?)")
	AssertRunError(t, "(empty? 42)")
	AssertRunError(t, "(empty? '() '())")
}
