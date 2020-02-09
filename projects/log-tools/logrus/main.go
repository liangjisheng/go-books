package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func main() {
	// demo1()
	// demo2()
	// jsonOutput()
	// thirdFormat()
	// hook()
	redisHook()
}

func demo1() {
	// 默认的级别为InfoLevel 所以为了能看到Trace和Debug日志
	// 我们在main函数第一行设置日志级别为TraceLevel
	logrus.SetLevel(logrus.TraceLevel)

	logrus.Trace("trade msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	// logrus.Fatal("fatal msg")
	// logrus.Panic("panic msg")

	// 设置在输出日志中添加文件名和方法信息
	logrus.SetReportCaller(true)

	// 有时候需要在输出中添加一些字段，可以通过调用logrus.WithField和logrus.WithFields实现
	logrus.WithFields(logrus.Fields{
		"name": "ljs",
		"age":  24,
	}).Info("info msg")

	// 如果在一个函数中的所有日志都需要添加某些字段 可以使用WithFields的返回值
	// 例如在 Web 请求的处理器中 日志都要加上user_id和ip字段
	requestLogger := logrus.WithFields(logrus.Fields{
		"user_id": 10010,
		"ip":      "127.0.0.1",
	})

	requestLogger.Info("info msg")
	requestLogger.Error("error msg")
}

func demo2() {
	// 同时将日志写到bytes.Buffer、标准输出和文件中
	writer1 := &bytes.Buffer{}
	writer2 := os.Stdout
	writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("create file log.txt failed: %v", err)
	}

	logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
	logrus.Info("info msg")
	fmt.Print(writer1.String())
}

func jsonOutput() {
	// logrus支持两种日志格式 文本和 JSON 默认为文本格式
	// 可以通过logrus.SetFormatter设置日志格式
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	// logrus.Fatal("fatal msg")
	// logrus.Panic("panic msg")
}

func thirdFormat() {
	// 除了内置的TextFormatter和JSONFormatter，还有不少第三方格式支持
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})

	logrus.Info("info msg")

	logrus.SetFormatter(&nested.Formatter{
		// HideKeys:        true,
		TimestampFormat: time.RFC3339,
		FieldsOrder:     []string{"name", "age"},
	})

	logrus.WithFields(logrus.Fields{
		"name": "dj",
		"age":  18,
	}).Info("info msg")
}
