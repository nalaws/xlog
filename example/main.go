package main

import (
	"fmt"

	"github.com/nalaws/xlog"
)

func main() {
	fmt.Println("begin.")
	defer func() {
		fmt.Println("end.")
	}()

	log := xlog.NewXlog()
	log.SetLogLevel(xlog.Info)
	log.Trace("a", "dd")
	log.Info("b", "info")
}
