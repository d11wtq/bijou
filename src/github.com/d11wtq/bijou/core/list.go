package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Basic list function, returns variadic args as a *List
func List(args *runtime.List) (runtime.Value, error) {
	return args, nil
}

// Return a new list appending a new head to a given list
func Cons(args *runtime.List) (runtime.Value, error) {
	var hd, tl runtime.Value
	if err := ReadArgs(args, &hd, &tl); err != nil {
		return nil, err
	}
	lst, ok := tl.(*runtime.List)
	if ok == false {
		return nil, &runtime.RuntimeError{"Bad data type: list required"}
	}

	return lst.Cons(hd), nil
}

// Return the head of the given list
func Head(args *runtime.List) (runtime.Value, error) {
	var lst runtime.Value
	if err := ReadArgs(args, &lst); err != nil {
		return nil, err
	}
	lst2, ok := lst.(*runtime.List)
	if ok == false {
		return nil, &runtime.RuntimeError{"Bad data type: list required"}
	}

	return lst2.Head(), nil
}

// Return the tail of the given list
func Tail(args *runtime.List) (runtime.Value, error) {
	var lst runtime.Value
	if err := ReadArgs(args, &lst); err != nil {
		return nil, err
	}
	lst2, ok := lst.(*runtime.List)
	if ok == false {
		return nil, &runtime.RuntimeError{"Bad data type: list required"}
	}

	return lst2.Tail(), nil
}
