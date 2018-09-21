# xx
shell 简化命令，常用命令精简
## 安装方法
1. 使用 `wget https://raw.githubusercontent.com/iuv/xx/master/xx ; bash xx install` 下载xx文件
2. 可以使用 `xx` 命令了
3. 更新使用 `xx update` 命令

## 使用方法
### docker 相关操作
1. 简化查询docker 命令，使用  
```shell
xx dl [dockername] [lines]
```
查询容器输出日志，`dockername` 支持模糊搜索 ,`lines` 输出行数, 默认10行

2. 简化进入docker bash 命令，使用  
```shell
xx de [dockername]
```
登入容器bash，`dockername` 支持模糊搜索

3. 简化启动docker 命令，使用
```shell
xx ds [dockername]
```
启动容器，dockername 支持模糊搜索

4. 简化重启docker 命令，使用
```shell
xx dr [dockername]
```
重新启动容器，dockername 支持模糊搜索

5. 简化停止docker 命令，使用
```shell
xx dk [dockername]
```
停止容器，dockername 支持模糊搜索

### k8s 相关操作

1. 简化查询namespace 命令，使用
```shell
xx kn
```

2. 简化查询pod 命令，使用
```shell
xx kp [namespace]
```
namespace 命名空间以!结尾支持模糊

3. 简化进入pod bash 命令，使用
```shell
xx ke [pod] [namespace]
```
登入pod bash，pod pod名称支持模糊搜索，namespace 所属命名空间以!结尾支持模糊

4. 简化查询pod 日志命令，使用
```shell
xx kl [pod] [namespace] [lines]
```
查询pod日志，pod名称支持模糊搜索，namespace 所属命名空间以!结尾支持模糊, lines 输出行数，默认10行

5. 简化查询deployments 命令，使用
```shell
xx kd [namespace]
```
namespace 命名空间以!结尾支持模糊

6. 简化查询ingress 命令，使用
```shell
xx ki [namespace]
```
namespace 命名空间以!结尾支持模糊

7. 简化查询service 命令，使用
```shell
xx ks [namespace]
```
namespace 命名空间以!结尾支持模糊

8. 简化查询pod describe命令，使用
```shell
xx kdp [pod] [namespace]
```
pod名称支持模糊搜索 namespace 命名空间以!结尾支持模糊

9. 简化查询ingress describe命令，使用
```shell
xx kdi [ingress] [namespace]
```
ingress名称支持模糊搜索 namespace 命名空间以!结尾支持模糊

10. 简化查询service describe命令，使用
```shell
xx kds [service] [namespace]
```
service名称支持模糊搜索 namespace 命名空间以!结尾支持模糊

11. 简化查询deployments describe命令，使用
```shell
xx kdd [deployments] [namespace]
```
service名称支持模糊搜索 namespace 命名空间以!结尾支持模糊
 
