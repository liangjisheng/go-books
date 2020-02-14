package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

// TailMgr ...
type TailMgr struct {
	// 因为我们的agent可能是读取多个日志文件，这里通过存储为一个map
	tailObjMap map[string]*TailObj
	lock       sync.Mutex
}

// TailObj ...
type TailObj struct {
	tail     *tail.Tail // 这里是每个读取日志文件的对象
	offset   int64      // 记录当前位置
	secLimit *SecondLimit
	// filename string
	logConf  logConfig
	exitChan chan bool
}

var tailMgr *TailMgr
var waitGroup sync.WaitGroup

// NewTailMgr ...
func NewTailMgr() *TailMgr {
	tailMgr = &TailMgr{
		tailObjMap: make(map[string]*TailObj, 16),
	}
	return tailMgr
}

// AddLogFile ...
func (t *TailMgr) AddLogFile(conf logConfig) (err error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	_, ok := t.tailObjMap[conf.LogPath]
	if ok {
		err = fmt.Errorf("duplicate filename %s", conf.LogPath)
		return
	}
	tail, err := tail.TailFile(conf.LogPath, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})

	tailobj := &TailObj{
		secLimit: NewSecondLimit(int32(conf.SendRate)),
		logConf:  conf,
		offset:   0,
		tail:     tail,
		exitChan: make(chan bool, 1),
	}

	t.tailObjMap[conf.LogPath] = tailobj
	logs.Info("map [%s]", t.tailObjMap)
	go tailobj.readLog()
	return
}

func (t *TailMgr) reloadConfig(logConfArr []logConfig) (err error) {
	for _, conf := range logConfArr {
		tailObj, ok := t.tailObjMap[conf.LogPath]
		if !ok {
			logs.Debug("conf:%v -- tailobj:%v", conf, tailObj)
			err = t.AddLogFile(conf)
			if err != nil {
				logs.Error("add log file failed,err:%v", err)
				continue
			}
			continue
		}
		tailObj.logConf = conf
		t.tailObjMap[conf.LogPath] = tailObj
		logs.Info(t.tailObjMap)
	}

	// 处理删除的日志收集配置
	for key, tailObj := range t.tailObjMap {
		var found = false
		for _, newValue := range logConfArr {
			if key == newValue.LogPath {
				found = true
				break
			}
		}
		if found == false {
			logs.Warn("log path:%s is remove", key)
			tailObj.exitChan <- true
			delete(t.tailObjMap, key)
		}
	}
	return
}

// Process ...
func (t *TailMgr) Process() {
	logChan := GetLogConf()
	for conf := range logChan {
		logs.Debug("log conf: %v", conf)
		var logConfArr []logConfig
		err := json.Unmarshal([]byte(conf), &logConfArr)
		if err != nil {
			logs.Error("unmarshal failed,err:%v conf:%v", err, conf)
			continue
		}
		logs.Debug("unmarshal succ conf:%v", logConfArr)
		err = t.reloadConfig(logConfArr)
		if err != nil {
			logs.Error("realod config from etcd failed err:%v", err)
			continue
		}
		logs.Debug("reaload from etcd success,config:%v", logConfArr)
	}

	// 开启线程去读日志文件
	// for _, tailObj := range t.tailObjMap {
	// 	waitGroup.Add(1)
	// 	go tailObj.readLog()
	// }
}

func (t *TailObj) readLog() {
	// 读取每行日志内容
	for line := range t.tail.Lines {
		if line.Err != nil {
			logs.Error("read line failed,err:%v", line.Err)
			continue
		}
		fmt.Println("text:", line.Text)
		str := strings.TrimSpace(line.Text)
		if len(str) == 0 || str[0] == '\n' {
			continue
		}

		kafkaSender.addMessage(line.Text, t.logConf.Topic)
		t.secLimit.Add(1)
		t.secLimit.Wait()

		select {
		case <-t.exitChan:
			logs.Warn("tail obj is exited,config:%v", t.logConf)
			return
		default:
		}
	}

	waitGroup.Done()
}

// RunServer ...
func RunServer() {
	tailMgr = NewTailMgr()
	// 这一部分是要调用tailf读日志文件推送到kafka中
	// for _, filename := range appConfig.LogFiles {
	// 	err := tailMgr.AddLogFile(filename)
	// 	if err != nil {
	// 		logs.Error("add log file failed,err:%v", err)
	// 		continue
	// 	}
	// }
	tailMgr.Process()
	waitGroup.Wait()
}
