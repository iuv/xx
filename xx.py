#!/usr/bin/python
#-*- coding:utf-8 -*-
import sys,os

#处理docker相关操作
def docker(a1, a2, a3):
	if(a1 == "dl"):
		dockerLog(a2, a3);
	elif(a1 == "de"):
		dockerExec(a2);
	return

# 查找docker窗口
def findDocker(name=" "):
	dsprint = os.popen("docker ps | grep \"%s\"" % name)
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

#处理docker 日志函数
# name 容器名称，支持模糊
# lines 显示log行数，默认10
def dockerLog(name = " ",lines = 10):
	dockername = findDocker(name)
	print("查看%s日志：" % dockername)
	os.system("docker logs --tail 10 -f %s" % dockername)
	return

#进入docker bash
# name 容器名称，支持模糊
def dockerExec(name = " "):
	dockername = findDocker(name)
	print("登入 %s bash：" % dockername)
	os.system("docker exec -it  %s bash" % dockername)
	return

#主方法执行
a1 = sys.argv[1]
a2 = sys.argv[2]
a3 = None
if(len(sys.argv) == 4):
	a3 = sys.argv[3]
if(str(a1).startswith("d")):
	docker(a1,a2,a3)
