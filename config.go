package xlog

// 日志配置
type XlogConfig struct {
	OnOff    bool    // 日志开关
	LogLevel Level   // 日志级别
	OutFile  bool    // 是否输出到文件
	XlogFile LogFile // 日志文件
}

// 日志文件
type LogFile struct {
	Dir           string // 文件路径
	Prefix        string // 文件前缀
	Suffix        string // 文件后缀
	Split         int    // 文件是否拆分 0: 不拆分, 1: 按时间拆分, 2: 按条数拆分
	SplitInterval int64  // 文件拆分间隔.单位: s
	SplitMax      int64  // 每个文件存多少条. 0: 不限制
}
