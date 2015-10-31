package core

import (
	"github.com/d11wtq/bijou/runtime"
)

var root runtime.Env

// Initialize the root environment
func init() {
	root = runtime.NewScope(nil)

	// primitives
	root.Def("nil", runtime.Nil)

	// functions
	root.Def("list", GoFunc(List))
	root.Def("head", GoFunc(Head))
	root.Def("tail", GoFunc(Tail))
}

// Get the initial root scope
func RootEnv() runtime.Env {
	return root
}
