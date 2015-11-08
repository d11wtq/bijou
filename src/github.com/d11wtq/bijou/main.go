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
	  "Return the first element in s for which predicate p is true."
	  (fn (p s)
	    (unless (empty? s)
		  (if (p (head s))
		    (head s)
			(some p (tail s))))))

	(some identity '(false nil 42 nil 7))

	(put "I like foo" 100)
	(put '(1 2 3) 4)
	(cons 42 "abc")
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
