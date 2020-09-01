package xlog

import (
	"context"
	"sync"
	"time"
)

const (
	layout = "2006/01/02 15:04:05"

	dayLayout  = "2006-01-02"
	hourLayout = "2006-01-02-15"
	timeLayout = "2006-01-02-15-04-05"
)

var (
	xlgCtx, xlgCancelFunc = context.WithCancel(context.TODO())
	schudelCheckTime      = 5 * time.Second

	xlgTimerStart = false
	xlgTimerMu    sync.Mutex
)

func XlogFileTimerStart() {
	if !xlgTimerStart {
		xlgTimerMu.Lock()
		if !xlgTimerStart {
			xlgTimerStart = true
			go managerLogFile()
		}
		xlgTimerMu.Unlock()
	}
}

func XlogFileTimerStop() {
	if xlgTimerStart {
		xlgTimerMu.Lock()
		if xlgTimerStart {
			xlgTimerStart = false
			xlgCancelFunc()
		}
		xlgTimerMu.Unlock()
	}
}

func managerLogFile() {
	t := time.NewTimer(schudelCheckTime)
	defer t.Stop()
	c, _ := context.WithCancel(xlgCtx)
loop:
	for {
		select {
		case <-c.Done():
			break loop
		case <-t.C:
			t.Reset(schudelCheckTime)
			closePreLogFile()
		}
	}
}

func closePreLogFile() {
	for _, xlg := range Instance().xlogs {
		if xlg != nil {
			if xlg.f != nil && xlg.preFile != nil {
				xlg.preFile.closeLogFile()
			}
		}
	}
}
