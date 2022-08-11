/*
 * @Date: 2022-02-14 14:44:34
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-08-01 17:02:39
 * @FilePath: \server\system\core\log\log.go
 */

package log

import (
	"os"
	"path"
	"time"

	"chunDoc/system/core/config"
	"chunDoc/system/util/sys"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func InitLog() *zap.Logger {

	if logger != nil {
		return logger
	}

	basedir := sys.ExecutePath() + "\\" //根目录

	hook := lumberjack.Logger{
		Filename:   path.Join(basedir+config.Instance().Zaplog.Director, time.Now().Format("2006-01-02")+".log"), // 日志文件路径
		MaxSize:    10,                                                                                           // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 50,                                                                                           // 日志文件最多保存多少个备份
		MaxAge:     10,                                                                                           // 文件最多保存多少天
		Compress:   true,                                                                                         // 是否压缩
		LocalTime:  true,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()

	logLevel := zap.DebugLevel
	switch config.Instance().Zaplog.Level {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	case "panic":
		logLevel = zap.PanicLevel
	case "fatal":
		logLevel = zap.FatalLevel
	default:
		logLevel = zap.InfoLevel
	}
	atomicLevel.SetLevel(logLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), // 编码器配置
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)), // 打印到文件
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serverName", "go2admin"))
	// 构造日志
	logger = zap.New(core, caller, development, filed)
	return logger
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

const (
	Debug = "debug"
	Info  = "info"
	Warn  = "warn"
	Sql   = "sql"
	Error = "error"
	Panic = "panic"
	Fatal = "fatal"
)

// 写日志
func Write(level string, msg string) {
	switch level {
	case Debug, Sql:
		logger.Debug(msg)
	case Info:
		logger.Info(msg)
	case Warn:
		logger.Warn(msg)
	case Error:
		logger.Error(msg)
	case Panic:
		logger.Panic(msg)
	case Fatal:
		logger.Fatal(msg)
	default:
		logger.Info(msg)
	}
}
