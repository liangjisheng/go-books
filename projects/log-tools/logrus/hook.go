package main

import (
	"io/ioutil"

	logredis "github.com/rogierlommers/logrus-redis-hook"
	"github.com/sirupsen/logrus"
)

// 钩子需要实现logrus.Hook接口
// github.com/sirupsen/logrus/hooks.go
// type Hook interface {
// 	Levels() []Level
// 	Fire(*Entry) error
// }
// Levels()方法返回感兴趣的日志级别，输出其他日志时不会触发钩子
// Fire是日志输出前调用的钩子方法

// AppHook ...
type AppHook struct {
	AppName string
}

// Levels ...
func (h *AppHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire ...
func (h *AppHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = h.AppName
	return nil
}

func hook() {
	h := &AppHook{AppName: "awesome-web"}
	logrus.AddHook(h)

	logrus.Info("info msg")
}

// logrus的第三方 Hook 很多，我们可以使用一些 Hook 将日志发送到 redis/mongodb 等存储中
// https://github.com/weekface/mgorus 将日志发送到 mongodb
// https://github.com/rogierlommers/logrus-redis-hook 将日志发送到 redis
// https://github.com/vladoatanasov/logrus_amqp 将日志发送到 ActiveMQ

func init() {
	hookConfig := logredis.HookConfig{
		Host:     "117.51.148.112",
		Port:     6379,
		Password: "password",
		Key:      "mykey",
		Format:   "v0",
		App:      "aweosome",
		Hostname: "117.51.148.112",
		TTL:      3600,
	}

	hook, err := logredis.NewHook(hookConfig)
	if err == nil {
		logrus.AddHook(hook)
	} else {
		logrus.Errorf("logredis error: %q", err)
	}
}

// 将日志发送到 redis
func redisHook() {
	logrus.Info("just some info logging...")

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"foo":    "bar",
		"this":   "that",
	}).Info("additional fields are being logged as well")

	logrus.SetOutput(ioutil.Discard)
	logrus.Info("This will only be sent to Redis")
}
