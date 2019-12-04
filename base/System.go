package base

import "strings"

// 获取当前运行系统
func Sys() string {
	ret := Exec("uname -a")
	if strings.Index(ret, "Darwin")>=0 {
		return "mac"
	} else if strings.Index(ret, "Linux")>=0 {
		return "linux"
	} else {
		return "win"
	}
}
