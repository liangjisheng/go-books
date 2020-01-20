package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"117.51.148.112:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer client.Close()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = client.Put(ctx, "/logagent/169.254.214.88/log_config", "../tailf-demo/my.log")
	defer cancelFunc()
	if err != nil {
		log.Println("cli.Put", err.Error())
		return
	}

	ctx, cancelFunc = context.WithTimeout(context.Background(), 2*time.Second)
	res, err := client.Get(ctx, "/logagent/169.254.214.88/log_config")
	defer cancelFunc()
	if err != nil {
		log.Println("cli.Get", err.Error())
		return
	}
	if len(res.Kvs) == 0 {
		log.Println("the key logagent is not exist")
		return
	}

	for k, v := range res.Kvs {
		fmt.Println("res", k, string(v.Key), string(v.Value))
	}
}
