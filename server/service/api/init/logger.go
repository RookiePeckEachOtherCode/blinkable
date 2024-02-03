package init

import (
	"blinkable/server/common/consts"
	"os"
	"path"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/fatih/color"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

type CustomTextFormatter struct {
	logrus.TextFormatter
	ForceColors   bool
	ColorInfo     *color.Color
	ColorWarning  *color.Color
	ColorError    *color.Color
	ColorCritical *color.Color
}

func InitLogger() {
	logFilePath := consts.HlogFIlePath
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}

	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}

	logger := hertzlogrus.NewLogger()
	logger.Logger().Formatter = &CustomTextFormatter{
		ForceColors:   true,
		ColorInfo:     color.New(color.FgBlue),
		ColorWarning:  color.New(color.FgYellow),
		ColorError:    color.New(color.FgRed),
		ColorCritical: color.New(color.BgRed, color.FgWhite),
	}

	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}

	logger.SetOutput(lumberjackLogger)
	logger.SetLevel(hlog.LevelDebug)

	logger.SetOutput(os.Stdout)
	hlog.SetLogger(logger)
}
