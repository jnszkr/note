package formatter

import (
	"fmt"
	"runtime"
)

var (
	reset = "\033[0m"
	red   = "\033[31m"
)

func init() {
	if runtime.GOOS == "windows" {
		reset = ""
		red = ""
	}
}

func Red(s string) string {
	return fmt.Sprintf("%s%s%s", red, s, reset)
}
