package xlog

import (
	"os"
	"time"
)

func (x *Xlog) output(level, tag, fname, fline, method, text string) {
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
	buf = append(buf, '\n')
	os.Stdout.Write(buf)
}
