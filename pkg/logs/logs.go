package logs

import (
	"fmt"
	"os"
)

// 打印失败信息，并退出应用
func Fatalf(format string, args ...interface{}) {
	fmt.Printf(format, args)
	os.Exit(1)
}
