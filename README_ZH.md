# xx 简化命令工具 v1.4.0
![License](https://img.shields.io/badge/license-MIT-4EB1BA)
[![wiki](https://img.shields.io/badge/Document-Wiki-green)](http://blog.jisuye.com/xx)

本工具为简化常用shell、docker、kubernetes命令使用

0.X版本使用python编写，需要运行环境支持py及需要sh脚本支持运行，为了解决运行环境依赖及保持单文件执行

1.X及以后版本将使用go语言编写

如需要自己编译，可clone本仓库并运行build.sh脚本
## 安装方法
1. Mac使用 `wget https://raw.githubusercontent.com/iuv/xx/master/build/mac/x86/xx` 下载xx文件(x86、arm64)
2. linux使用 `wget https://raw.githubusercontent.com/iuv/xx/master/build/linux/x86/xx` 下载xx文件(x86、arm、arm64)
3. 执行`chmod +x xx; ./xx install` 安装
4. 可以使用 `xx` 命令了
5. 更新使用 `xx update` 命令
6. 使用 `xx h` 获取帮助, `xx zh` 获取中文帮助

> 国内用户可尝试使用以下加速链接：  
> `wget https://gh.wget.cool/https:/raw.githubusercontent.com/iuv/xx/master/build/mac/x86/xx`   
> `wget https://gh.wget.cool/https:/raw.githubusercontent.com/iuv/xx/master/build/linux/x86/xx`   
> `wget https://cdn.jsdelivr.net/gh/iuv/xx@master/build/mac/x86/xx`  
> `wget https://cdn.jsdelivr.net/gh/iuv/xx@master/build/linux/x86/xx`

## 使用示例
![demo](https://raw.githubusercontent.com/iuv/xx/master/demo.gif?sanitize=true)

## 使用帮助如下：

![help](https://raw.githubusercontent.com/iuv/xx/master/xx_zh.svg?sanitize=true)

### shell:
1. xx ip [port] 获取本地ip及公网ip(如果有外网)
```
 示例结果：
 Local IP: 172.16.112.12
 HTTP Server: http://172.16.112.12
 HTTP Server: http://172.16.112.12:8080
 Public Network IP: 8.8.8.8
 可选输出
 HTTP Server: http://172.16.112.12:[port]
```
2. xx ps [str] 获取进程，根据str模糊搜索，并高亮显示

### docker(参数为空且需要后续参数时使用"@"占位):
1、运行docker命令，使用
```shell
xx dr [imageName] [containerName] [port]
或
xx drun [imageName] [containerName] [port]
```
默认使用后台进程启动 imageName 镜像名支持模糊搜索，containerName 设置容器名，port 映射的端口号
支持"8080:8080"和"8080"两种方式，其"8080"会自动补全为"8080:8080"

2、查询docker容器日志命令，使用

```shell
xx dl [dockername] [lines]
或
xx dlog [dockername] [lines]
```
查询容器输出日志，dockername 支持镜像/容器名模糊搜索 ,lines 输出行数, 默认100行

3、进入docker bash命令，使用
```shell
xx de [dockername]
或
xx dexec [dockername]
```
进入容器bash，dockername 支持镜像/容器名模糊搜索

4、启动docker容器命令，使用
```shell
xx ds [dockername]
或
xx dstart [dockername]
```
启动容器，dockername 支持镜像/容器名模糊搜索

5、重启docker 命令，使用
```shell
xx drs [dockername]
或
xx drestart [dockername]
```
重新启动容器，dockername 支持镜像/容器名模糊搜索

6、停止docker 命令，使用
```shell
xx dk [dockername]
或
xx dstop [dockername]
```
停止容器，dockername 支持镜像/容器名模糊搜索

7、查找docker镜像，使用
```shell
xx di [imageName]
或
xx dimages [imageName]
```
查找镜像，imageName 支持模糊搜索

8、拉取docker镜像，使用
```shell
xx dpl [imageName]
或
xx dpull [imageName]
```
拉取镜像，imageName，镜像全路径

9、推送docker镜像，使用
```shell
xx dph [imageName]
或
xx dpush [imageName]
```
推送镜像，imageName 支持模糊搜索

10、docker镜像打tag，使用
```shell
xx dt [imageName] [tagname]
或
xx dtag [imageName] [tagname]
```
镜像打tag，imageName 支持模糊搜索， tagname 需要打的tag名称

11、docker查看所有容器，使用
```shell
xx dps [dockername]
```
查看所有容器（运行中和停止的），dockername 支持镜像/容器名模糊搜索

12、docker删除镜像及使用该镜像启动的容器，使用
```shell
xx drm [imageName]
```
删除镜像及使用该镜像启动的所有容器 ，imageName 支持模糊搜索

13、docker本地-容器互相复制文件，使用
```shell
# 容器内文件复制到本地
xx dc [dockerName]:[filePath] [localPath]
# 本地文件复制到容器内
xx dc [localPath] [dockerName]:[filePath]
或
# 容器内文件复制到本地
xx dcp [dockerName]:[filePath] [localPath]
# 本地文件复制到容器内
xx dcp [localPath] [dockerName]:[filePath]
```
docker本地-容器互相复制文件 ，dockerName 容器名支持模糊搜索 filePath 容器内文件/文件夹路径 localPath 本地文件路径  eg: xx dc mysql:/tmp/a.sql .

14、docker将镜像保存为本地文件，使用
```shell
xx dsa [imageName] [fileName]
或
xx dsave [imageName] [fileName]
```
docker将镜像保存为本地文件，imageName 支持模糊搜索 fileName 保存的文件名

15、docker从本地文件导入镜像，使用
```shell
xx dlo [fileName]
或
xx dload [fileName]
```
docker从本地文件导入镜像，fileName 需要导入的文件名

16、docker将运行的容器保存为镜像，使用
```shell
xx dco [dockerName] [imageName]
或
xx dcommit [dockerName] [imageName]
```
docker将运行的容器保存为镜像，dockerName 容器名称支持模糊搜索 imageName 保存的镜像名

17、docker查看镜像创建历史，使用
```shell
xx dh [imageName]
或
xx dhistory [imageName]
```
docker查看镜像创建历史，imageName 镜像名支持模糊搜索

18、docker构建镜像(在Dockerfile所在目录下执行)
```shell
xx db [imageName]
或
xx dbuild [imageName]
```
docker构建镜像，在Dockerfile所在目录下执行，imageName为镜像名

### k8s(参数为空且需要后续参数时使用"@"占位):
#### 1.exec
1.1 进入pod bash 命令，使用
```shell
xx ke [pod] [namespace] [sh]
或
xx kexe [pod] [namespace] [sh]
```
登入pod bash，pod pod名称支持模糊搜索，namespace 所属命名空间支持模糊, sh 默认为bash,有特殊可传入（/bin/目录下）
#### 2.log
2.1 查询pod 日志命令，使用
```shell
xx kl [pod] [namespace] [lines]
或
xx klog [pod] [namespace] [lines]
```
查询pod日志，pod名称支持模糊搜索，namespace 所属命名空间支持模糊, lines 输出行数，默认100行

#### 3.query
3.1 查询namespace 命令，使用
```shell
xx kn [namespace]
或
xx kns [namespace]
或
xx knamespace [namespace]
```
namespace 支持模糊搜索

3.2 查询pod 命令，使用
```shell
xx kp [pod] [namespace]
或
xx kpod [pod] [namespace]
```
pod 模糊匹配pod，如要查询全部 namespace 命名空间支持模糊匹配

3.3 查询deployments 命令，使用
```shell
xx kd [deployment] [namespace]
或
xx kdeployment [deployment] [namespace]
```
deployment 名称支持模糊，namespace 命名空间支持模糊

3.4 查询ingress 命令，使用
```shell
xx ki [ingress] [namespace]
或
xx kingress [ingress] [namespace]
```
ingress 名称支持模糊，namespace 命名空间支持模糊

3.5 查询service 命令，使用
```shell
xx ks [service] [namespace]
或
xx kservice [service] [namespace]
```
service名称支持模糊，namespace 命名空间支持模糊

3.6 查询configmap 命令，使用
```shell
xx kc [configmap] [namespace]
或
xx kconfigmap [configmap] [namespace]
```
configmap名称支持模糊，namespace 命名空间支持模糊

3.7 查询secret 命令，使用
```shell
xx ksec [secret] [namespace]
或
xx ksecret [secret] [namespace]
```
secret名称支持模糊，namespace 命名空间支持模糊

3.8 查询statefulset 命令，使用
```shell
xx kss [statefulset] [namespace]
或
xx kstatefulset [statefulset] [namespace]
```
statefulset名称支持模糊，namespace 命名空间支持模糊

3.9 查询 CR 命令, 使用
```shell
xx kcr [cr] [cr key] [namespace]
```
cr 具体的资源类型 cr key 关键字支持模糊, namespace 命名空间支持模糊

### 4.describe
4.1 查询pod describe命令，使用
```shell
xx kpd [pod] [namespace] [key]
或
xx kpodd [pod] [namespace] [key]
```
pod名称支持模糊搜索，namespace 命名空间支持模糊, key 支持在返回内容中搜索关键字

4.2 查询ingress describe命令，使用
```shell
xx kid [ingress] [namespace] [key]
或
xx kingressd [ingress] [namespace] [key]
```
ingress名称支持模糊搜索，namespace 命名空间支持模糊, key 支持在返回内容中搜索关键字

4.3 查询service describe命令，使用
```shell
xx ksd [service] [namespace] [key]
或
xx kserviced [service] [namespace] [key]
```
service名称支持模糊搜索，namespace 命名空间支持模糊, key 支持在返回内容中搜索关键字

4.4 查询deployment describe命令，使用
```shell
xx kdd [deployment] [namespace] [key]
或
xx kdeploymentd [deployment] [namespace] [key]
```
deployment名称支持模糊搜索，namespace 命名空间支持模糊, key 支持在返回内容中搜索关键字

4.5 查询configmap describe命令，使用
```shell
xx kcd [configmap] [namespace] [key]
或
xx kconfigmapd [configmap] [namespace] [key]
```
configmap名称支持模糊搜索，namespace 命名空间支持模糊, key 支持在返回内容中搜索关键字

4.6 查询secret describe命令，使用
```shell
xx ksecd [secret] [namespace] [key]
或
xx ksecretd [secret] [namespace] [key]
```
secret名称支持模糊搜索，namespace 命名空间支持模糊, key 支持在返回内容中搜索关键字

4.7 查询statefulset describe命令，使用
```shell
xx kssd [statefulset] [namespace] [key]
或
xx kstatefulsetd [statefulset] [namespace] [key]
```
statefulset名称支持模糊搜索，namespace 命名空间支持模糊, key 支持在返回内容中搜索关键字

4.8 查询 CR describe 命令, 使用
```shell
xx kcrd [cr] [cr key] [namespace] [key]
```
cr 资源类型 ,cr key 关键字支持模糊搜索, namespace 命名空间支持模糊搜索, key 支持在返回内容中搜索关键字

#### 5.yaml
5.1 保存pod yaml命令，使用
```shell
xx kpy [pod] [namespace] [file]
或
xx kpody [pod] [namespace] [file]
```
pod名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

5.2 保存ingress yaml命令，使用
```shell
xx kiy [ingress] [namespace] [file]
或
xx kingressy [ingress] [namespace] [file]
```
ingress名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

5.3 保存service yaml命令，使用
```shell
xx ksy [service] [namespace] [file]
或
xx kservicey [service] [namespace] [file]
```
service名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

5.4 保存deployment yaml命令，使用
```shell
xx kdy [deployment] [namespace] [file]
或
xx kdeploymenty [deployment] [namespace] [file]
```
deployment名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

5.5 保存configmap yaml命令，使用
```shell
xx kcy [configmap] [namespace] [file]
或
xx kconfigmapy [configmap] [namespace] [file]
```
configmap名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

5.6 保存secret yaml命令，使用
```shell
xx ksecy [secret] [namespace] [file]
或
xx ksecrety [secret] [namespace] [file]
```
secret名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

5.7 保存statefulset yaml命令，使用
```shell
xx kssy [statefulset] [namespace] [file]
或
xx kstatefulsety [statefulset] [namespace] [file]
```
statefulset名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

5.8 保存cr yaml 命令, 使用
```shell
xx kcry [cr] [cr key] [namespace] [file]
```
cr 资源类型, cr key 支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

#### 6.delete
6.1 删除pod命令，使用
```shell
xx kpdel [pod] [namespace]
或
xx kpoddel [pod] [namespace]
```
pod名称支持模糊搜索，namespace 命名空间支持模糊

6.2 删除ingress命令，使用
```shell
xx kidel [ingress] [namespace]
或
xx kingressdel [ingress] [namespace]
```
ingress名称支持模糊搜索，namespace 命名空间支持模糊

6.3 删除service命令，使用
```shell
xx ksdel [service] [namespace]
或
xx kservicedel [service] [namespace]
```
service名称支持模糊搜索，namespace 命名空间支持模糊

6.4 删除deployment命令，使用
```shell
xx kddel [deployment] [namespace]
或
xx kdeploymentdel [deployment] [namespace]
```
deployment名称支持模糊搜索，namespace 命名空间支持模糊

6.5 删除configmap命令，使用
```shell
xx kcdel [configmap] [namespace]
或
xx kconfigmapdel [configmap] [namespace]
```
configmap名称支持模糊搜索，namespace 命名空间支持模糊

6.6 删除secret命令，使用
```shell
xx ksecdel [secret] [namespace]
或
xx ksecretdel [secret] [namespace]
```
secret名称支持模糊搜索，namespace 命名空间支持模糊

6.7 删除statefulset命令，使用
```shell
xx kssdel [statefulset] [namespace]
或
xx kstatefulsetdel [statefulset] [namespace]
```
statefulset名称支持模糊搜索，namespace 命名空间支持模糊

6.8 删除 cr 资源命令，使用
```shell
xx kcrdel [cr] [cr key] [namespace]
```
cr 资源类型, cr key 支持模糊搜索，namespace 命名空间支持模糊

#### 7.apply
7.1 应用yaml配置文件命令，使用
```shell
xx ka [file]
或
xx kapply [file]
```
file yaml配置文件

#### 8.copy
8.1 从pod容器中复制文件命令，使用
```shell
xx kcopy [pod] [namespace] [srcFile] [saveFile]
```
pod名称支持模糊搜索，namespace 命名空间支持模糊, srcFile 容器中要复制的文件路径, saveFile 本地保存路径

#### 9.edit
9.1 编辑pod命令，使用
```shell
xx kpe [pod] [namespace]
或
xx kpode [pod] [namespace]
```
pod名称支持模糊搜索，namespace 命名空间支持模糊

9.2 编辑ingress命令，使用
```shell
xx kie [ingress] [namespace]
或
xx kingresse [ingress] [namespace]
```
ingress名称支持模糊搜索，namespace 命名空间支持模糊

9.3 编辑service命令，使用
```shell
xx kse [service] [namespace]
或
xx kservicee [service] [namespace]
```
service名称支持模糊搜索，namespace 命名空间支持模糊

9.4 编辑deployment命令，使用
```shell
xx kde [deployment] [namespace]
或
xx kdeploymente [deployment] [namespace]
```
deployment名称支持模糊搜索，namespace 命名空间支持模糊

9.5 编辑configmap命令，使用
```shell
xx kce [configmap] [namespace]
或
xx kconfigmape [configmap] [namespace]
```
configmap名称支持模糊搜索，namespace 命名空间支持模糊

9.6 编辑secret命令，使用
```shell
xx ksece [secret] [namespace]
或
xx ksecrete [secret] [namespace]
```
secret名称支持模糊搜索，namespace 命名空间支持模糊

9.7 编辑statefulset命令，使用
```shell
xx ksse [statefulset] [namespace]
或
xx kstatefulsete [statefulset] [namespace]
```
statefulset名称支持模糊搜索，namespace 命名空间支持模糊

9.8 编辑 cr 资源命令，使用
```shell
xx kcre [cr] [cr key] [namespace]
```
cr 资源类型, cr key 支持模糊搜索，namespace 命名空间支持模糊

