package global

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger LoggerInterface

type LoggerInterface interface {
	Debug(a ...interface{})
	Debugf(format string, a ...interface{})
	Info(a ...interface{})
	Infof(format string, a ...interface{})
	Warn(a ...interface{})
	Warnf(format string, a ...interface{})
	Error(a ...interface{})
	Errorf(format string, a ...interface{})
	Fatal(a ...interface{})
	Fatalf(format string, a ...interface{})
	Sync()
}

func initLogger(level string, encoder string) error {
	Logger = NewGinLogger(level, encoder)
	return nil
}

type GinLogger struct {
	logger *zap.SugaredLogger
}

func NewGinLogger(level string, encoder string) *GinLogger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 初始化配置文件的Level
	var zapLevel zapcore.Level
	switch level {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "warn":
		zapLevel = zap.WarnLevel
	case "error":
		zapLevel = zap.ErrorLevel
	case "dpanic":
		zapLevel = zap.DPanicLevel
	case "panic":
		zapLevel = zap.PanicLevel
	case "fatal":
		zapLevel = zap.FatalLevel
	default:
		zapLevel = zap.InfoLevel
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapLevel)

	// 设置option
	options := []zap.Option{zap.AddCaller(), zap.AddCallerSkip(1)}

	var zapEncoder zapcore.Encoder
	if encoder == "json" {
		zapEncoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		zapEncoder = zapcore.NewConsoleEncoder(encoderConfig)
		options = append(options, zap.AddStacktrace(zap.ErrorLevel))
	}

	// 设置日志文件切割
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "/var/log/goweb/gin.log",
		MaxSize:    100,
		MaxBackups: 7,
		MaxAge:     30,
		Compress:   true,
	}

	core := zapcore.NewCore(
		zapEncoder,
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger)), // 打印到控制台和文件
		atomicLevel,
	)

	zapLogger := zap.New(core, options...)
	return &GinLogger{
		logger: zapLogger.Sugar(),
	}
}

func (l *GinLogger) Debug(a ...interface{}) {
	l.logger.Debug(a)
}

func (l *GinLogger) Debugf(format string, a ...interface{}) {
	l.logger.Debug(fmt.Sprintf(format, a))
}

func (l *GinLogger) Info(a ...interface{}) {
	l.logger.Info(a)
}

func (l *GinLogger) Infof(format string, a ...interface{}) {
	l.logger.Info(fmt.Sprintf(format, a))
}

func (l *GinLogger) Warn(a ...interface{}) {
	l.logger.Warn(a)
}

func (l *GinLogger) Warnf(format string, a ...interface{}) {
	l.logger.Warn(fmt.Sprintf(format, a))
}

func (l *GinLogger) Error(a ...interface{}) {
	l.logger.Error(a)
}

func (l *GinLogger) Errorf(format string, a ...interface{}) {
	l.logger.Error(fmt.Sprintf(format, a))
}

func (l *GinLogger) Fatal(a ...interface{}) {
	l.logger.Fatal(a)
}

func (l *GinLogger) Fatalf(format string, a ...interface{}) {
	l.logger.Fatal(fmt.Sprintf(format, a))
}

func (l *GinLogger) Sync() {
	_ = l.logger.Sync()
}
