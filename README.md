## 描述
golang日志输出

终端输出：
2020/06/11 07:15:16 info \<tag\> [main.go:18] (main.main): \<log info\>

示例:
{"time":"yyyy/MM/dd HH:mm:ss", "level":"info", "tag":"###", "methon":"package.methon", "text":"xxxx"}

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
