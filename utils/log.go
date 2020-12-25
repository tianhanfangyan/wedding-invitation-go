package utils

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

func InitLogger() error {
	path, _ := os.Getwd()
	fileName := path + GetConf().String("log::filename")
	_, err := os.Stat(fileName)
	if err != nil {
		err := os.MkdirAll(filepath.Dir(fileName), os.ModePerm)
		if err != nil {
			panic("无法创建日志目录")
			return err
		}
	}

	// 自定义日志
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(getLogLevel()),
		Development: true,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",    // 输出时间的key名
			LevelKey:       "level", // 输出日志级别的key名
			NameKey:        "logger",
			CallerKey:      "lineno", // 输出行号的key名
			MessageKey:     "msg",    // 输出信息的key名
			StacktraceKey:  "trace",
			LineEnding:     zapcore.DefaultLineEnding,      // 每行的分隔符
			EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 将日志级别字符串转化为小写
			EncodeTime:     zapcore.ISO8601TimeEncoder,     // 输出的时间格式
			EncodeDuration: zapcore.SecondsDurationEncoder, // 执行消耗的时间转化成浮点型的秒
			EncodeCaller:   zapcore.ShortCallerEncoder,     // 以包/文件:行号 格式化调用堆栈
		},
		OutputPaths:      []string{fileName, "stdout"},
		ErrorOutputPaths: []string{fileName, "stdout"},
	}

	// 构造日志
	globalLogger, err = cfg.Build()
	if err != nil {
		return err
	}

	return nil
}

// 得到日志级别
func getLogLevel() zapcore.Level {
	env := GetEnv()
	switch env {
	case ENV_DEV:
		return zap.DebugLevel
	case ENV_TEST:
		return zap.InfoLevel
	case ENV_PROD:
		return zap.InfoLevel
	default:
		return zap.DebugLevel
	}
}

// 同步缓冲日志
func Sync() error {
	return globalLogger.Sync()
}

func Sugar() *zap.SugaredLogger {
	return GetLog().Sugar()
}

func GetLog() *zap.Logger {
	return globalLogger
}
