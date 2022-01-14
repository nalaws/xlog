package main

import (
	"fmt"
	"sync"

	"github.com/nalaws/xlog"
)

var (
	tag = "tag"
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
	log.Trace(tag, "bb")
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
	defer log.Close()
	log.SetLogLevel(xlog.Trace)
	err := log.SetTeeFile(true)
	/*conf := xlog.LogFile{
		Split:         1,
		SplitInterval: 1 * 60 * 60,
	}
	log.SetXlogFileConfig(&conf)*/
	if err != nil {
		fmt.Println(err)
	}
	_, e := xlog.IsFileExist("e:/golang/src/github.com/nalaws/xlog/example/1.log")
	log.Trace("tag", e)
	log.Trace("error", e)
}

func test4() {
	log := xlog.NewXlog()
	defer log.Close()
	err := log.SetTeeFile(true)
	if err != nil {
		fmt.Println(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2000)
	go func() {
		for i := 0; i < 1000; i++ {
			log.Info("info", i)
			//log.Info("", i)
			wg.Done()
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			log.Error("error", i)
			//log.Error("", i)
			wg.Done()
		}
	}()

	wg.Wait()
}
