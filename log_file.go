package xlog

import (
	"encoding/json"
	"errors"
	"fmt"
)

// 示例:
// {"time": yyyyMMdd HH:mm:ss	, "level": "trace", "tag":"xxx", "methon":"package.methon", "id":"xxx", "dd":"xxxx"}

// 日志文件结构
type XlogFile struct {
	CreateTime string `json:"time"`
	LogLevel   string `json:"level"`
	Tag        string `json:"tag"`

	Name   string `json:"file"`
	Line   int    `json:"line"`
	Methon string `json:"methon"`

	Text interface{} `json:"text"`
}

// 追加日志
func (x *XlogFile) Append() error {
	if f == nil {
		return errors.New("日志没有初始化")
	}

	bs, err := json.Marshal(*x)
	if err != nil {
		x.Text = fmt.Sprintf("xlog json.Marshal(XlogFile) error: %s", err.Error())
		es, _ := json.Marshal(*x)
		es = append(es, '\n')
		f.Write(es)
		return err
	}
	bs = append(bs, '\n')
	f.Write(bs)
	return nil
}
