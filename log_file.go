package xlog

// 示例:
// {"time": yyyyMMdd HH:mm:ss	, "level": "trace", "tag":"xxx", "methon":"package.methon", "id":"xxx", "dd":"xxxx"}

// 日志文件结构
type XlogFile struct {
	CreateTime string `json:"time"`
	LogLevel   string `json:"level"`
	Tag        string `json:"tag"`

	Name   string `json:"file"`
	Line   uint32 `json:"line"`
	Methon string `json:"methon"`

	Text interface{} `json:"text"`
}

// 追加日志
func (x *XlogFile) Append() error {

}
