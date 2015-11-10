package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Return argument v unchanged.
// Usage: (identity v)
func Identity(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var v runtime.Value
	if err := runtime.ReadArgs(args, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Return the logical inversion of v.
// Usage: (not x)
func Not(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	var v runtime.Value
	if err := runtime.ReadArgs(args, &v); err != nil {
		return nil, err
	}

	if v == runtime.Nil || v == runtime.False {
		return runtime.True, nil
	} else {
		return runtime.False, nil
	}
}
