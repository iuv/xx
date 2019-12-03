package base

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

// 执行前输出命令
func Execp(cmdStr string) string  {
	fmt.Println(cmdStr)
	return Exec(cmdStr)
}
// 执行命令并返回内容
func Exec(cmdStr string) string{
	var ret string
	cmds := strings.Split(cmdStr, " ")
	cmd := exec.Command(cmds[0], cmds[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	ret += string(opBytes)
	return ret
}

// 执行命令并切换到命令进程
func Run(cmdStr string)  {
	fmt.Println(cmdStr)
	cmds := strings.Split(cmdStr, " ")
	cmd := exec.Command(cmds[0],cmds[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}
