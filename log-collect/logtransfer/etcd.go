package main

import (
	"context"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"go.etcd.io/etcd/clientv3"
)

var etcdClient *clientv3.Client
var topicConfChan chan string

func initEtcd(addr []string, keyfmt string, ipArrays []string, timeout time.Duration) (err error) {
	var keys []string
	for _, ip := range ipArrays {
		keys = append(keys, fmt.Sprintf(keyfmt, ip))
	}

	topicConfChan = make(chan string, 8)
	logs.Debug("etcd watch key:%v, timeout:%d", keys, timeout)

	etcdClient, err = clientv3.New(clientv3.Config{
		Endpoints:   addr,
		DialTimeout: timeout,
	})
	if err != nil {
		logs.Warn("init etcd client failed, err:%v", err)
		return
	}

	logs.Debug("init etcd succ")
	waitGroup.Add(1)

	for _, key := range keys {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		///logagent/192.168.2.100/log_config
		resp, err := etcdClient.Get(ctx, key)
		cancel()
		if err != nil {
			logs.Warn("get key %s failed, err:%v", key, err)
			continue
		}

		for _, ev := range resp.Kvs {
			logs.Debug(" %q : %q\n", ev.Key, ev.Value)
			topicConfChan <- string(ev.Value)
		}
	}
	go WatchEtcd(keys)
	return
}

// WatchEtcd ...
func WatchEtcd(keys []string) {
	defer waitGroup.Done()
	var watchChans []clientv3.WatchChan
	for _, key := range keys {
		rch := etcdClient.Watch(context.Background(), key)
		watchChans = append(watchChans, rch)
	}

	for {
		for _, watchC := range watchChans {
			select {
			case wresp := <-watchC:
				for _, ev := range wresp.Events {
					logs.Debug("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
					topicConfChan <- string(ev.Kv.Value)
				}
			default:
			}
		}
		time.Sleep(time.Second)
	}
}

// GetLogConf ...
func GetLogConf() chan string {
	return topicConfChan
}
