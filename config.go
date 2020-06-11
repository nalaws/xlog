package xlog

// 日志配置
type XConfig struct {
	LogLevel Level    // 日志级别
	OutFile  bool     // 是否输出到文件
	XlogFile *LogFile // 日志文件
}

// 日志文件
type LogFile struct {
	Path string
	// 多少条写一次日志
	// 多久写一次日志
	// 按小时生成日志
	// 按日生成日志
}
