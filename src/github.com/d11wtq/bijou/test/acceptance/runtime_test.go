package acceptance_test

import (
	"github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestRead(t *testing.T) {
	AssertRunEqual(
		t, `(read "(1 2)")`,
		runtime.EmptyList.
			Append(runtime.Int(1)).
			Append(runtime.Int(2)),
	)
	AssertRunError(t, `(read "(1 2")`)
	AssertRunError(t, `(read 1 2)`)
	AssertRunError(t, `(read)`)
}

func TestEval(t *testing.T) {
	AssertRunEqual(t, `(eval '(head '(1 2 3)))`, runtime.Int(1))
	AssertRunError(t, `(eval)`)
	AssertRunError(t, `(eval 1 2)`)
}

func TestApply(t *testing.T) {
	AssertRunEqual(t, `(apply * '(2 3))`, runtime.Int(6))
	AssertRunError(t, `(apply * 2 3)`)
	AssertRunError(t, `(apply * 2)`)
	AssertRunError(t, `(apply *)`)
	AssertRunError(t, `(apply)`)
}

func TestInspect(t *testing.T) {
	AssertRunEqual(t, `(inspect '("a" "b"))`, runtime.String(`("a" "b")`))
	AssertRunEqual(t, `(inspect 42)`, runtime.String("42"))
	AssertRunError(t, `(inspect 1 2)`)
	AssertRunError(t, `(inspect)`)
}
