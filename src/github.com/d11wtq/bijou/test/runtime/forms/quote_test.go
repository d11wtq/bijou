package forms_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestQuoteReturnsArgumentUnevaluated(t *testing.T) {
	form := EmptyList.Cons(Symbol("x")).Cons(Symbol("quote"))
	v, err := form.Eval(test.FakeEnv())

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if v != Symbol("x") {
		t.Fatalf(`expected v == Symbol("x"), got %s`, v)
	}
}
