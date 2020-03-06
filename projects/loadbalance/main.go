package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"loadbalance/balance"
)

func main() {
	var insts []*balance.Instance
	for i := 0; i < 16; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		one := balance.NewInstance(host, 8080)
		insts = append(insts, one)
	}

	var balanceName = balance.RandomBalancer
	if len(os.Args) > 1 {
		balanceName = os.Args[1]
	}

	for {
		inst, err := balance.DoBalance(balanceName, insts)
		if err != nil {
			fmt.Println("no balance err", err)
			continue
		}
		fmt.Println(inst)
		time.Sleep(time.Second)
	}
}
