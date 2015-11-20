package acceptance_test

import (
	"github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestAdd(t *testing.T) {
	AssertRunEqual(t, "(+)", runtime.Int(0))
	AssertRunEqual(t, "(+ 42)", runtime.Int(42))
	AssertRunEqual(t, "(+ 42 5)", runtime.Int(47))
	AssertRunError(t, "(+ 42 'x)")
}

func TestSub(t *testing.T) {
	AssertRunEqual(t, "(-)", runtime.Int(0))
	AssertRunEqual(t, "(- 42)", runtime.Int(-42))
	AssertRunEqual(t, "(- 42 5)", runtime.Int(37))
	AssertRunError(t, "(- 42 'x)")
}

func TestMul(t *testing.T) {
	AssertRunEqual(t, "(*)", runtime.Int(1))
	AssertRunEqual(t, "(* 42)", runtime.Int(42))
	AssertRunEqual(t, "(* 42 5)", runtime.Int(210))
	AssertRunError(t, "(* 42 'x)")
}

func TestDiv(t *testing.T) {
	AssertRunEqual(t, "(/)", runtime.Int(1))
	AssertRunEqual(t, "(/ 42)", runtime.Int(0))
	AssertRunEqual(t, "(/ 42 5)", runtime.Int(8))
	AssertRunError(t, "(/ 42 'x)")
}
