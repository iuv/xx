package shell

import (
	"fmt"
	"strings"
	"xx/base"
)

func Ip(port string){
	ret := base.Exec("ifconfig")
	en0 := strings.Index(ret, "en0");
	inet := en0+strings.Index(ret[en0:], "inet ")+5
	end := inet+strings.Index(ret[inet:], " ")
	ret = ret[inet:end]
	fmt.Println("Local IP: "+ret)
	fmt.Println("HTTP Server: http://"+ret)
	fmt.Println("HTTP Server: http://"+ret+":8080")
	if port != "" {
		fmt.Println("HTTP Server: http://"+ret+":"+port)
	}
}
