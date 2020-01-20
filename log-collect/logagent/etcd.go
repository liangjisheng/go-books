package main

import (
	"context"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/clientv3"
)

// Client etcd client
var Client *clientv3.Client
var logConfChan chan string

func initEtcd(addr []string, keyfmt string, timeout time.Duration) (err error) {
	var keys []string
	for _, ip := range ipArrays {
		// keyfmt = /logagent/%s/log_config
		keys = append(keys, fmt.Sprintf(keyfmt, ip))
	}

	logConfChan = make(chan string, 10)
	logs.Debug("etcd watch key:%v timeout:%v", keys, timeout)

	Client, err = clientv3.New(clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error(err.Error())
		return
	}
	logs.Debug("init etcd success")
	waitGroup.Add(1)

	for _, key := range keys {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		// 从etcd中获取要收集日志的信息
		resp, err := Client.Get(ctx, key)
		cancel()
		if err != nil {
			logs.Warn("get key %s failed,err:%v", key, err)
			continue
		}

		for _, ev := range resp.Kvs {
			logs.Debug("%q : %q\n", ev.Key, ev.Value)
			logConfChan <- string(ev.Value)
		}
	}
	go WatchEtcd(keys)
	return
}

// WatchEtcd ...
func WatchEtcd(keys []string) {
	defer waitGroup.Done()

	// 这里用于检测当需要收集的日志信息更改时及时更新
	var watchChans []clientv3.WatchChan
	for _, key := range keys {
		rch := Client.Watch(context.Background(), key)
		watchChans = append(watchChans, rch)
	}

	for {
		for _, watchC := range watchChans {
			select {
			case wresp := <-watchC:
				for _, ev := range wresp.Events {
					logs.Debug("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
					logConfChan <- string(ev.Kv.Value)
				}
			default:
			}
		}
		time.Sleep(time.Second)
	}
}

// GetLogConf ...
func GetLogConf() chan string {
	return logConfChan
}
