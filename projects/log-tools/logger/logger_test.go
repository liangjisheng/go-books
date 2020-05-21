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
		time.Sleep(200 * time.Millisecond)
	}
}

func BenchmarkLog(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("info")
	}
}
