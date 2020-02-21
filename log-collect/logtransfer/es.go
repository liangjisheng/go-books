package main

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/astaxie/beego/logs"
	"github.com/olivere/elastic"
)

var waitGroup sync.WaitGroup
var client *elastic.Client

type logConfig struct {
	Topic    string `json:"topic"`
	LogPath  string `json:"log_path"`
	Service  string `json:"service"`
	SendRate int    `json:"send_rate"`
}

func initEs(addr string) (err error) {
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(addr))
	if err != nil {
		logs.Error("connect to es error:%v", err)
		return
	}
	logs.Debug("conn to es success")
	return
}

func reloadKafka(topicArray []string) {
	for _, topic := range topicArray {
		kafkaMgr.AddTopic(topic)
	}
}

func reload() {
	// GetLogConf() 从channel中获topic信息，而这部分信息是从etcd放进去的
	for conf := range GetLogConf() {
		var logArray []logConfig
		var topicArray []string
		err := json.Unmarshal([]byte(conf), &logArray)
		if err != nil {
			logs.Error("unmarshal failed,err:%v conf:%v", err, conf)
			continue
		}
		for i := 0; i < len(logArray); i++ {
			topicArray = append(topicArray, logArray[i].Topic)
		}
		reloadKafka(topicArray)
	}
}

// Run ...
func Run(esThreadNum int) (err error) {
	go reload()
	for i := 0; i < esThreadNum; i++ {
		waitGroup.Add(1)
		go sendToEs()
	}
	waitGroup.Wait()
	return
}

// EsMessage ...
type EsMessage struct {
	Message string
}

func sendToEs() {
	// 从msgChan中读取日志内容并扔到elasticsearch中
	for msg := range GetMessage() {
		var esMsg EsMessage
		esMsg.Message = msg.line
		_, err := client.Index().
			Index(msg.topic).
			Type(msg.topic).
			BodyJson(esMsg).
			Do(context.Background())
		if err != nil {
			logs.Error("send to es failed,err:%v", err)
			continue
		}
		logs.Debug("send to es success")
	}
	waitGroup.Done()
}
