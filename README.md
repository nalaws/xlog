## 描述
golang日志输出

终端输出：
`2020/06/11 07:15:16 info <tag> [main.go:18] (main.main): <log info>`

文件输出:
`{"time":"yyyy/MM/dd HH:mm:ss", "level":"info", "tag":"###", "methon":"package.methon", "text":"xxxx"}`

## 使用
```golang
import (
	"github.com/nalaws/xlog"
)

var (
	tag = "tag"
)

func test() {
    log := xlog.NewXlog()
    defer log.Close()
    log.SetLogLevel(xlog.Info) // 默认trace级别
    log.Info(tag, "info")
}
```

**注意**：如果对包进行二次封装注意设置堆栈层数

```golang
func test() {
    log := xlog.NewXlog()
    defer log.Close()
    log.SetLogLevel(xlog.Info) // 默认trace级别
    log.SetSkip(1) // 如果进行二次封装,根据日志输出注意调整
    log.Info(tag, "info")
}
```
