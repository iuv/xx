package base

import "strings"

// 获取当前运行系统
func Sys() string {
	ret := Exec("uname -a")
	if strings.Index(ret, "Darwin") >= 0 {
		return "mac"
	} else if strings.Index(ret, "Linux") >= 0 {
		return "linux"
	} else {
		return "win"
	}
}

// 获取cpu架构
func Cpu() string {
	ret := Exec("uname -a")
	ret = strings.ToLower(ret)
	if strings.Index(ret, "aarch64") >= 0 || strings.Index(ret, "arm64") >= 0 {
		return "arm64"
	} else if strings.Index(ret, "aarch32") >= 0 || strings.Index(ret, "arm32") >= 0 {
		return "arm"
	} else if strings.Index(ret, "amd64") >= 0 {
		return "x86"
	} else {
		return ""
	}
}
