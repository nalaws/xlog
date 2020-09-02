package xlog

import (
	"fmt"
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

		fmt.Println("2", mth)
	}

	if pc, file, line, ok := runtime.Caller(1); ok {
		f := runtime.FuncForPC(pc)
		pos := strings.LastIndex(file, "/")
		nm = file[pos+1:]
		ln = line
		mth = f.Name()

		fmt.Println("1", mth)
	}

	if pc, file, line, ok := runtime.Caller(3); ok {
		f := runtime.FuncForPC(pc)
		pos := strings.LastIndex(file, "/")
		nm = file[pos+1:]
		ln = line
		mth = f.Name()

		fmt.Println("3", mth)
	}

	return nm, ln, mth
}
