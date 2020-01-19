package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	filename := "../my.log"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	for i := 0; i < 100; i++ {
		content := fmt.Sprintf("%d\n", i)
		_, err := file.WriteString(content)
		if err != nil {
			continue
		}
		time.Sleep(2 * time.Second)
		fmt.Print(content)
	}
}
