package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
	"github.com/d11wtq/bijou/test"
	"testing"
)

func TestRunWithValidInput(t *testing.T) {
	res, err := Run(
		`(def head
		   (fn (hd & tl)
		     hd))

		 (head 42 7 23)`,
		test.FakeEnv(),
	)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if res != Int(42) {
		t.Fatalf(`expected res == Int(42), got %s`, res)
	}
}

func TestRunWithAccessToEnv(t *testing.T) {
	fun := &Func{
		Env:    test.FakeEnv(),
		Params: EmptyList,
		Body:   EmptyList.Append(Int(42)),
	}

	env := test.FakeEnv()
	env.Def("answer-to-life", fun)

	res, err := Run(
		`(answer-to-life)`,
		env,
	)
	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}
	if res != Int(42) {
		t.Fatalf(`expected res == Int(42), got %s`, res)
	}
}

func TestRunWithInvalidInput(t *testing.T) {
	res, err := Run(
		`(def head
		   (fn wat?`,
		test.FakeEnv(),
	)
	if err == nil {
		t.Fatalf(`expected err != nil, got nil`)
	}
	if res != nil {
		t.Fatalf(`expected res == nil, got %s`, res)
	}
}
