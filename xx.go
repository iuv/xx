package main

import (
	"fmt"
	"github.com/iuv/xx/base"
	"github.com/iuv/xx/docker"
	"github.com/iuv/xx/k8s"
	"github.com/iuv/xx/shell"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Get help with \"xx help\" command ")
		return
	}
	s := args[1]
	a1, a2, a3, a4, _ := setArg(args)
	switch s[0:1] {
	case "d":
		docker.Docker(s[1:], a1, a2, a3)
	case "k":
		k8s.K8s(s[1:], a1, a2, a3, a4)
	case "s":
		shell.Shell(s[1:], a1, a2, a3, a4)
	default:
		switch s {
		case "ip":
			shell.Ip(a1)
		case "v", "version", "-v":
			base.V()
		case "h", "-h", "help":
			base.Help(a1)
		case "zh":
			base.HelpZh(a1)
		case "ps":
			shell.Ps(a1)
		case "install":
			install()
		case "update":
			update()
		default:
			fmt.Println("command not found: " + s)
		}
	}
}

// 安装
func install() {
	base.Execp("sudo mv xx /usr/local/bin")
	base.Execp("chmod 775 /usr/local/bin/xx")
}

// 更新
func update() {
	ret := base.Exec("curl https://jianpage.com/xx/")
	if strings.Index(ret, base.Version) > 0 {
		fmt.Println("The " + base.Version + " is already the latest version")
		return
	}
	sys := base.Sys()
	if sys == "mac" {
		base.Run("wget https://raw.githubusercontent.com/iuv/xx/master/build/mac/xx")
	} else if sys == "linux" {
		base.Run("wget https://raw.githubusercontent.com/iuv/xx/master/build/linux/xx")
	}
	base.Execp("chmod 775 xx")
	base.Execp("./xx install")
}
func setArg(args []string) (string, string, string, string, string) {
	a1, a2, a3, a4, a5 := "", "", "", "", ""
	if len(args) >= 3 {
		a1 = args[2]
	}
	if len(args) >= 4 {
		a2 = args[3]
	}
	if len(args) >= 5 {
		a3 = args[4]
	}
	if len(args) >= 6 {
		a4 = args[5]
	}
	if len(args) >= 7 {
		a5 = args[6]
	}
	return a1, a2, a3, a4, a5
}
