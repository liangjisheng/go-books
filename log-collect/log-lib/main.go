package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/logs"
)

func main() {
	config := make(map[string]interface{})
	config["filename"] = "./logcollect.log"
	config["level"] = logs.LevelTrace
	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed,err：", err)
		return
	}
	logs.SetLogger(logs.AdapterFile, string(configStr))
	logs.Debug("this is a debug,my name is %s", "stu01")
	logs.Info("this is a info,my name is %s", "stu02")
	logs.Trace("this is trace my name is %s", "stu03")
	logs.Warn("this is a warn my name is %s", "stu04")
}
