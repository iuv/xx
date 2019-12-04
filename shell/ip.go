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
		exe("eth0", "addr:", port)
	}
}

func exe(e, key, port string )  {
	ret := base.Exec("ifconfig")
	en0 := strings.Index(ret, e)
	inet := en0+strings.Index(ret[en0:], key)+5
	end := inet+strings.Index(ret[inet:], " ")
	ret = ret[inet:end]
	fmt.Println("Local IP: "+ret)
	fmt.Println("HTTP Server: http://"+ret)
	fmt.Println("HTTP Server: http://"+ret+":8080")
	if port != "" {
		fmt.Println("HTTP Server: http://"+ret+":"+port)
	}
}
