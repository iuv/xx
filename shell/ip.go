package shell

import (
	"fmt"
	"strings"
	"xx/base"
)

func Ip(port string){
	sys := base.Sys();
	if sys == "mac" {
		exe("en0", "inet ", port)
	} else if sys == "linux"{
		exe("eth0", "inet ", port)
	}
}

func exe(e, key, port string )  {
	ret := base.Exec("ifconfig")
	en0 := strings.Index(ret, e)
	inet := en0+strings.Index(ret[en0:], key)+5
	end := inet+strings.Index(ret[inet:], " ")
	ret = ret[inet:end]
	addr := strings.Index(ret, "addr:")
	if addr>0 {
		ret = ret[addr+5:]
	}
	fmt.Println("Local IP: "+ret)
	fmt.Println("HTTP Server: http://"+ret)
	fmt.Println("HTTP Server: http://"+ret+":8080")
	if port != "" {
		fmt.Println("HTTP Server: http://"+ret+":"+port)
	}
	// 获取外网ip
	ret = base.Exec("curl ifconfig.io")
	if ret != "" && !strings.HasPrefix(ret, "curl") {
		fmt.Println("Public Network IP: "+ret)
	}
}
