package main

import (
	"sync"

	"github.com/gohouse/gorose"
)

var (
	once  sync.Once
	engin *gorose.Engin
)

func init() {
	BootGorose()
	UserInit()
}

func main() {
	BootGin()
}
