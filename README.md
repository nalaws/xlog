## 描述
重新封装golang日志输出

终端输出：
2020/06/11 07:15:16 info ### [main.go:18] (main.main): xxxxx

示例:
{"time":yyyy/MM/dd HH:mm:ss, "level":"info", "tag":"###", "methon":"package.methon", "text":"xxxx"}

## 使用
```golang
import (
	"github.com/nalaws/xlog"
)

func test() {
    log := xlog.NewXlog()
    log.SetLogLevel(xlog.Info) // 默认trace级别
    log.Info("tag", "info")
}
```
