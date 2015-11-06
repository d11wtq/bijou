package main

import (
	"fmt"
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
)

func main() {
	env := core.RootEnv()
	src := `
	(def unless
	  (macro (cond then)
	    (list (quote if)
		      (list (quote not) cond)
			  then)))

	(def or
	  (macro (a b)
	    (list (quote if)
		      a
			  a
			  b)))

	(def some
	  (fn (s)
	    (unless (empty? s)
		  (or (head s)
		      (some (tail s))))))

	(def x '(false nil 42 nil))
	(some x)
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
