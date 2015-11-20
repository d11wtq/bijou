package main

import (
	"fmt"
	"github.com/d11wtq/bijou/core"
	"github.com/d11wtq/bijou/runtime"
	"io"
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

// Return the source string of the input
func readSrc() string {
	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])
		haltIf(err)
		return readIo(file)
	} else {
		return readIo(os.Stdin)
	}
}

// Return a source string from an io.Reader
func readIo(s io.Reader) string {
	src, err := ioutil.ReadAll(s)
	haltIf(err)
	return string(src)
}

// Command line entry point.
// Usage: ./bin/bijou source.bjx
func main() {
	_, err := runtime.Run(readSrc(), core.RootEnv())
	haltIf(err)
	os.Exit(ExitOK)
}
