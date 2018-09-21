#!/usr/bin/python
#-*- coding:utf-8 -*-
import sys,os

#处理docker相关操作
def docker(a1, a2, a3, a4):
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
        k8sget(a2, "pod")
    elif(a1 == "kl"):
        podLog(a2, a3, a4);
    elif(a1 == "ke"):
        podBash(a2, a3);
    elif(a1 == "ks"):
        k8sget(a2, "service")
    elif(a1 == "ki"):
        k8sget(a2, "ingress")
    elif(a1 == "kd"):
        k8sget(a2, "deploy")
    elif(a1 == "kdp"):
        k8sDesc(a3,"pod", a2);
    elif(a1 == "kdi"):
        k8sDesc(a3,"ingress", a2);
    elif(a1 == "kds"):
        k8sDesc(a3,"service", a2);
    elif(a1 == "kdd"):
        k8sDesc(a3,"deployments", a2);
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
        namespace 命名空间以!结尾支持模糊
      3.简化进入pod bash 命令，使用
        xx ke [pod] [namespace]
        登入pod bash，pod pod名称支持模糊搜索，namespace 所属命名空间以!结尾支持模糊
      4.简化查询pod 日志命令，使用
        xx kl [pod] [namespace] [lines]
        查询pod日志，pod名称支持模糊搜索，namespace 所属命名空间以!结尾支持模糊, lines 输出行数，默认10行
      5.简化查询deployments 命令，使用
        xx kd [namespace]
        namespace 命名空间以!结尾支持模糊
      6.简化查询ingress 命令，使用
        xx ki [namespace]
        namespace 命名空间以!结尾支持模糊
      7.简化查询service 命令，使用
        xx ks [namespace]
        namespace 命名空间以!结尾支持模糊
      8.简化查询pod describe命令，使用
        xx kdp [pod] [namespace]
        pod名称支持模糊搜索 namespace 命名空间以!结尾支持模糊
      9.简化查询ingress describe命令，使用
        xx kdi [ingress] [namespace]
        ingress名称支持模糊搜索 namespace 命名空间以!结尾支持模糊
      10.简化查询service describe命令，使用
        xx kds [service] [namespace]
        service名称支持模糊搜索 namespace 命名空间以!结尾支持模糊
      11.简化查询deployments describe命令，使用
        xx kdd [deployments] [namespace]
        service名称支持模糊搜索 namespace 命名空间以!结尾支持模糊
    """ % version
    return

def versions( version ):
    print("xx v%s" % version)

#处理k8s查看namespace
def k8sNamespace():
    print("查看namespace，命令：kubectl get namespace")
    os.system("kubectl get namespace")

# 查找pod
def findPod(pod = "", namespace = ""):
    return findX("pod", pod, namespace)

# 查找namespace
def findNameSpace(namespace = ""):
    return findX("namespace", namespace, "")

# 查找x  t:类别  x:名称（支持模糊） namespace 命名空间
def findX(t ,x = "", namespace = ""):
    dsprint = os.popen("kubectl get %s %s | grep \"%s\"" % (t, namespace, x))
    res = dsprint.readlines()
    fixres = []
    xname = ""
    for r in res:
        rs = r.split()
        if rs[0].find(x) >=0 :
            fixres.append(rs[0])
    if len(fixres) == 0:
        print("未找到%s，请检查输入" % (t))
    elif len(fixres) == 1:
        xname = fixres[0]
    elif len(fixres) > 1:
        for i in range(0, len(fixres)):
            print("%d:%s" % (i+1,fixres[i]))
        m = int(raw_input("请选择%s："%(t)))
        if m > 0 and m <= len(fixres) :
            xname = fixres[m-1]
    return xname


# 处理namespace
def fixNS(namespace):
    if namespace is None or namespace == "":
        namespace = ""
    elif namespace.find("!")<0 and namespace[0:3] != " -n":
        namespace = " -n " + namespace
    elif namespace.find("!")>=0:
        namespace = findNameSpace(namespace[0:-1])
        namespace = " -n " + namespace
    return namespace

#处理k8s pod 日志函数
# pod pod名称，支持模糊
# namespace 命名空间
# lines 显示log行数，默认10
def podLog(pod = "", namespace = "",lines = 10):
    if lines is None:
        lines = 10
    namespace = fixNS(namespace)
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
    namespace = fixNS(namespace)
    podname = findPod(pod, namespace)
    if podname != "":
        command = "kubectl exec -it %s %s -- bash" % (podname, namespace)
        print("进入%s bash，命令：%s" % (pod, command))
        os.system(command)
    return


def k8sget(namespace, cmd):
    k8sCmd(namespace, cmd, "get", "")

def k8sDesc(namespace, cmd, x):
    ns = fixNS(namespace)
    x = findX(cmd, x, ns);
    k8sCmd(ns, cmd, "describe", x)

#执行方法  修复namespace
def k8sCmd(namespace, cmd, t, pod):
    namespace = fixNS(namespace)
    command = "kubectl %s %s %s %s" % (t, cmd, pod, namespace)
    print("获取%s 下的%s，命令：%s" % (namespace, cmd, command))
    os.system(command)
    return

#主方法执行
version = "0.6" # 版本号
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
