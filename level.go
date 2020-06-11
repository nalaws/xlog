package xlog

type Level uint8

const (
	Trace Level = 1 << iota
	Info
	Debug
	Warn
	Error
	Fatal
)

func (l *Level) String() string {
	switch *l {
	case Trace:
		return "trace"
	case Info:
		return "info"
	case Debug:
		return "debug"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Fatal:
		return "fatal"
	}
	return ""
}
