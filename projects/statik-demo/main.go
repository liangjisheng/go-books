//go:generate statik -src=./assets -include=*.txt
//go:generate go fmt statik/statik.go

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	_ "statik"

	"github.com/rakyll/statik/fs"
)

// Before building, run go generate.
func main() {
	statikFS, err := fs.New()
	if err != nil {
		fmt.Printf("fs.New() err: %v\n", err)
		os.Exit(1)
	}

	file, err := statikFS.Open("/tmp/b.txt")
	if err != nil {
		fmt.Printf("statikFS.Open err: %v\n", err)
		os.Exit(1)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("ioutil.ReadAll err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("content: %s\n", string(content))
}
