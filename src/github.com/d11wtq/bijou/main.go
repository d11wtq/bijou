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

	(def coalesce
	  (fn (x)
	    (unless (= () x)
		  (or (head x)
		      (coalesce (tail x))))))

	(coalesce (list nil nil 42 nil nil nil))

	(def test
	  (fn (x y & z)
	    z))

	(tail (test 1 2 3 4 5))


	"This is an example string.

	It spans several lines.

	It supports tabs:

	a\tb\tc\td

	It allows escaped strings:

	\"Like this\"

	Neat!"
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
