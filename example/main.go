package main

import (
	"fmt"
	"sync"

	"github.com/nalaws/xlog"
)

func main() {
	fmt.Println("begin.")
	defer func() {
		fmt.Println("end.")
	}()

	// test1()
	test2()
}

func test1() {
	log := xlog.NewXlog()
	log.SetLogLevel(xlog.Info)
	log.Trace("a", "dd")
	log.Info("b", "info")
}

func test2() {
	log := xlog.NewXlog()
	wg := sync.WaitGroup{}
	wg.Add(2000)
	go func() {
		log.SetLogLevel(xlog.Info)
		for i := 0; i < 1000; i++ {
			log.Info("#info#", i)
			wg.Done()
		}
	}()

	go func() {
		log.SetLogLevel(xlog.Error)
		for i := 0; i < 1000; i++ {
			log.Error("#error#", i)
			wg.Done()
		}
	}()

	wg.Wait()
}
