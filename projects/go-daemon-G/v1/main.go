package main

import (
	"fmt"
	"time"
)

// Fibonacci ...
type Fibonacci struct {
	a, b int
}

// NewFibonacci ...
func NewFibonacci() *Fibonacci {
	return &Fibonacci{a: 0, b: 1}
}

// Run ...
func (f *Fibonacci) Run() {
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(f.b)
			f.a, f.b = f.b, f.a+f.b
		}
	}()
}

func main() {
	NewFibonacci().Run()
	select {}
}
