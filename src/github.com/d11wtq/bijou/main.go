package main

import (
	"fmt"
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
)

func main() {
	env := core.RootEnv()
	src := `
	(def forty-two
	  (fn () 42))

	(if (forty-two)
	  (head (list 1 2 3))
	  8)
	`

	app, err := runtime.ReadSrc(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := runtime.EvalDo(env, app)
	if err != nil {
		fmt.Println(err)
		return
	}
	// FIXME: There is a bug with evaluation of args
	fmt.Println(res)
}
