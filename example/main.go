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

	test1()
	//test2()
	//test3()
	//test4()
}

func test1() {
	log := xlog.NewXlog()
	log.SetLogLevel(xlog.Trace)
	log.Trace("t", "bb")
	log.Trace("t", "cc", "dd")
	log.Info("i", "info")
	log.Debug("d", "debug")
	log.Warn("w", "warn")
	log.Error("e", "error")
	log.Fatal("f", "fatal")
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

func test3() {
	log := xlog.NewXlog()
	log.SetLogLevel(xlog.Trace)
	err := log.SetTeeFile(true)
	if err != nil {
		fmt.Println(err)
	}
	_, e := xlog.IsFileExist("d:/2.log")
	log.Trace("tag", e)
	log.Trace("error", e)
}

func test4() {
	log := xlog.NewXlog()
	err := log.SetTeeFile(true)
	if err != nil {
		fmt.Println(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2000)
	go func() {
		for i := 0; i < 1000; i++ {
			log.Info("#info#", i)
			wg.Done()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			log.Error("#error#", i)
			wg.Done()
		}
	}()

	wg.Wait()
}
