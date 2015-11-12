package main

import (
	"fmt"
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
)

func main() {
	src := `())`

	res, err := runtime.Run(src, core.RootEnv())

	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println("error:", err)
	}
}
