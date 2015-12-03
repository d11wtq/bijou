package core

import (
	"github.com/d11wtq/bijou/runtime"
	"os"
)

var root runtime.Env

// Initialize the root environment
func init() {
	root = runtime.NewScope(nil)

	// builtin ports
	root.Def("stdin", GoIoPort(os.Stdin, nil))
	root.Def("stdout", GoIoPort(nil, os.Stdout))
	root.Def("stderr", GoIoPort(nil, os.Stderr))

	// runtime functions
	root.Def("read", GoFunc(Read))
	root.Def("eval", GoFunc(Eval))
	root.Def("apply", GoFunc(Apply))
	root.Def("inspect", GoFunc(Inspect))

	// logical functions
	root.Def("identity", GoFunc(Identity))
	root.Def("not", GoFunc(Not))

	// comparison functions
	root.Def("=", GoFunc(Eq))
	root.Def(">", GoFunc(Gt))
	root.Def("<", GoFunc(Lt))

	// arithmetic functions
	root.Def("+", GoFunc(Add))
	root.Def("-", GoFunc(Sub))
	root.Def("*", GoFunc(Mul))
	root.Def("/", GoFunc(Div))

	// list functions
	root.Def("list", GoFunc(List))
	root.Def("cons", GoFunc(Cons))

	// sequence functions
	root.Def("head", GoFunc(Head))
	root.Def("tail", GoFunc(Tail))
	root.Def("put", GoFunc(Put))
	root.Def("empty?", GoFunc(Empty))

	// port functions
	root.Def("send!", GoFunc(PortWrite))
	root.Def("receive!", GoFunc(PortAccept))
	root.Def("read!", GoFunc(PortRead))

	// concurrency functions
	root.Def("spawn", GoFunc(Spawn))
}

// Get the initial root scope
func RootEnv() runtime.Env {
	return root.Extend()
}
