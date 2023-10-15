package errs

import "log"

func HandleErrWithPanic(str string, err error) {
	if err != nil {
		log.Panicf("%s: %s", str, err.Error())
	}
}
func HandleErrWithNormal(str string, err error) {
	if err != nil {
		log.Printf("%s: %s", str, err.Error())
	}
}
