package main

import (
	"fmt"

	"github.com/robfig/cron"
)

func main() {
	c := cron.New()

	// @every 1m
	// @every 1h
	// @every 1m2s
	// @daily
	c.AddFunc("@every 1s", func() {
		fmt.Println("task start in 1 seconds")
	})

	c.Start()
	select {}
}
