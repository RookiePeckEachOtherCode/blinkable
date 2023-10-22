package errno

import (
	"blinkable/pkg/zap"
)

var (
	zlog = zap.Init()
)

func HandleErrWithPanic(where string, str string, err error) {
	if err != nil {
		zlog.Panicf("[%s] %s: %s", where, str, err)
		return
	}
	zlog.Panicf("[%s] %s", where, str)
}

func HandleErrWithError(where string, str string, err error) {
	if err != nil {
		zlog.Error("[%s] %s: %s", where, str, err)
		return
	}
	zlog.Errorf("[%s] %s", where, str)
}
func HandleErrWithFatal(where string, str string, err error) {
	if err != nil {
		zlog.Fatalf("[%s] %s: %s", where, str, err)
		return
	}
	zlog.Infof("[%s] %s", where, str)
}
func HandleERrWithDebug(where string, str string, err error) {
	if err != nil {
		zlog.Debugf("[%s] %s %s", where, str, err)
		return
	}
	zlog.Debugf("[%s] %s", where, str)
}
