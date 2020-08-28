package xlog

import (
	"sync"
)

type Xlogs map[string]Xlog

var (
	xlogs Xlogs
	mu    sync.Mutex
)

func init() {
	xlogs = make(map[string]Xlog)
}

func (xs Xlogs) Add(tag string, log *Xlog) {
	mu.Lock()
	defer mu.Unlock()
	xlogs[tag] = *log
}
