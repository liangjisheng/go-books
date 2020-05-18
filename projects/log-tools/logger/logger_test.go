package logger_test

import (
	"testing"
	"time"

	"logger"
)

func TestLog(t *testing.T) {
	for {
		logger.Info("info")
		logger.Debug("debug")
		logger.Warn("warn")
		logger.Error("error")

		logger.JInfo("info")
		logger.JDebug("debug")
		logger.JWarn("warn")
		logger.JError("error")

		time.Sleep(200 * time.Millisecond)
	}
}
