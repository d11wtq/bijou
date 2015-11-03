package main

import (
	"fmt"
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
)

func main() {
	env := core.RootEnv()
	src := `
	(def coalesce
	  (fn (x)
	    (if (= () x)
		  nil
		  (if (head x)
		    (head x)
		    (coalesce (tail x))))))

	(coalesce (list nil nil nil nil nil nil))
	(= (list 42 ()) (list 42 ()))
	true
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
