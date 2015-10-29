package main

import (
	"fmt"
	"github.com/d11wtq/bijou/runtime"
)

func main() {
	env := runtime.NewScope(nil)
	src := `
	(def forty-two
	  (fn () 42))

	(if (forty-two)
	  7
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
	fmt.Println(res)
}
