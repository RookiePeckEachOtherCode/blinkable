package zap

import (
	"blinkable/pkg/viper"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	cfg       = viper.Load("log")
	infoPath  = cfg.Viper.GetString("info_path")
	errorPath = cfg.Viper.GetString("error_path")
	devPath   = cfg.Viper.GetString("dev_path")
	mod       = cfg.Viper.GetString("mod")
)

// InitLogger 初始化zap
func Init() *zap.SugaredLogger {
	highPriority := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= zap.ErrorLevel
	})

	lowPriority := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l < zap.ErrorLevel && l >= zap.DebugLevel
	})

	encoder := getEncoder()

	infoSyncer := getInfoWriter()
	infoCore := zapcore.NewCore(encoder, infoSyncer, lowPriority)

	errorSyncer := getErrorWriter()
	errorCore := zapcore.NewCore(encoder, errorSyncer, highPriority)

	var core zapcore.Core

	if mod == "dev" {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

		core = zapcore.NewTee(
			infoCore,
			errorCore,
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewTee(infoCore, errorCore)
	}

	logger := zap.New(core, zap.AddCaller())
	sugarLogger := logger.Sugar()

	return sugarLogger
}

// 自定义日志输出格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 获取INFO的Writer
func getInfoWriter() zapcore.WriteSyncer {
	//lumberJack:日志切割归档
	lumberJackLogger := &lumberjack.Logger{
		Filename:   infoPath,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// 获取ERROR的Writer
func getErrorWriter() zapcore.WriteSyncer {
	//lumberJack:日志切割归档
	lumberJackLogger := &lumberjack.Logger{
		Filename:   errorPath,
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}
