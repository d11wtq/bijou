package acceptance_test

import (
	"github.com/d11wtq/bijou/runtime"
	"testing"
)

func TestEq(t *testing.T) {
	AssertRunEqual(t, "(=)", runtime.True)
	AssertRunEqual(t, "(= 42)", runtime.True)
	AssertRunEqual(t, "(= 42 7)", runtime.False)
	AssertRunEqual(t, "(= 42 42 42)", runtime.True)
	AssertRunEqual(t, "(= 42 7 42 42)", runtime.False)
}

func TestGt(t *testing.T) {
	AssertRunEqual(t, "(>)", runtime.True)
	AssertRunEqual(t, "(> 42)", runtime.True)
	AssertRunEqual(t, "(> 42 7)", runtime.True)
	AssertRunEqual(t, "(> 42 41 40)", runtime.True)
	AssertRunEqual(t, "(> 42 7 41 40)", runtime.False)
}

func TestLt(t *testing.T) {
	AssertRunEqual(t, "(<)", runtime.True)
	AssertRunEqual(t, "(< 42)", runtime.True)
	AssertRunEqual(t, "(< 7 42)", runtime.True)
	AssertRunEqual(t, "(< 40 41 42)", runtime.True)
	AssertRunEqual(t, "(< 40 41 7 42)", runtime.False)
}
