package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Return true if all arguments are equivalent.
// Usage: (= & args)
func Eq(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	if !args.Empty() {
		for !args.Tail().Empty() {
			if !args.Head().Eq(args.Tail().Head()) {
				return runtime.False, nil
			}
			args = args.Tail()
		}
	}

	return runtime.True, nil
}

// Return true if each argument is greater than the previous.
// Usage: (> & args)
func Gt(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	if !args.Empty() {
		for !args.Tail().Empty() {
			if !args.Head().Gt(args.Tail().Head()) {
				return runtime.False, nil
			}
			args = args.Tail()
		}
	}

	return runtime.True, nil
}

// Return true if each argument is less than the previous.
// Usage: (< & args)
func Lt(env runtime.Env, args runtime.Sequence) (runtime.Value, error) {
	if !args.Empty() {
		for !args.Tail().Empty() {
			if !args.Head().Lt(args.Tail().Head()) {
				return runtime.False, nil
			}
			args = args.Tail()
		}
	}

	return runtime.True, nil
}
