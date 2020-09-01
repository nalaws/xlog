package xlog

import (
	"context"
	"sync"
	"time"
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
