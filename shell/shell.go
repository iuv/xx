package shell

import(
	"fmt"
)

func Shell(cmd, s1, s2, s3, s4 string){
	switch cmd {
	case "up": sup(s1)
	default: fmt.Println("command not found: s"+cmd)
	}
}

func sup(s1 string){
	// 处理上一条输入
}