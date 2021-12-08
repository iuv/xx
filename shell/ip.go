package shell

import (
	"fmt"
	"github.com/iuv/xx/base"
	"strings"
)

func Ip(port string) {
	sys := base.Sys()
	var ret int
	if sys == "mac" {
		ret = exe("en0", "inet ", port)
	} else if sys == "linux" {
		ret = exe("eth0", "inet ", port)
		if ret < 0 {
			ret = exe("ens33", "inet ", port)
		}
	}
	if ret < 0 {
		def(port)
	}
}

func def(port string) {
	ret := base.Exec("ifconfig")
	rets := strings.Split(ret, "\n")
	for _, ip := range rets {
		inet := strings.Index(ip, "inet ")
		if inet < 0 {
			continue
		}
		inet = inet + 5
		end := inet + strings.Index(ip[inet:], " ")
		ip = ip[inet:end]
		addr := strings.Index(ip, "addr:")
		if addr > 0 {
			ip = ip[addr+5:]
		}
		fmt.Println("Local IP: " + ip)
		if port != "" {
			fmt.Println("HTTP Server: http://" + ip + ":" + port)
		} else {
			fmt.Println("HTTP Server: http://" + ip)
			fmt.Println("HTTP Server: http://" + ip + ":8080")
		}
	}
	// 获取外网ip
	ret = base.Exec("curl ifconfig.io")
	if ret != "" && !strings.HasPrefix(ret, "curl") {
		fmt.Println("Public Network IP: " + ret)
	}
}

func exe(e, key, port string) int {
	ret := base.Exec("ifconfig")
	en0 := strings.Index(ret, e)
	if en0 < 0 {
		return en0
	}
	inet := en0 + strings.Index(ret[en0:], key) + 5
	end := inet + strings.Index(ret[inet:], " ")
	ret = ret[inet:end]
	addr := strings.Index(ret, "addr:")
	if addr > 0 {
		ret = ret[addr+5:]
	}
	fmt.Println("Local IP: " + ret)
	fmt.Println("HTTP Server: http://" + ret)
	fmt.Println("HTTP Server: http://" + ret + ":8080")
	if port != "" {
		fmt.Println("HTTP Server: http://" + ret + ":" + port)
	}
	// 获取外网ip
	ret = base.Exec("curl ifconfig.io")
	if ret != "" && !strings.HasPrefix(ret, "curl") {
		fmt.Println("Public Network IP: " + ret)
	}
	return 0
}
