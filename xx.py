#!/usr/bin/python
#-*- coding:utf-8 -*-
import sys,os

#处理docker相关操作
def docker(a1, a2, a3, a4):
    print("xxx:%s"%a1)
    print("xxx:%s"%a4)
    if(a1 == "dl"):
        dockerLog(a2, a3);
    elif(a1 == "de"):
        dockerExec(a2);
    elif(a1 == "dr"):
        dockerRestart(a2);
    elif(a1 == "ds"):
        dockerStart(a2);
    elif(a1 == "dk"):
        dockerStop(a2);
    elif(a1 == "kn"):
        k8sNamespace();
    elif(a1 == "kp"):
        k8sPod(a2);
    elif(a1 == "kl"):
        podLog(a2, a3, a4);
    elif(a1 == "ke"):
        podBash(a2, a3);
    return

# 查找docker窗口
def findDocker(name=" ", para = ""):
    dsprint = os.popen("docker ps %s | grep \"%s\"" % (para,name))
    res = dsprint.readlines()
    fixres = []
    dockername = ""
    for r in res:
        rs = r.split()
        if rs[len(rs)-1].find(name) >=0 :
            fixres.append(rs[len(rs)-1])
    if len(fixres) == 0:
        print("未找到容器，请检查输入")
    elif len(fixres) == 1:
        dockername = fixres[0]
    elif len(fixres) > 1:
        for i in range(0, len(fixres)):
            print("%d:%s" % (i+1,fixres[i]))
        t = int(raw_input("请选择容器："))
        if t > 0 and t <= len(fixres) :
            dockername = fixres[t-1]
    return dockername
#重启docker
# name 窗口名称支持模糊
def dockerRestart(name= " "):
    dockername = findDocker(name)
    if dockername != "":
        t = raw_input("确定重启容器%s 吗？[y/n]：" % dockername)
        if(t == "y" or t == "Y"):
            os.system("docker restart %s " % dockername)
            os.system("docker logs --tail %s -f %s" % (10,dockername))
    return

#启动docker
# name 窗口名称支持模糊
def dockerStart(name= " "):
    dockername = findDocker(name, "-a")
    if dockername != "":
        print("启动容器%s：" % dockername)
        os.system("docker start %s " % dockername)
        os.system("docker logs --tail %s -f %s" % (10,dockername))
    return

#停止docker
# name 窗口名称支持模糊
def dockerStop(name= " "):
    dockername = findDocker(name)
    if dockername != "":
        t = raw_input("确定停止容器%s 吗？[y/n]：" % dockername)
        if(t == "y" or t == "Y"):
            os.system("docker stop %s " % dockername)
    return

#处理docker 日志函数
# name 容器名称，支持模糊
# lines 显示log行数，默认10
def dockerLog(name = " ",lines = 10):
    if lines is None:
        lines = 10
    dockername = findDocker(name)
    if dockername != "":
        print("查看%s日志：" % dockername)
        os.system("docker logs --tail %s -f %s" % (lines,dockername))
    return

#进入docker bash
# name 容器名称，支持模糊
def dockerExec(name = " "):
    dockername = findDocker(name)
    if dockername != "":
        print("登入 %s bash：" % dockername)
        os.system("docker exec -it  %s bash" % dockername)
    return

def helps( version ):
    print """
        xx工具（%s）使用帮助
    了解更多请访问 https://github.com/iuv/xx
    docker 相关操作
      1.简化查询docker 命令，使用
        xx dl [dockername] [lines]
        查询容器输出日志，dockername 支持模糊搜索, lines 输出行数，默认10行
      2.简化进入docker bash 命令，使用
        xx de [dockername]
        登入容器bash，dockername 支持模糊搜索
      3.简化启动docker 命令，使用
        xx ds [dockername]
        启动容器，dockername 支持模糊搜索
      4.简化重启docker 命令，使用
        xx dr [dockername]
        重新启动容器，dockername 支持模糊搜索
      5.简化停止docker 命令，使用
        xx dk [dockername]
        停止容器，dockername 支持模糊搜索
    k8s 相关操作
      1.简化查询namespace 命令，使用
        xx kn
      2.简化查询pod 命令，使用
        xx kp [namespace]
        namespace 命名空间
      3.简化进入pod bash 命令，使用
        xx ke [pod] [namespace]
        登入pod bash，pod pod名称支持模糊搜索，namespace 所属命名空间
      4.简化查询pod 日志命令，使用
        xx kl [pod] [namespace] [lines]
        查询pod日志，pod名称支持模糊搜索，namespace 所属命名空间, lines 输出行数，默认10行
    """ % version
    return

def versions( version ):
    print("xx v%s" % version)

#处理k8s查看namespace
def k8sNamespace():
    print("查看namespace，命令：kubectl get namespace")
    os.system("kubectl get namespace")
#处理k8s查看pod
# namespace 命名空间
def k8sPod(namespace = ""):
    print("查看pod，命令：kubectl get pod %s" % (namespace))
    if namespace is None:
        namespace = ""
    else:
        namespace = " -n " + namespace
    os.system("kubectl get pod %s" % namespace)
    return

# 查找pod
def findPod(pod = "", namespace = ""):
    dsprint = os.popen("kubectl get pod %s | grep \"%s\"" % (namespace, pod))
    res = dsprint.readlines()
    fixres = []
    podname = ""
    for r in res:
        rs = r.split()
        if rs[0].find(pod) >=0 :
            fixres.append(rs[0])
    if len(fixres) == 0:
        print("未找到pod，请检查输入")
    elif len(fixres) == 1:
        podname = fixres[0]
    elif len(fixres) > 1:
        for i in range(0, len(fixres)):
            print("%d:%s" % (i+1,fixres[i]))
        t = int(raw_input("请选择pod："))
        if t > 0 and t <= len(fixres) :
            podname = fixres[t-1]
    return podname

#处理k8s pod 日志函数
# pod pod名称，支持模糊
# namespace 命名空间
# lines 显示log行数，默认10
def podLog(pod = "", namespace = "",lines = 10):
    if lines is None:
        lines = 10
    if namespace is None:
        namespace = ""
    else:
        namespace = " -n " + namespace
    podname = findPod(pod, namespace)
    if podname != "":
        command = "kubectl logs --tail %s -f %s %s" % (lines, podname, namespace);
        print("查看%s日志，命令：%s" % (pod, command))
        os.system(command)
    return

#处理k8s pod 进入bash
# pod pod名称，支持模糊
# namespace 命名空间
def podBash(pod = "", namespace = ""):
    if namespace != "":
        namespace = " -n " + namespace;
    podname = findPod(pod, namespace)
    if podname != "":
        command = "kubectl exec -it %s %s -- bash" % (podname, namespace)
        print("进入%s bash，命令：%s" % (pod, command))
        os.system(command)
    return


#主方法执行
version = "0.5" # 版本号
a1 = sys.argv[1]
a2 = None
a3 = None
a4 = None
if(len(sys.argv) == 3):
    a2 = sys.argv[2]
if(len(sys.argv) == 4):
    a2 = sys.argv[2]
    a3 = sys.argv[3]
if(len(sys.argv) == 5):
    a2 = sys.argv[2]
    a3 = sys.argv[3]
    a4 = sys.argv[4]
if(str(a1).startswith("d") or str(a1).startswith("k")):
    docker(a1,a2,a3,a4)
elif a1 == "-h" or a1 == "help":
    helps(version)
elif a1 == "-v" or a1 == "version":
    versions(version)
