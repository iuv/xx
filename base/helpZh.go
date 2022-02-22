package base

import (
	"fmt"
)

var titleZh string = "        xx 简化命令工具 " + Version
var describeZh = `
本工具为简化常用shell、docker、kubernetes命令使用
有使用问题请访问https://github.com/iuv/xx 提交issues,
使用帮助如下：
`
var shellZh string = `
shell: 
  1. xx ip [port] 获取本地ip及公网ip(如果有外网)
  示例结果：
  Local IP: 172.16.112.12
  HTTP Server: http://172.16.112.12
  HTTP Server: http://172.16.112.12:8080
  Public Network IP: 8.8.8.8
  可选输出
  HTTP Server: http://172.16.112.12:[port]

  2. xx ps [str] 获取进程，根据str模糊搜索，并高亮显示
`
var dockerZh string = `
docker(参数为空且需要后续参数时使用"@"占位):
  1、运行docker命令，使用
    xx dr [imageName] [containerName] [port]
    或
    xx drun [imageName] [containerName] [port]
    默认使用后台进程启动 imageName 镜像名支持模糊搜索，containerName 设置容器名，port 映射的端口号
    支持"8080:8080"和"8080"两种方式，其"8080"会自动补全为"8080:8080"

  2、查询docker容器日志命令，使用
    xx dl [dockername] [lines]
    或
    xx dlog [dockername] [lines]
    查询容器输出日志，dockername 支持镜像/容器名模糊搜索 ,lines 输出行数, 默认100行

  3、进入docker bash命令，使用
    xx de [dockername]
    或
    xx dexec [dockername]
    进入容器bash，dockername 支持镜像/容器名模糊搜索

  4、启动docker容器命令，使用
    xx ds [dockername]
    或
    xx dstart [dockername]
    启动容器，dockername 支持镜像/容器名模糊搜索

  5、重启docker 命令，使用
    xx drs [dockername]
    或
    xx drestart [dockername]
    重新启动容器，dockername 支持镜像/容器名模糊搜索

  6、停止docker 命令，使用
    xx dk [dockername]
    或
    xx dstop [dockername]
    停止容器，dockername 支持镜像/容器名模糊搜索

  7、查找docker镜像，使用
    xx di [imageName]
    或
    xx dimages [imageName]
    查找镜像，imageName 支持模糊搜索

  8、拉取docker镜像，使用
    xx dpl [imageName]
    或
    xx dpull [imageName]
    拉取镜像，imageName，镜像全路径

  9、推送docker镜像，使用
    xx dph [imageName]
    或
    xx dpush [imageName]
    推送镜像，imageName 支持模糊搜索

  10、docker镜像打tag，使用
    xx dt [imageName] [tagname]
    或
    xx dtag [imageName] [tagname]
    镜像打tag，imageName 支持模糊搜索， tagname 需要打的tag名称

  11、docker查看所有容器，使用
    xx dps [dockername]
    查看所有容器（运行中和停止的），dockername 支持镜像/容器名模糊搜索

  12、docker删除镜像及使用该镜像启动的容器，使用
    xx drm [imageName]
    删除镜像及使用该镜像启动的所有容器 ，imageName 支持模糊搜索

  13、docker本地-容器互相复制文件，使用
    # 容器内文件复制到本地
    xx dc [dockerName]:[filePath] [localPath]
    # 本地文件复制到容器内
    xx dc [localPath] [dockerName]:[filePath]
    或
    # 容器内文件复制到本地
    xx dcp [dockerName]:[filePath] [localPath]
    # 本地文件复制到容器内
    xx dcp [localPath] [dockerName]:[filePath]
    docker本地-容器互相复制文件 ，dockerName 容器名支持模糊搜索 filePath 容器内文件/文件夹路径 localPath 本地文件路径  eg: xx dc mysql:/tmp/a.sql .
    
  14、docker将镜像保存为本地文件，使用
    xx dsa [imageName] [fileName]
    或
    xx dsave [imageName] [fileName]
    docker将镜像保存为本地文件，imageName 支持模糊搜索 fileName 保存的文件名
    
  15、docker从本地文件导入镜像，使用
    xx dlo [fileName]
    或
    xx dload [fileName]
    docker从本地文件导入镜像，fileName 需要导入的文件名
    
  16、docker将运行的容器保存为镜像，使用
    xx dco [dockerName] [imageName]
    或
    xx dcommit [dockerName] [imageName]
    docker将运行的容器保存为镜像，dockerName 容器名称支持模糊搜索 imageName 保存的镜像名
    
  17、docker查看镜像创建历史，使用
    xx dh [imageName]
    或
    xx dhistory [imageName]
    docker查看镜像创建历史，imageName 镜像名支持模糊搜索
  18、docker构建镜像(在Dockerfile所在目录下执行)
	xx db [imageName]
	或
	xx dbuild [imageName]
    docker构建镜像，在Dockerfile所在目录下执行，imageName为镜像名
`
var k8sZh = `
k8s(参数为空且需要后续参数时使用"@"占位):
  1、查询namespace 命令，使用
    xx kn [keyword]
    或
    xx kns [keyword]
    或
    xx knamespace [keyword]
    keyword 支持模糊搜索

  2、查询pod 命令，使用
    xx kp [keyword] [namespace]
    或
    xx kpod [keyword] [namespace]
    keyword 模糊匹配pod，如要查询全部 namespace 命名空间支持模糊匹配

  3、进入pod bash 命令，使用
    xx ke [pod] [namespace] [sh]
    或
    xx kexe [pod] [namespace] [sh]
    登入pod bash，pod pod名称支持模糊搜索，namespace 所属命名空间支持模糊, sh 默认为bash,有特殊可传入（/bin/目录下）

  4、查询pod 日志命令，使用
    xx kl [pod] [namespace] [lines]
    或
    xx klog [pod] [namespace] [lines]
    查询pod日志，pod名称支持模糊搜索，namespace 所属命名空间支持模糊, lines 输出行数，默认100行

  5、查询deployments 命令，使用
    xx kd [deployment] [namespace]
    或
    xx kdeployment [deployment] [namespace]
    deployment 名称支持模糊，namespace 命名空间支持模糊

  6、查询ingress 命令，使用
    xx ki [ingress] [namespace]
    或
    xx kingress [ingress] [namespace]
    ingress 名称支持模糊，namespace 命名空间支持模糊

  7、查询service 命令，使用
    xx ks [service] [namespace]
    或
    xx kservice [service] [namespace]
    service名称支持模糊，namespace 命名空间支持模糊

  8、查询configmap 命令，使用
    xx kc [configmap] [namespace]
    或
    xx kconfigmap [configmap] [namespace]
    configmap名称支持模糊，namespace 命名空间支持模糊

  9、查询secret 命令，使用
    xx ksec [secret] [namespace]
    或
    xx ksecret [secret] [namespace]
    secret名称支持模糊，namespace 命名空间支持模糊

  10、查询statefulset 命令，使用
    xx kss [statefulset] [namespace]
    或
    xx kstatefulset [statefulset] [namespace]
    statefulset名称支持模糊，namespace 命名空间支持模糊

  11、查询pod describe命令，使用
    xx kpd [pod] [namespace]
    或
    xx kpodd [pod] [namespace]
    pod名称支持模糊搜索，namespace 命名空间支持模糊

  12、查询ingress describe命令，使用
    xx kid [ingress] [namespace]
    或
    xx kingressd [ingress] [namespace]
    ingress名称支持模糊搜索，namespace 命名空间支持模糊

  13、查询service describe命令，使用
    xx ksd [service] [namespace]
    或
    xx kserviced [service] [namespace]
    service名称支持模糊搜索，namespace 命名空间支持模糊

  14、查询deployment describe命令，使用
    xx kdd [deployment] [namespace]
    或
    xx kdeploymentd [deployment] [namespace]
    deployment名称支持模糊搜索，namespace 命名空间支持模糊

  15、查询configmap describe命令，使用
    xx kcd [configmap] [namespace]
    或
    xx kconfigmapd [configmap] [namespace]
    configmap名称支持模糊搜索，namespace 命名空间支持模糊

  16、查询secret describe命令，使用
    xx ksecd [secret] [namespace]
    或
    xx ksecretd [secret] [namespace]
    secret名称支持模糊搜索，namespace 命名空间支持模糊

  17、查询statefulset describe命令，使用
    xx kssd [statefulset] [namespace]
    或
    xx kstatefulsetd [statefulset] [namespace]
    statefulset名称支持模糊搜索，namespace 命名空间支持模糊

  18、保存pod yaml命令，使用
    xx kpy [pod] [namespace] [file]
    或
    xx kpody [pod] [namespace] [file]
    pod名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

  19、保存ingress yaml命令，使用
    xx kiy [ingress] [namespace] [file]
    或
    xx kingressy [ingress] [namespace] [file]
    ingress名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

  20、保存service describe命令，使用
    xx ksy [service] [namespace] [file]
    或
    xx kservicey [service] [namespace] [file]
    service名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

  21、保存deployment yaml命令，使用
    xx kdy [deployment] [namespace] [file]
    或
    xx kdeploymenty [deployment] [namespace] [file]
    deployment名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

  22、保存configmap yaml命令，使用
    xx kcy [configmap] [namespace] [file]
    或
    xx kconfigmapy [configmap] [namespace] [file]
    configmap名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

  23、保存secret yaml命令，使用
    xx ksecy [secret] [namespace] [file]
    或
    xx ksecrety [secret] [namespace] [file]
    secret名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

  24、保存statefulset yaml命令，使用
    xx kssy [statefulset] [namespace] [file]
    或
    xx kstatefulsety [statefulset] [namespace] [file]
    statefulset名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名
    
  25、 删除pod命令，使用
    xx kpdel [pod] [namespace]
    或
    xx kpoddel [pod] [namespace]
    pod名称支持模糊搜索，namespace 命名空间支持模糊

  26、删除ingress命令，使用
    xx kidel [ingress] [namespace]
    或
    xx kingressdel [ingress] [namespace]
    ingress名称支持模糊搜索，namespace 命名空间支持模糊

  27、删除service命令，使用
    xx ksdel [service] [namespace]
    或
    xx kservicedel [service] [namespace]
    service名称支持模糊搜索，namespace 命名空间支持模糊

  28、删除deployment命令，使用
    xx kddel [deployment] [namespace]
    或
    xx kdeploymentdel [deployment] [namespace]
    deployment名称支持模糊搜索，namespace 命名空间支持模糊

  29、删除configmap命令，使用
    xx kcdel [configmap] [namespace]
    或
    xx kconfigmapdel [configmap] [namespace]
    configmap名称支持模糊搜索，namespace 命名空间支持模糊

  30、删除secret命令，使用
    xx ksecdel [secret] [namespace]
    或
    xx ksecretdel [secret] [namespace]
    secret名称支持模糊搜索，namespace 命名空间支持模糊

  31、删除statefulset命令，使用
    xx kssdel [statefulset] [namespace]
    或
    xx kstatefulsetdel [statefulset] [namespace]
    statefulset名称支持模糊搜索，namespace 命名空间支持模糊

  32、应用yaml配置文件命令，使用
    xx ka [file]
    或
    xx kapply [file]
    file yaml配置文件

  33、从pod容器中复制文件命令，使用
    xx kcopy [pod] [namespace] [srcFile] [saveFile]
    pod名称支持模糊搜索，namespace 命名空间支持模糊, srcFile 容器中要复制的文件路径, saveFile 本地保存路径
`

func HelpZh(key string) {
	fmt.Println(titleZh)
	fmt.Print(describeZh)
	switch key {
	case "shell":
		fmt.Print(shellZh)
	case "docker":
		fmt.Print(dockerZh)
	case "k8s":
		fmt.Print(k8sZh)
	default:
		fmt.Print(shellZh)
		fmt.Print(dockerZh)
		fmt.Print(k8sZh)
	}
}
