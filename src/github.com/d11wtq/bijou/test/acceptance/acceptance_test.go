package acceptance_test

import (
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"testing"
)

// Within running test t, execute string src and compare the result with wanted.
// If an error is returned, or the result is not what was expected, fail.
func AssertRunEqual(t *testing.T, src string, wanted runtime.Value) {
	AssertRunEqualWithEnv(t, src, wanted, core.RootEnv())
}

// Within running test t, execute string src and expect an error.
// If no error is returned, fail.
func AssertRunError(t *testing.T, src string) {
	AssertRunErrorWithEnv(t, src, core.RootEnv())
}

func AssertRunEqualWithEnv(
	t *testing.T,
	src string,
	wanted runtime.Value,
	env runtime.Env,
) {
	actual, err := runtime.Run(src, env)

	if err != nil {
		t.Fatalf(`expected err == nil, got %s`, err)
	}

	if !actual.Eq(wanted) {
		t.Fatalf(`expected %s => %s, got %s`, src, wanted, actual)
	}
}

func AssertRunErrorWithEnv(t *testing.T, src string, env runtime.Env) {
	actual, err := runtime.Run(src, env)

	if err == nil {
		t.Fatalf(`expected %s to error, but got %s`, src, actual)
	}
}
