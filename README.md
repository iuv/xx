# xx 简化命令工具 v1.0
本工具为简化常用shell、docker、kubernetes命令使用

0.X版本使用python编写，需要运行环境支持py及需要sh脚本支持运行，为了解决运行环境依赖及保持单文件执行

1.X及以后版本将使用go语言编写

如需要自己编译，可clone本仓库并运行build.sh脚本
## 安装方法
1. 使用 `wget https://raw.githubusercontent.com/iuv/xx/master/xx ; bash xx install` 下载xx文件
2. 可以使用 `xx` 命令了
3. 更新使用 `xx update` 命令

## 使用帮助如下：

### shell:
1. xx ip [port] 获取本地ip
```
 示例结果：
 Local IP: 172.16.112.12
 HTTP Server: http://172.16.112.12
 HTTP Server: http://172.16.112.12:8080
 可选输出
 HTTP Server: http://172.16.112.12:[port]
```
2. xx ps [str] 获取进程，根据str模糊搜索，并高亮显示

### docker(参数为空且需要后续参数时使用"@"占位):
1、查询docker日志命令，使用

```shell
xx dl [dockername] [lines]
```

查询容器输出日志，dockername 支持模糊搜索 ,lines 输出行数, 默认100行

2、进入docker bash 命令，使用
```shell
xx de [dockername]
```
登入容器bash，dockername 支持模糊搜索

3、启动docker 命令，使用
```shell
xx ds [dockername]
```
启动容器，dockername 支持模糊搜索

4、重启docker 命令，使用
```shell
xx dr [dockername]
```
重新启动容器，dockername 支持模糊搜索

5、停止docker 命令，使用
```shell
xx dk [dockername]
```
停止容器，dockername 支持模糊搜索

### k8s(参数为空且需要后续参数时使用"@"占位):
1、查询namespace 命令，使用
```shell
xx kn [keyword]
```
keyword 支持模糊搜索

2、查询pod 命令，使用
```shell
xx kp [keyword] [namespace]
```
keyword 模糊匹配pod，如要查询全部 namespace 命名空间支持模糊匹配

3、进入pod bash 命令，使用
```shell
xx ke [pod] [namespace] [sh]
```
登入pod bash，pod pod名称支持模糊搜索，namespace 所属命名空间支持模糊, sh 默认为bash,有特殊可传入（/bin/目录下）

4、查询pod 日志命令，使用
```shell
xx kl [pod] [namespace] [lines]
```
查询pod日志，pod名称支持模糊搜索，namespace 所属命名空间支持模糊, lines 输出行数，默认100行

5、查询deployments 命令，使用
```shell
xx kd [deployment] [namespace]
```
deployment 名称支持模糊，namespace 命名空间支持模糊

6、查询ingress 命令，使用
```shell
xx ki [ingress] [namespace]
```
ingress 名称支持模糊，namespace 命名空间支持模糊

7、查询service 命令，使用
```shell
xx ks [service] [namespace]
```
service名称支持模糊，namespace 命名空间支持模糊

8、查询pod describe命令，使用
```shell
xx kpd [pod] [namespace]
```
pod名称支持模糊搜索，namespace 命名空间支持模糊

9、查询ingress describe命令，使用
```shell
xx kid [ingress] [namespace]
```
ingress名称支持模糊搜索，namespace 命名空间支持模糊

10、查询service describe命令，使用
```shell
xx ksd [service] [namespace]
```
service名称支持模糊搜索，namespace 命名空间支持模糊

11、查询deployment describe命令，使用
```shell
xx kdd [deployment] [namespace]
```
deployment名称支持模糊搜索，namespace 命名空间支持模糊

12、查询pod yaml命令，使用
```shell
xx kpy [pod] [namespace] [file]
```
pod名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

13、查询ingress yaml命令，使用
```shell
xx kiy [ingress] [namespace] [file]
```
ingress名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

14、查询service describe命令，使用
```shell
xx ksy [service] [namespace] [file]
```
service名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名

15、查询deployment yaml命令，使用
```shell
xx kdy [deployment] [namespace] [file]
```
deployment名称支持模糊搜索，namespace 命名空间支持模糊，file 保存到文件名
  
16、 删除pod命令，使用
```shell
xx kpdel [pod] [namespace]
```
pod名称支持模糊搜索，namespace 命名空间支持模糊

17、查询ingress命令，使用
```shell
xx kidel [ingress] [namespace]
```
ingress名称支持模糊搜索，namespace 命名空间支持模糊

18、删除service命令，使用
```shell
xx ksdel [service] [namespace]
```
service名称支持模糊搜索，namespace 命名空间支持模糊

19、删除deployment命令，使用
```shell
xx kddel [deployment] [namespace]
```
deployment名称支持模糊搜索，namespace 命名空间支持模糊

20、应用yaml配置文件命令，使用
```shell
xx ka [file]
```
file yaml配置文件

21、从pod容器中复制文件命令，使用
```shell
xx kcp [pod] [namespace] [srcFile] [saveFile]
```
pod名称支持模糊搜索，namespace 命名空间支持模糊, srcFile 容器中要复制的文件路径, saveFile 本地保存路径
