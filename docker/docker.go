package docker

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"strconv"
	"strings"
	"xx/base"
)

func Docker(cmd,a1,a2,a3 string){
	switch cmd {
	case "i","images": images(a1)
	case "e","exec": bash(a1)
	case "l","log": logs(a1, a2)
	case "s","start": start(a1)
	case "rs","restart": restart(a1)
	case "k","stop": stop(a1)
	case "pl","pull": pull(a1)
	case "ph","push": push(a1)
	case "t","tag": tag(a1, a2)
	case "rm": removeall(a1)
	case "ps": psall(a1)
	case "r","run": run(a1, a2, a3)
	default: fmt.Println("command not found: d"+cmd)
	}
}

// 启动容器
func run(imageName, containerName, port string){
	name := findImageName(imageName)
	if name == "" {
		fmt.Println("Image not found!")
		return
	}
	if strings.Index(port, ":") < 0 && port != "" && port != "@"{
		port = port+":"+port
	} 
	runCmd := "docker run -d"
	if containerName != "@" && containerName != "" {
		runCmd = runCmd + " --name="+containerName
	}
	if port != "@" && port != "" {
		runCmd = runCmd + " -p "+port
	}
	runCmd = runCmd+" "+name
	base.Run(runCmd)
}
// 删除images 及所有container
func removeall(key string){
	name := findImageName(key)
	if name == "" {
		return
	}
	prompt:= &survey.Confirm{
		Message:  "Are you sure remove all container and images by : "+name,
	}
	yesOrNo := false
	survey.AskOne(prompt, &yesOrNo)
	if yesOrNo {
		ret := base.Exec("docker ps -a")
		ret = base.FindByKey(ret, name)
		if len(ret) > 0{
			rets := strings.Split(ret, "\n")
			for i:=0; i<len(rets); i++ {
				strs := strings.Fields(rets[i])
				if len(strs) == 0 {
					continue
				}
				rmCmd:= "docker rm "+strs[0]
				base.Execp(rmCmd)
			}
		}
		startCmd := "docker rmi "+name
		base.Execp(startCmd)
	}
}
// 打tag
func tag(key string, t string){
	name := findImageName(key)
	if name == "" {
		return
	}
	startCmd := "docker tag "+ name +" "+t
	base.Run(startCmd)
}
// 推送镜像
func push(key string){
	name := findImageName(key)
	if name == "" {
		return
	}
	startCmd := "docker push "+name
	base.Run(startCmd)
}

// 拉取镜像
func pull(key string){
	if key == "" {
		return
	}
	startCmd := "docker pull "+key
	base.Run(startCmd)
}

// 停止容器
func stop(key string)  {
	name := findPsAllName(key)
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
		base.Run(startCmd)
	}
}
// 重启容器
func restart(key string)  {
	name := findPsAllName(key)
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
	name := findPsAllName(key)
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
	name := findPsAllName(key)
	if name == "" {
		return
	}
	cmd := "docker logs --tail "+line+" -f "+name
	base.Run(cmd)
}

// 进入容器bash
func bash(a1 string){
	name := findPsAllName(a1)
	if name == "" {
		return
	}
	cmd := "docker exec -it "+name+" /bin/bash"
	base.Run(cmd)
}
// 查找相关容器名称
func findPsAllName(key string) string{
	cmd := "docker ps -a"
	return findNameByCmd(key, cmd)
}
// 查询镜像名称
func findImageName(key string) string{
	cmd := "docker images"
	return findNameByCmd(key, cmd)
}
// 基础查找名称
func findNameByCmd(key string, cmd string) string{
	if key == "@" {
		key = ""
	}
	var rs []string
	ret := base.Exec(cmd)
	rets := strings.Split(ret, "\n")
	for i:=1; i<len(rets); i++{
		strs := strings.Fields(rets[i])
		if len(strs) == 0 {
			continue
		}
		s1 := strs[0]  // images name
		s2 := strs[len(strs)-1] // container name
		s3 := strs[1] // ps -a 时 images name
		if strings.Index(s1, key) >=0 {
			s1 = s1+":"+strs[1]
			rs = append(rs, s1)
		} else if strings.Index(s2, key) >=0 {
			rs = append(rs, s2)
		} else if strings.Index(s3, key) >=0 {
			s3 = strs[0]
			rs = append(rs, s3)
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
		if name=="" || name=="q" {
			return ""
		}
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
// 查看所有容器
func psall(a1 string){
	ret := base.Exec("docker ps -a")
	if a1 != "" {
		ret = base.FindByKey(ret, a1)
	}
	fmt.Println(ret)
}
