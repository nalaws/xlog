package xlog

import (
	"runtime"
	"strings"
)

// skip: 跳过堆栈数。2: 直接使用, 每多封装一次值加1
func parseAttribute(skip int) (string, int, string) {
	var nm, mth string
	var ln int
	if pc, file, line, ok := runtime.Caller(skip); ok {
		f := runtime.FuncForPC(pc)
		pos := strings.LastIndex(file, "/")
		nm = file[pos+1:]
		ln = line
		mth = f.Name()
		pos = strings.LastIndex(mth, "/")
		if pos > 0 {
			mth = mth[pos+1:]
		}
	}

	return nm, ln, mth
}
