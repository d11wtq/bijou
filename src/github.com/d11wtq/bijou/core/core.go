package core

import (
	"github.com/d11wtq/bijou/runtime"
)

var root runtime.Env

// Initialize the root environment
func init() {
	root = runtime.NewScope(nil)

	// runtime functions
	root.Def("read", GoFunc(Read))
	root.Def("eval", GoFunc(Eval))
	root.Def("inspect", GoFunc(Inspect))

	// logical functions
	root.Def("identity", GoFunc(Identity))
	root.Def("not", GoFunc(Not))

	// comparison functions
	root.Def("=", GoFunc(Eq))

	// arithmetic functions
	root.Def("+", GoFunc(Add))
	root.Def("-", GoFunc(Sub))
	root.Def("/", GoFunc(Div))

	// list functions
	root.Def("list", GoFunc(List))
	root.Def("cons", GoFunc(Cons))

	// sequence functions
	root.Def("head", GoFunc(Head))
	root.Def("tail", GoFunc(Tail))
	root.Def("put", GoFunc(Put))
	root.Def("empty?", GoFunc(Empty))
}

// Get the initial root scope
func RootEnv() runtime.Env {
	return root.Extend()
}
