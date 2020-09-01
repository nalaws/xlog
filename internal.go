package xlog

import (
	"runtime"
	"strings"
)

func parseAttribute() (string, int, string) {
	var nm, mth string
	var ln int
	if pc, file, line, ok := runtime.Caller(2); ok {
		f := runtime.FuncForPC(pc)
		pos := strings.LastIndex(file, "/")
		nm = file[pos+1:]
		ln = line
		mth = f.Name()
	}

	return nm, ln, mth
}
