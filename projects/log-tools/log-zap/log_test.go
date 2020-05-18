package logger_test

import (
	"fmt"
	"testing"

	"logger"
)

func TestInitZapV2Logger(t *testing.T) {
	for i := 0; i < 2; i++ {
		logger.Debug(fmt.Sprintf("debug log %d", i+1))
		logger.Info(fmt.Sprintf("info log %d", i+1))
		// logger.Warn(fmt.Sprintf("warn log %d", i+1))
		// logger.Error(fmt.Sprintf("error log %d", i+1))
	}

	// appName := time.Now().Format("2006-01-02")
	// lg := NewLogger(SetAppName(appName), SetDevelopment(true), SetLevel(zap.DebugLevel))
	// for i := 0; i < 10; i++ {
	// 	lg.Debug(fmt.Sprint("debug log ", 1), zap.Int("line", 47))
	// 	lg.Info(fmt.Sprint("Info log ", 2), zap.Any("level", "1231231231"))
	// 	lg.Warn(fmt.Sprint("warn log ", 3), zap.String("level", `{"a":"4","b":"5"}`))
	// 	lg.Error(fmt.Sprint("err log ", 4), zap.String("level", `{"a":"7","b":"8"}`))
	// }
}
