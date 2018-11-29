package logs

import (
	"fmt"
	"os"
)

func Fatalf(format string, args ...interface{}) {
	fmt.Printf(format, args)
	os.Exit(1)
}