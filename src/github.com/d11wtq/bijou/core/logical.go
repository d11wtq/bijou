package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Logical not
func Not(args *runtime.List) (runtime.Value, error) {
	var v runtime.Value
	if err := ReadArgs(args, &v); err != nil {
		return nil, err
	}

	if v == runtime.Nil || v == runtime.False {
		return runtime.True, nil
	} else {
		return runtime.False, nil
	}
}
