package xlog

import (
	"fmt"
	"strconv"
	"time"
)

// 打印trace日志
func (x *Xlog) Trace(tag string, a ...interface{}) {
	if tag == "" {
		tag = x.appName
	}
	if !x.logSwitch {
		return
	}
	if x.logLevel > Trace {
		return
	}

	f, l, m := parseAttribute()

	if !x.teeFile { // 输出到终端
		x.output("trace", tag, f, strconv.Itoa(l), m, fmt.Sprintln(a...))
		return
	}

	xlt := fmt.Sprintln(a...)
	xldata := XlogData{
		CreateTime: time.Now().Format(layout),
		LogLevel:   "trace",
		Tag:        tag,
		Name:       f,
		Line:       l,
		Methon:     m,
		Text:       xlt[:len(xlt)-1],
	}
	xlf := Instance().XlogFile(&x.fileConf, tag)
	xlbytes, _ := xldata.Bytes()
	xlf.Log2File(xlbytes)
}

// 打印info日志
func (x *Xlog) Info(tag string, a ...interface{}) {
	if tag == "" {
		tag = x.appName
	}
	if !x.logSwitch {
		return
	}
	if x.logLevel > Info {
		return
	}

	f, l, m := parseAttribute()

	if !x.teeFile { // 输出到终端
		x.output("info", tag, f, strconv.Itoa(l), m, fmt.Sprintln(a...))
		return
	}

	if x.teeFile {
		xlt := fmt.Sprintln(a...)
		xldata := XlogData{
			CreateTime: time.Now().Format(layout),
			LogLevel:   "info",
			Tag:        tag,
			Name:       f,
			Line:       l,
			Methon:     m,
			Text:       xlt[:len(xlt)-1],
		}
		xlf := Instance().XlogFile(&x.fileConf, tag)
		xlbytes, _ := xldata.Bytes()
		xlf.Log2File(xlbytes)
		return
	}

}

// 打印debug日志
func (x *Xlog) Debug(tag string, a ...interface{}) {
	if tag == "" {
		tag = x.appName
	}
	if !x.logSwitch {
		return
	}
	if x.logLevel > Debug {
		return
	}

	f, l, m := parseAttribute()

	if !x.teeFile { // 输出到终端
		x.output("debug", tag, f, strconv.Itoa(l), m, fmt.Sprintln(a...))
		return
	}

	if x.teeFile {
		xlt := fmt.Sprintln(a...)
		xldata := XlogData{
			CreateTime: time.Now().Format(layout),
			LogLevel:   "debug",
			Tag:        tag,
			Name:       f,
			Line:       l,
			Methon:     m,
			Text:       xlt[:len(xlt)-1],
		}
		xlf := Instance().XlogFile(&x.fileConf, tag)
		xlbytes, _ := xldata.Bytes()
		xlf.Log2File(xlbytes)
		return
	}
}

// 打印warn日志
func (x *Xlog) Warn(tag string, a ...interface{}) {
	if tag == "" {
		tag = x.appName
	}
	if !x.logSwitch {
		return
	}
	if x.logLevel > Warn {
		return
	}

	f, l, m := parseAttribute()

	if !x.teeFile { // 输出到终端
		x.output("warn", tag, f, strconv.Itoa(l), m, fmt.Sprintln(a...))
		return
	}

	if x.teeFile {
		xlt := fmt.Sprintln(a...)
		xldata := XlogData{
			CreateTime: time.Now().Format(layout),
			LogLevel:   "warn",
			Tag:        tag,
			Name:       f,
			Line:       l,
			Methon:     m,
			Text:       xlt[:len(xlt)-1],
		}
		xlf := Instance().XlogFile(&x.fileConf, tag)
		xlbytes, _ := xldata.Bytes()
		xlf.Log2File(xlbytes)
		return
	}
}

// 打印error日志
func (x *Xlog) Error(tag string, a ...interface{}) {
	if tag == "" {
		tag = x.appName
	}
	if !x.logSwitch {
		return
	}
	if x.logLevel > Error {
		return
	}

	f, l, m := parseAttribute()

	if !x.teeFile { // 输出到终端
		x.output("error", tag, f, strconv.Itoa(l), m, fmt.Sprintln(a...))
		return
	}

	if x.teeFile {
		xlt := fmt.Sprintln(a...)
		xldata := XlogData{
			CreateTime: time.Now().Format(layout),
			LogLevel:   "error",
			Tag:        tag,
			Name:       f,
			Line:       l,
			Methon:     m,
			Text:       xlt[:len(xlt)-1],
		}
		xlf := Instance().XlogFile(&x.fileConf, tag)
		xlbytes, _ := xldata.Bytes()
		xlf.Log2File(xlbytes)
		return
	}
}

// 打印fatal日志
func (x *Xlog) Fatal(tag string, a ...interface{}) {
	if tag == "" {
		tag = x.appName
	}
	if !x.logSwitch {
		return
	}

	f, l, m := parseAttribute()

	if !x.teeFile { // 输出到终端
		x.output("fatal", tag, f, strconv.Itoa(l), m, fmt.Sprintln(a...))
		return
	}

	if x.teeFile {
		xlt := fmt.Sprintln(a...)
		xldata := XlogData{
			CreateTime: time.Now().Format(layout),
			LogLevel:   "fatal",
			Tag:        tag,
			Name:       f,
			Line:       l,
			Methon:     m,
			Text:       xlt[:len(xlt)-1],
		}
		xlf := Instance().XlogFile(&x.fileConf, tag)
		xlbytes, _ := xldata.Bytes()
		xlf.Log2File(xlbytes)
		return
	}
}
