package xlog

import (
	"fmt"
	"time"
)

// 示例:
// {"time": yyyyMMdd HH:mm:ss	, "level": "trace", "tag":"xxx", "methon":"package.methon", "id":"xxx", "dd":"xxxx"}

type Xlog struct {
	logSwitch bool  // 日志总开关  true: 输出日志, false: 不输出日志
	logLevel  Level // 定义日志级别
	teeFile   bool  // 是否输出到文件
	Second    int   // 0: 立即输出到文件 60*60: 按小时输出到文件 60*60*24 按天输出到文件
}

type XlogInfo struct {
	CreateTime time.Time `json:"time"`
	LogLevel   string    `json:"level"`
	Tag        string    `json:"tag"`

	Name   string `json:"file"`
	Line   uint32 `json:"line"`
	Methon string `json:"methon"`

	Text interface{} `json:"text"`
}

func NewXlog() *Xlog {
	return &Xlog{
		logSwitch: true,
		logLevel:  Trace,
		teeFile:   false,
	}
}

// 设置日志开关
func (x *Xlog) SetLogSwitch(b bool) {
	x.logSwitch = b
}

// 设置日志级别
func (x *Xlog) SetLogLevel(l Level) {
	x.logLevel = l
}

// 设置日志存储到文件
func (x *Xlog) SetTeeFile(b bool) {
	x.teeFile = b
}

// 打印trace日志
func (x *Xlog) Trace(tag string, a ...interface{}) {
	if !x.logSwitch {
		return
	}
	if x.logLevel > Trace {
		return
	}

	f, l, m := parseAttribute()
	fmt.Printf("%s %s %s [%s:%d] (%s): ", time.Now().Format(layout), "trace", tag, f, l, m)
	fmt.Println(a...)
}

// 打印info日志
func (x *Xlog) Info(tag string, a ...interface{}) {
	if !x.logSwitch {
		return
	}
	if x.logLevel > Info {
		return
	}

	f, l, m := parseAttribute()
	fmt.Printf("%s %s %s [%s:%d] (%s): ", time.Now().Format(layout), "info", tag, f, l, m)
	fmt.Println(a...)
}

// 打印debug日志
func (x *Xlog) Debug(tag string, a ...interface{}) {
	if !x.logSwitch {
		return
	}
	if x.logLevel > Debug {
		return
	}

	f, l, m := parseAttribute()
	fmt.Printf("%s %s %s [%s:%d] (%s): ", time.Now().Format(layout), "debug", tag, f, l, m)
	fmt.Println(a...)
}

// 打印warn日志
func (x *Xlog) Warn(tag string, a ...interface{}) {
	if !x.logSwitch {
		return
	}
	if x.logLevel > Warn {
		return
	}

	f, l, m := parseAttribute()
	fmt.Printf("%s %s %s [%s:%d] (%s): ", time.Now().Format(layout), "warn", tag, f, l, m)
	fmt.Println(a...)
}

// 打印error日志
func (x *Xlog) Error(tag string, a ...interface{}) {
	if !x.logSwitch {
		return
	}
	if x.logLevel > Error {
		return
	}

	f, l, m := parseAttribute()
	fmt.Printf("%s %s %s [%s:%d] (%s): ", time.Now().Format(layout), "error", tag, f, l, m)
	fmt.Println(a...)
}

// 打印fatal日志
func (x *Xlog) Fatal(tag string, a ...interface{}) {
	if !x.logSwitch {
		return
	}

	f, l, m := parseAttribute()
	fmt.Printf("%s %s %s [%s:%d] (%s): ", time.Now().Format(layout), "fatal", tag, f, l, m)
	fmt.Println(a...)
}
