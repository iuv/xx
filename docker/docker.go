package docker

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"strconv"
	"strings"
	"xx/base"
)

func Docker(cmd,a1,a2 string){
	switch cmd {
	case "i": images(a1)
	case "e": bash(a1)
	case "l": logs(a1, a2)
	case "s": start(a1)
	case "r": restart(a1)
	case "k": stop(a1)
	default: fmt.Println("command not found: d"+cmd)
	}
}
// 停止容器
func stop(key string)  {
	name := findName(key)
	if name == "" {
		return
	}
	prompt:= &survey.Confirm{
		Message:  "Are you sure stop : "+name,
	}
	yesOrNo := false
	survey.AskOne(prompt, &yesOrNo)
	if yesOrNo {
		startCmd := "docker stop "+name
		base.Execp(startCmd)
	}
}
// 重启容器
func restart(key string)  {
	name := findName(key)
	if name == ""{
		return
	}
	prompt:= &survey.Confirm{
		Message:  "Are you sure restart : "+name,
	}
	yesOrNo := false
	survey.AskOne(prompt, &yesOrNo)
	if yesOrNo {
		startCmd := "docker restart "+name
		logCmd := "docker logs --tail 100 -f "+name
		base.Execp(startCmd)
		base.Run(logCmd)
	}
}
// 启动容器
func start(key string)  {
	name := findName(key)
	if name == ""{
		return
	}
	startCmd := "docker start "+name
	logCmd := "docker logs --tail 100 -f "+name
	base.Execp(startCmd)
	base.Run(logCmd)
}
// 查看容器控制台日志 a1：容器名（模糊） a2:查看行数，默认100
func logs(key, line string)  {
	if line == "" {
		line = "100"
	}
	name := findName(key)
	if name == "" {
		return
	}
	cmd := "docker logs --tail "+line+" -f "+name
	base.Run(cmd)
}

// 进入容器bash
func bash(a1 string){
	name := findName(a1)
	if name == "" {
		return
	}
	cmd := "docker exec -it "+name+" /bin/bash"
	base.Run(cmd)
}
// 查找相交容器名称
func findName(key string) string{
	if key == "@" {
		key = ""
	}
	var rs []string
	ret := base.Exec("docker ps -a")
	rets := strings.Split(ret, "\n")
	for i:=1; i<len(rets)-1; i++{
		strs := strings.Fields(rets[i])
		if len(strs) == 0 {
			continue
		}
		s1 := strs[0]
		s2 := strs[len(strs)-1]
		if strings.Index(s1, key) >=0 {
			rs = append(rs, s1)
		}
		if strings.Index(s2, key) >=0 {
			rs = append(rs, s2)
		}
	}
	if len(rs) == 1{
		return rs[0]
	}
	name := ""
	if len(rs) == 0 {
		fmt.Println("Not Found!")
		return name
	}
	var prompt survey.Prompt
	if len(rs)<=base.SEL_LIMIT {
		prompt = &survey.Select{
			Message: "Choose a option:",
			Options: rs,
		}
	} else {
		msg := ""
		for i,s := range rs {
			msg += strconv.Itoa(i)+": "+s+"\n"
		}
		prompt = &survey.Input{
			Message:  "The list:\n"+msg+"Input index:",
		}
	}
	survey.AskOne(prompt, &name)

	if len(rs)>base.SEL_LIMIT {
		idx, _ :=strconv.Atoi(name)
		name = rs[idx]
	}
	return name
}
// 查询镜像
func images(a1 string){
	ret := base.Exec("docker images")
	if a1 != "" {
		ret = base.FindByKey(ret, a1)
	}
	fmt.Println(ret)
}
