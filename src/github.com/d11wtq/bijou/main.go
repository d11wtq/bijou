package main

import (
	"fmt"
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"io/ioutil"
	"os"
)

// Interpreter exit codes
const (
	ExitOK  = 0
	ExitErr = 255
)

// Terminate with exit code 255 on error
func haltIf(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(ExitErr)
	}
}

// Command line entry point.
// Usage: ./bin/bijou < source.bjx
func main() {
	src, err := ioutil.ReadAll(os.Stdin)
	haltIf(err)
	res, err := runtime.Run(string(src), core.RootEnv())
	haltIf(err)
	fmt.Fprintln(os.Stdout, res)
	os.Exit(ExitOK)
}
