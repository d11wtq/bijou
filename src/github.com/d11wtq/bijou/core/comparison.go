package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Logical equality
func Eq(env runtime.Env, args *runtime.List) (runtime.Value, error) {
	if args != runtime.EmptyList {
		for args.Next != runtime.EmptyList {
			if !args.Data.Eq(args.Next.Data) {
				return runtime.False, nil
			}
			args = args.Next
		}
	}

	return runtime.True, nil
}
