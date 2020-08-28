// +build windows

package xlog

import (
	"os"
	"syscall"
	"time"
)

var (
	kernel32    *syscall.LazyDLL  = syscall.NewLazyDLL(`kernel32.dll`)
	proc        *syscall.LazyProc = kernel32.NewProc(`SetConsoleTextAttribute`)
	CloseHandle *syscall.LazyProc = kernel32.NewProc(`CloseHandle`)
)

type XCmdColor int

const (
	black       XCmdColor = iota // 黑色
	blue                         // 蓝色
	green                        // 绿色
	cyan                         // 青色
	red                          // 红色
	purple                       // 紫色
	yellow                       // 黄色
	lightGray                    // 淡灰色（系统默认值）
	gray                         // 灰色
	lightBlue                    // 亮蓝色
	lightGreen                   // 亮绿色
	lightCyan                    // 亮青色
	lightRed                     // 亮红色
	lightPurple                  // 亮紫色
	lightYellow                  // 亮黄色
	white                        // 白色
)

func (x *Xlog) output(level, tag, fname, fline, method, text string) {
	x.lock.Lock()
	defer x.lock.Unlock()
	// windows 如果做服务器, 这儿可以优化, 可以参考系统源码 Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(x.levelColors(level)))
	defer func() {
		backHandle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(lightGray))
		CloseHandle.Call(backHandle)
		CloseHandle.Call(handle)
	}()

	buf := []byte{}
	buf = append(buf, []byte(time.Now().Format(layout))...)
	buf = append(buf, ' ')
	buf = append(buf, []byte(level)...)
	buf = append(buf, ' ')
	buf = append(buf, []byte(tag)...)
	buf = append(buf, []byte{' ', '['}...)
	buf = append(buf, []byte(fname)...)
	buf = append(buf, ' ')
	buf = append(buf, []byte(fline)...)
	buf = append(buf, []byte{']', ' ', '('}...)
	buf = append(buf, []byte(method)...)
	buf = append(buf, []byte{')', ':', ' '}...)
	buf = append(buf, []byte(text)...)
	os.Stdout.Write(buf)
}

func (x *Xlog) levelColors(level string) XCmdColor {
	switch level {
	case "trace":
		return gray
	case "info":
		return green
	case "debug":
		return blue
	case "warn":
		return yellow
	case "error":
		return red
	case "fatal":
		return purple
	default:
		return black
	}
}
