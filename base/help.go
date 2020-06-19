package base

import (
	"fmt"
)

var title string = "        xx 简化命令工具 "+Version
var describe = `
本工具为简化常用shell、docker、kubernetes命令使用
有使用问题请访问https://github.com/iuv/xx 提交issues,
使用帮助如下：
`
var shell string = `
shell: 
  1. xx ip [port] 获取本地ip
    示例结果：
    Local IP: 172.16.112.12
    HTTP Server: http://172.16.112.12
    HTTP Server: http://172.16.112.12:8080
    # 可选输出
    HTTP Server: http://172.16.112.12:[port]
  2. xx ps [str] 获取进程，根据str模糊搜索，并高亮显示
`
var docker string = `
docker(参数为空且需要后续参数时使用"@"占位):
  1、查询docker日志命令，使用
    xx dl [dockername] [lines]
    查询容器输出日志，dockername 支持模糊搜索 ,lines 输出行数, 默认100行
  
  2、进入docker bash 命令，使用
    xx de [dockername]
    进入容器bash，dockername 支持模糊搜索
  
  3、启动docker 命令，使用
    xx ds [dockername]
    启动容器，dockername 支持模糊搜索
  
  4、重启docker 命令，使用
    xx dr [dockername]
    重新启动容器，dockername 支持模糊搜索
  
  5、停止docker 命令，使用
    xx dk [dockername]
    停止容器，dockername 支持模糊搜索

  6、查找docker镜像，使用
    xx di [dockername]
    查找镜像，dockername 支持模糊搜索

  7、拉取docker镜像，使用
    xx dpl [dockername]
    拉取镜像，dockername，镜像全路径

  8、推送docker镜像，使用
    xx dph [dockername]
    推送镜像，dockername 支持模糊搜索

  9、docker镜像打tag，使用
    xx dt [dockername] [tagname]
    镜像打tag，dockername 支持模糊搜索， tagname 需要打的tag名称
`
var k8s = `
k8s(参数为空且需要后续参数时使用"@"占位):
  1、查询namespace 命令，使用
    xx kn [keyword]
    keyword 支持模糊搜索
  
  2、查询pod 命令，使用
    xx kp [keyword] [namespace]
    keyword 模糊匹配pod，如要查询全部 使用"@"占位
    namespace 命名空间支持模糊匹配
  
  3、进入pod bash 命令，使用
    xx ke [pod] [namespace] [sh]
    登入pod bash，pod pod名称支持模糊搜索，namespace 所属命名空间支持模糊, sh 默认为bash,有特殊可传入（/bin/目录下）
  
  4、查询pod 日志命令，使用
    xx kl [pod] [namespace] [lines]
    查询pod日志，pod名称支持模糊搜索，namespace 所属命名空间支持模糊, lines 输出行数，默认100行
  
  5、查询deployments 命令，使用
    xx kd [deployment] [namespace]
    deployment 名称支持模糊，namespace 命名空间支持模糊
  
  6、查询ingress 命令，使用
    xx ki [ingress] [namespace]
    ingress 名称支持模糊，namespace 命名空间支持模糊
  
  7、查询service 命令，使用
    xx ks [service] [namespace]
    service名称支持模糊，namespace 命名空间支持模糊
  
  8、查询pod describe命令，使用
    xx kpd [pod] [namespace]
    pod名称支持模糊搜索，namespace 命名空间支持模糊
  
  9、查询ingress describe命令，使用
    xx kid [ingress] [namespace]
    ingress名称支持模糊搜索，namespace 命名空间支持模糊
  
  10、查询service describe命令，使用
    xx ksd [service] [namespace]
    service名称支持模糊搜索，namespace 命名空间支持模糊
  
  11、查询deployment describe命令，使用
    xx kdd [deployment] [namespace]
    deployment名称支持模糊搜索，namespace 命名空间支持模糊

  12、查询pod yaml命令，使用
    xx kpy [pod] [namespace] [file]
    pod名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名
  
  13、查询ingress yaml命令，使用
    xx kiy [ingress] [namespace] [file]
    ingress名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名
  
  14、查询service describe命令，使用
    xx ksy [service] [namespace] [file]
    service名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名
  
  15、查询deployment yaml命令，使用
    xx kdy [deployment] [namespace] [file]
    deployment名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名
  16、 删除pod命令，使用
    xx kpdel [pod] [namespace]
    pod名称支持模糊搜索，namespace 命名空间支持模糊
  
  17、查询ingress命令，使用
    xx kidel [ingress] [namespace]
    ingress名称支持模糊搜索，namespace 命名空间支持模糊
  
  18、删除service命令，使用
    xx ksdel [service] [namespace]
    service名称支持模糊搜索，namespace 命名空间支持模糊
  
  19、删除deployment命令，使用
    xx kddel [deployment] [namespace]
    deployment名称支持模糊搜索，namespace 命名空间支持模糊

  20、应用yaml配置文件命令，使用
    xx ka [file]
    file yaml配置文件

  21、从pod容器中复制文件命令，使用
    xx kcp [pod] [namespace] [srcFile] [saveFile]
    pod名称支持模糊搜索，namespace 命名空间支持模糊, srcFile 容器中要复制的文件路径, saveFile 本地保存路径
`
func Help(key string){
	fmt.Println(title)
	fmt.Print(describe)
	switch key {
	case "shell": fmt.Print(shell)
	case "docker": fmt.Print(docker)
	case "k8s": fmt.Print(k8s)
	default:
		fmt.Print(shell)
		fmt.Print(docker)
		fmt.Print(k8s)
	}
}
