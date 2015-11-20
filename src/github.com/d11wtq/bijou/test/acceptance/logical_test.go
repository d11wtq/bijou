package acceptance_test

import (
	"github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestNot(t *testing.T) {
	AssertRunEqual(t, "(not nil)", runtime.True)
	AssertRunEqual(t, "(not 0)", runtime.False)
	AssertRunEqual(t, "(not true)", runtime.False)
	AssertRunEqual(t, "(not false)", runtime.True)
	AssertRunError(t, "(not)")
	AssertRunError(t, "(not 1 2)")
}

func TestIdentity(t *testing.T) {
	AssertRunEqual(t, "(identity nil)", runtime.Nil)
	AssertRunEqual(t, "(identity 42)", runtime.Int(42))
	AssertRunEqual(t, "(identity 'foo)", runtime.Symbol("foo"))
	AssertRunError(t, "(identity)")
	AssertRunError(t, "(identity 1 2)")
}
