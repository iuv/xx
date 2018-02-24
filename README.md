# xx
shell 简化命令，常用命令精简
## 安装方法
1. 使用 `wget https://raw.githubusercontent.com/iuv/xx/master/xx` 下载xx文件
2. 使用命令 `bash xx install` 安装文件
3. 可以使用 `xx` 命令了
4. 更新使用 `xx update` 命令

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
3.简化启动docker 命令，使用
```shell
xx ds [dockername]
```
启动容器，dockername 支持模糊搜索
4.简化重启docker 命令，使用
```shell
xx dr [dockername]
```
重新启动容器，dockername 支持模糊搜索
5.简化停止docker 命令，使用
```shell
xx dk [dockername]
```
停止容器，dockername 支持模糊搜索

