package errno

import (
	"blinkable/pkg/zap"
)

var (
	zlog = zap.Init()
)

func HandleErrWithPanic(str string, err error) {
	if err != nil {
		zlog.Panicf("%s: %s", str, err.Error())
	}
}
func HandleErrWithNormal(str string, err error) {
	if err != nil {
		zlog.Errorf("%s: %s", str, err.Error())
	}
}
func HandleErrWithFatal(str string, err error) {
	if err != nil {
		zlog.Fatalf("%s: %s", str, err.Error())
	}
}

func Error(str string) {
	zlog.Error(str)
}
