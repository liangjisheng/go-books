package logger

import (
	"io"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var textLogger *zap.SugaredLogger
var jsonLogger *zap.SugaredLogger

func init() {
	cfg := zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "timestamp",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}

	// 设置一些基本日志格式 具体含义还比较好理解 直接看zap源码也不难懂
	textEncoder := zapcore.NewConsoleEncoder(cfg)
	jsonEncoder := zapcore.NewJSONEncoder(cfg)

	// 实现两个判断日志等级的 interface
	debugLevle := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel
	})
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.WarnLevel
	})
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.ErrorLevel
	})

	// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
	textDebugWriter := getWriter("./logs/trade-text-debug.log")
	textInfoWriter := getWriter("./logs/trade-text-info.log")
	textWarnWriter := getWriter("./logs/trade-text-warn.log")
	textErrorWriter := getWriter("./logs/trade-text-error.log")

	jsonDebugWriter := getWriter("./logs/trade-json-debug.log")
	jsonInfoWriter := getWriter("./logs/trade-json-info.log")
	jsonWarnWriter := getWriter("./logs/trade-json-warn.log")
	jsonErrorWriter := getWriter("./logs/trade-json-error.log")

	// 最后创建具体的 Logger
	textCore := zapcore.NewTee(
		zapcore.NewCore(textEncoder, zapcore.AddSync(textDebugWriter), debugLevle),
		zapcore.NewCore(textEncoder, zapcore.AddSync(textInfoWriter), infoLevel),
		zapcore.NewCore(textEncoder, zapcore.AddSync(textWarnWriter), warnLevel),
		zapcore.NewCore(textEncoder, zapcore.AddSync(textErrorWriter), errorLevel),
	)

	testLog := zap.New(textCore, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	textLogger = testLog.Sugar()

	jsonCore := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(jsonDebugWriter), debugLevle),
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(jsonInfoWriter), infoLevel),
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(jsonWarnWriter), warnLevel),
		zapcore.NewCore(jsonEncoder, zapcore.AddSync(jsonErrorWriter), errorLevel),
	)

	jsonLog := zap.New(jsonCore, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	jsonLogger = jsonLog.Sugar()
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的 Logger 实际生成的文件名 trade-text-debug-YY-mm-dd-HH-MM-SS.log
	// trade_debug 是指向最新日志的链接
	// 保存7天内的日志 每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		// 没有使用go风格反人类的format格式
		// strings.Replace(filename, ".log", "", -1)+"-%Y-%m-%d-%H-%M-%S.log",
		strings.Replace(filename, ".log", "", -1)+"-%Y-%m-%d.log",
		rotatelogs.WithLinkName(filename),
		// rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
	}
	return hook
}

// Debug ....
func Debug(args ...interface{}) {
	textLogger.Debug(args...)
}

// Debugf ...
func Debugf(template string, args ...interface{}) {
	textLogger.Debugf(template, args...)
}

// Info ...
func Info(args ...interface{}) {
	textLogger.Info(args...)
}

// Infof ...
func Infof(template string, args ...interface{}) {
	textLogger.Infof(template, args...)
}

// Warn ...
func Warn(args ...interface{}) {
	textLogger.Warn(args...)
}

// Warnf ...
func Warnf(template string, args ...interface{}) {
	textLogger.Warnf(template, args...)
}

// Error ...
func Error(args ...interface{}) {
	textLogger.Error(args...)
}

// Errorf ...
func Errorf(template string, args ...interface{}) {
	textLogger.Errorf(template, args...)
}

// DPanic ...
func DPanic(args ...interface{}) {
	textLogger.DPanic(args...)
}

// DPanicf ...
func DPanicf(template string, args ...interface{}) {
	textLogger.DPanicf(template, args...)
}

// Panic ...
func Panic(args ...interface{}) {
	textLogger.Panic(args...)
}

// Panicf ...
func Panicf(template string, args ...interface{}) {
	textLogger.Panicf(template, args...)
}

// Fatal ...
func Fatal(args ...interface{}) {
	textLogger.Fatal(args...)
}

// Fatalf ...
func Fatalf(template string, args ...interface{}) {
	textLogger.Fatalf(template, args...)
}

// JDebug ....
func JDebug(args ...interface{}) {
	jsonLogger.Debug(args...)
}

// JDebugf ...
func JDebugf(template string, args ...interface{}) {
	jsonLogger.Debugf(template, args...)
}

// JInfo ...
func JInfo(args ...interface{}) {
	jsonLogger.Info(args...)
}

// JInfof ...
func JInfof(template string, args ...interface{}) {
	jsonLogger.Infof(template, args...)
}

// JWarn ...
func JWarn(args ...interface{}) {
	jsonLogger.Warn(args...)
}

// JWarnf ...
func JWarnf(template string, args ...interface{}) {
	jsonLogger.Warnf(template, args...)
}

// JError ...
func JError(args ...interface{}) {
	jsonLogger.Error(args...)
}

// JErrorf ...
func JErrorf(template string, args ...interface{}) {
	jsonLogger.Errorf(template, args...)
}

// JDPanic ...
func JDPanic(args ...interface{}) {
	jsonLogger.DPanic(args...)
}

// JDPanicf ...
func JDPanicf(template string, args ...interface{}) {
	jsonLogger.DPanicf(template, args...)
}

// JPanic ...
func JPanic(args ...interface{}) {
	jsonLogger.Panic(args...)
}

// JPanicf ...
func JPanicf(template string, args ...interface{}) {
	jsonLogger.Panicf(template, args...)
}

// JFatal ...
func JFatal(args ...interface{}) {
	jsonLogger.Fatal(args...)
}

// JFatalf ...
func JFatalf(template string, args ...interface{}) {
	jsonLogger.Fatalf(template, args...)
}
