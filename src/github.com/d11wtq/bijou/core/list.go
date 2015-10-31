package core

import (
	"github.com/d11wtq/bijou/runtime"
)

// Basic list function, returns variadic args as a *List
func List(args *runtime.List) (runtime.Value, error) {
	return args, nil
}

// Return the head of the given list
func Head(args *runtime.List) (runtime.Value, error) {
	// FIXME: Validate arity, extract args easier
	// FIXME: Need a way to register functions with a fixed number of args
	// FIXME: Needs to do the BadArity error
	return args.Head().(*runtime.List).Head(), nil
}

// Return the tail of the given list
func Tail(args *runtime.List) (runtime.Value, error) {
	return args.Head().(*runtime.List).Tail(), nil
}
