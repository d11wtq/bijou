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
	root.Def("true", runtime.True)
	root.Def("false", runtime.False)

	// runtime functions
	root.Def("read", GoFunc(Read))
	root.Def("eval", GoFunc(Eval))

	// logical functions
	root.Def("not", GoFunc(Not))

	// comparison functions
	root.Def("=", GoFunc(Eq))

	// list functions
	root.Def("list", GoFunc(List))
	root.Def("cons", GoFunc(Cons))

	// sequence functions
	root.Def("head", GoFunc(Head))
	root.Def("tail", GoFunc(Tail))
}

// Get the initial root scope
func RootEnv() runtime.Env {
	return root
}
