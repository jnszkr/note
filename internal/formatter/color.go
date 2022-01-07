package formatter

import (
	"fmt"
	"regexp"
	"runtime"
)

type Color string

var (
	Reset  Color = "\033[0m"
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Blue   Color = "\033[34m"
	Purple Color = "\033[35m"
	Cyan   Color = "\033[36m"
	Gray   Color = "\033[37m"
	White  Color = "\033[97m"
)

var winos = false

func init() {
	winos = runtime.GOOS == "windows"
	if winos {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

// Highlight would highlight the substring from s with the given color.
func Highlight(s, substr string, c Color) string {
	if winos {
		return s
	}
	re := regexp.MustCompile(fmt.Sprintf("(?i)(%s)", substr))
	return re.ReplaceAllString(s, fmt.Sprintf("%s$1%s", c, Reset))
}
