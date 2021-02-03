package main

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func demo1() {
	var g errgroup.Group

	// 启动第一个子任务,它执行成功
	g.Go(func() error {
		time.Sleep(1 * time.Second)
		fmt.Println("exec #1")
		// return errors.New("failed to exec #1")
		return nil
	})

	// 启动第二个子任务，它执行失败
	g.Go(func() error {
		time.Sleep(2 * time.Second)
		fmt.Println("exec #2")
		return errors.New("failed to exec #2")
	})

	// 启动第三个子任务，它执行成功
	g.Go(func() error {
		time.Sleep(3 * time.Second)
		fmt.Println("exec #3")
		return nil
	})

	// 等待三个任务都完成后返回，如果有多个子任务返回error，则会返回第一个error，所有子任务执行成功则返回nil
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully exec all")
	} else {
		fmt.Println("failed:", err)
	}
}
