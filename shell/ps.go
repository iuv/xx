package shell

import (
	"fmt"
	"xx/base"
)
func Ps(arg string){
	var ret string
	if arg != "" {
		ret = base.Exec("ps -ef")
		ret = base.FindByKey(ret, arg)
	} else {
		ret = base.Exec("ps")
	}
	fmt.Println(ret)
}
