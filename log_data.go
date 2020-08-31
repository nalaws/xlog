package xlog

import (
	"encoding/json"
	"fmt"
)

// {"time": yyyyMMdd HH:mm:ss, "level": "trace", "tag":"xxx", "file":"main.go","line":65,"methon":"package.methon", "text":"xxxx"}
type XlogData struct {
	CreateTime string `json:"time"`
	LogLevel   string `json:"level"`
	Tag        string `json:"tag"`

	Name   string `json:"file"`
	Line   int    `json:"line"`
	Methon string `json:"methon"`

	Text interface{} `json:"text"`
}

func (x *XlogData) Bytes() ([]byte, error) {
	bs, err := json.Marshal(*x)
	if err != nil {
		x.Text = fmt.Sprintf("xlog json.Marshal(XlogFile) error: %s", err.Error())
		es, _ := json.Marshal(*x)
		es = append(es, '\n')
		return es, err
	}
	return bs, nil
}
