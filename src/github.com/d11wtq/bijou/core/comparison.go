package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Logical equality
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
