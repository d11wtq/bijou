package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Basic list function, returns variadic args as a *List
func List(env runtime.Env, args *runtime.List) (runtime.Value, error) {
	return args, nil
}

// Return a new list appending a new head to a given list
func Cons(env runtime.Env, args *runtime.List) (runtime.Value, error) {
	var hd, tl runtime.Value
	if err := runtime.ReadArgs(args, &hd, &tl); err != nil {
		return nil, err
	}
	lst, ok := tl.(*runtime.List)
	if ok == false {
		return nil, &runtime.RuntimeError{"Bad data type: list required"}
	}

	return lst.Cons(hd), nil
}
