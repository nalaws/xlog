// +build linux

package xlog

import (
	"fmt"
	"os"
	"time"
)

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
//
// 代码 意义
// -------------------------
//  0  终端默认设置
//  1  高亮显示
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  不可见
// fmt.Printf(" %c[%d;%dm%s###%c[0m ", 0x1B, 0, 32, "", 0x1B)

var (
	colorsPrefix []byte
	colorsSuffix []byte
)

func init() {
	colorsPrefix = []byte(fmt.Sprintf("%c[0;", 0x1B))
	colorsSuffix = []byte(fmt.Sprintf("%c[0m", 0x1B))
}

func (x *Xlog) output(level, tag, fname, fline, method, text string) {
	x.lock.Lock()
	defer x.lock.Unlock()
	buf := []byte{}
	buf = append(buf, colorsPrefix...)
	buf = append(buf, x.levelColors(level)...)
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
	buf = append(buf, colorsSuffix...)
	os.Stdout.Write(buf)
}

func (x *Xlog) levelColors(level string) []byte {
	switch level {
	case "trace":
		return []byte("37m")
	case "info":
		return []byte("32m")
	case "debug":
		return []byte("36m")
	case "warn":
		return []byte("33m")
	case "error":
		return []byte("31m")
	case "fatal":
		return []byte("35m")
	default:
		return []byte("37m")
	}
}
