package main

import (
	"fmt"
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
)

func main() {
	env := core.RootEnv()
	src := `
	(def any?
	  (fn (x) x))

	(def coalesce
	  (fn (x)
	    (if (any? x)
		  (if (head x)
		    (head x)
			(coalesce (tail x))))))

	(coalesce (list nil nil nil 42 nil 7))
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
