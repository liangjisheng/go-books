package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type logConfig struct {
	Topic    string `json:"topic"`
	LogPath  string `json:"log_path"`
	Service  string `json:"service"`
	SendRate int    `json:"send_rate"`
}

func setConfig(key string) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"117.51.148.112:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer client.Close()

	logConfs := []logConfig{
		logConfig{
			Topic:    "nginx_log",
			LogPath:  "../tailf-demo/my.log",
			Service:  "nginx_log",
			SendRate: 100,
		},
	}
	logConfsBytes, err := json.Marshal(logConfs)
	if err != nil {
		log.Println(err.Error())
		return
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	_, err = client.Put(ctx, key, string(logConfsBytes))
	defer cancelFunc()
	if err != nil {
		log.Println("cli.Put", err.Error())
		return
	}

	ctx, cancelFunc = context.WithTimeout(context.Background(), 2*time.Second)
	res, err := client.Get(ctx, key)
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

func main() {
	key := "/logagent/169.254.214.88/log_config"
	setConfig(key)
	key = "/logtransfer/169.254.214.88/log_config"
	setConfig(key)
}
