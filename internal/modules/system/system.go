package system

import (
	"runtime"
)

func GetSeparator() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}

	return "\n"
}
