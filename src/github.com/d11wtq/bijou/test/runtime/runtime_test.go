package runtime_test

import (
	. "github.com/d11wtq/bijou/runtime"
)

// Return an implementation of Env, suitable for testing
func FakeEnv() Env {
	return NewScope(nil)
}
