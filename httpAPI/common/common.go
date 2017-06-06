package common

import "runtime"

func Fn() string {
	if pc, _, _, ok := runtime.Caller(1); ok {
		fn := runtime.FuncForPC(pc)
		return fn.Name()
	}
	panic("couldn't determine function name")
}
