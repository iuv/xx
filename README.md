[中文说明](https://github.com/iuv/xx/blob/master/README_ZH.md)

# xx simplified command tool v1.4.0
![License](https://img.shields.io/badge/license-MIT-4EB1BA)
[![wiki](https://img.shields.io/badge/Document-Wiki-green)](http://blog.jisuye.com/xx)

This tool simplifies the use of common shell, docker, and kubernetes commands

The 0.X version is written in python, and the running environment needs to support py and the sh script to support running. In order to solve the running environment dependency and maintain single-file execution

1.X and later versions will be written in go language

If you need to compile it yourself, you can clone this repository and run the build.sh script
## demo
![demo](https://raw.githubusercontent.com/iuv/xx/master/demo.gif?sanitize=true)
## installation method
1. Mac use `wget https://raw.githubusercontent.com/iuv/xx/master/build/mac/x86/xx` to download xx file(x86/arm64)
2. linux uses `wget https://raw.githubusercontent.com/iuv/xx/master/build/linux/x86/xx` to download xx file(x86/arm/arm64)
3. Execute `chmod +x xx; ./xx install` to install
4. You can use the `xx` command
5. Update using `xx update` command
6. Use `xx h` for help, `xx zh` for Chinese help

## Use help as follows:

![help](https://raw.githubusercontent.com/iuv/xx/master/xx.svg?sanitize=true)

### shell:
1. xx ip [port] Get local ip and public network ip (if there is an external network)
````
 Example result:
 Local IP: 172.16.112.12
 HTTP Server: http://172.16.112.12
 HTTP Server: http://172.16.112.12:8080
 Public Network IP: 8.8.8.8
 optional output
 HTTP Server: http://172.16.112.12:[port]
````
2. xx ps [str] Get the process, fuzzy search according to str, and highlight

### docker (parameters are empty and use "@" placeholder when subsequent parameters are required):
1. Run the docker command, use
```shell
xx dr [image] [container] [port]
or
xx drun [image] [container] [port]
````
By default, the background process is used to start the docker image name supports fuzzy search, container sets the container name, port is the port number of the mapping
Support "8080:8080" and "8080" two methods, "8080" will be automatically completed as "8080:8080"

2. To query the docker container log command, use

```shell
xx dl [container] [lines]
or
xx dlog [container] [lines]
````
Query container output log, container supports image/container name fuzzy search, lines is the number of output lines, default 100 lines

3. Enter the docker bash command and use
```shell
xx de [container]
or
xx dexec [container]
````
Enter the container bash, container supports image/container name fuzzy search

4. Start the docker container command, use
```shell
xx ds [container]
or
xx dstart [container]
````
Start the container, container supports image/container name fuzzy search

5. Restart the docker command and use
```shell
xx drs [container]
or
xx drestart [container]
````
Restart the container, container supports image/container name fuzzy search

6. Stop the container command and use
```shell
xx dk [container]
or
xx dstop [container]
````
Stop the container, container supports image/container name fuzzy search

7. Find the docker image and use
```shell
xx di [image]
or
xx dimages [image]
````
Find images, image supports fuzzy search

8. Pull the docker image and use
```shell
xx dpl [imageFullPath]
or
xx dpull [imageFullPath]
````
Pull image, image full path

9. Push the docker image and use
```shell
xx dph [image]
or
xx dpush [image]
````
Push image, image supports fuzzy search

10. Set tag of the docker image and use
```shell
xx dt [image] [tag]
or
xx dtag [image] [tag]
````
Image tagging, image supports fuzzy search, tag is the name of the tag that needs to be tagged

11, docker view all containers, use
```shell
xx dps [container]
````
View all containers (running and stopped), container supports image/container name fuzzy search

12. Docker deletes the image and the container started using the image, using
```shell
xx drm [image]
````
Delete the image and all containers started using the image, image supports fuzzy search

13. Docker local-containers copy files to each other, use
```shell
# Copy the files in the container to the local
xx dc [container]:[filePath] [localPath]
# Copy local files to the container
xx dc [localPath] [container]:[filePath]
or
# Copy the files in the container to the local
xx dcp [container]:[filePath] [localPath]
# Copy local files to the container
xx dcp [localPath] [container]:[filePath]
````
Docker local-containers copy files to each other, container name supports fuzzy search filePath container file/folder path localPath local file path eg: xx dc mysql:/tmp/a.sql .

14. Docker saves the image as a local file and uses
```shell
xx dsa [image] [fileName]
or
xx dsave [image] [fileName]
````
Docker saves the image as a local file, image supports fuzzy search for the file name saved by fileName

15. Docker imports images from local files, using
```shell
xx dlo [fileName]
or
xx dload [fileName]
````
Docker imports images from local files, fileName is the file name to be imported

16. Docker saves the running container as an image, using
```shell
xx dco [container] [image]
or
xx dcommit [container] [image]
````
Docker saves the running container as an image, container name supports fuzzy search image saved image name

17, docker view the image creation history, use
```shell
xx dh [image]
or
xx dhistory [image]
````
docker view the image creation history, image image name supports fuzzy search

18. Docker builds the image (execute in the directory where the Dockerfile is located)
```shell
xx db [image]
or
xx dbuild [image]
````
Docker builds an image and executes it in the directory where the Dockerfile is located. image is the image name

### k8s ("@" is used when the parameter is empty and subsequent parameters are required):
#### 1.exec
1.1 Enter the pod bash command and use
```shell
xx ke [pod] [namespace] [sh]
or
xx kexe [pod] [namespace] [sh]
````
Log in to pod bash, the pod pod name supports fuzzy search, the namespace to which the namespace belongs supports fuzzy search, sh defaults to bash, and there are special ones that can be passed in (under /bin/ directory)

#### 2.log
2.1 To query the pod log command, use
```shell
xx kl [pod] [namespace] [lines]
or
xx klog [pod] [namespace] [lines]
````
Query the pod log, the pod name supports fuzzy search, the namespace to which the namespace belongs supports fuzzy search, lines is the number of output lines, the default is 100 lines

#### 3.query
3.1 Query the namespace command, use
```shell
xx kn [namespace]
or
xx kns [namespace]
or
xx knamespace [namespace]
````
namespace supports fuzzy search

3.2 To query the pod command, use
```shell
xx kp [pod] [namespace]
or
xx kpod [pod] [namespace]
````
pod Fuzzy matching pod, if you want to query all namespace namespaces support fuzzy matching

3.3 Query the deployments command, use
```shell
xx kd [deployment] [namespace]
or
xx kdeployment [deployment] [namespace]
````
Deployment name supports ambiguity, namespace namespace supports ambiguity 

3.4 Query the ingress command, use
```shell
xx ki [ingress] [namespace]
or
xx kingress [ingress] [namespace]
````
ingress name supports ambiguity, namespace namespace supports ambiguity

3.5 Query the service command, use
```shell
xx ks [service] [namespace]
or
xx kservice [service] [namespace]
````
Service name supports ambiguity, namespace namespace supports ambiguity

3.6 Query the configmap command, use
```shell
xx kc [configmap] [namespace]
or
xx kconfigmap [configmap] [namespace]
````
configmap names support ambiguity, namespace namespaces support ambiguity

3.7 To query the secret command, use
```shell
xx ksec [secret] [namespace]
or
xx ksecret [secret] [namespace]
````
The secret name supports ambiguity, and the namespace namespace supports ambiguity

3.8 To query the statefulset command, use
```shell
xx kss [statefulset] [namespace]
or
xx kstatefulset [statefulset] [namespace]
````
Statefulset names support ambiguity, namespace namespaces support ambiguity

3.9 To query the CR command, use
```shell
xx kcr [cr] [cr key] [namespace]
```
cr type ,cr key support ambiguity, namespace namespaces support ambiguity

#### 4.describe
4.1 To query the pod describe command, use
```shell
xx kpd [pod] [namespace] [key]
or
xx kpodd [pod] [namespace] [key]
````
pod name supports fuzzy search, namespace namespace supports fuzzy search, key use by grep content

4.2 To query the ingress describe command, use
```shell
xx kid [ingress] [namespace] [key]
or
xx kingressd [ingress] [namespace] [key]
````
ingress name supports fuzzy search, namespace namespace supports fuzzy search, key use by grep content

4.3 To query the service describe command, use
```shell
xx ksd [service] [namespace] [key]
or
xx kserviced [service] [namespace] [key]
````
The service name supports fuzzy search, and the namespace namespace supports fuzzy search, key use by grep content

4.4 To query the deployment describe command, use
```shell
xx kdd [deployment] [namespace] [key]
or
xx kdeploymentd [deployment] [namespace] [key]
````
The deployment name supports fuzzy search, and the namespace namespace supports fuzzy search, key use by grep content

4.5 To query the configmap describe command, use
```shell
xx kcd [configmap] [namespace] [key]
or
xx kconfigmapd [configmap] [namespace] [key]
````
configmap name supports fuzzy search, namespace namespace supports fuzzy search, key use by grep content

4.6 To query the secret describe command, use
```shell
xx ksecd [secret] [namespace] [key]
or
xx ksecretd [secret] [namespace] [key]
````
The secret name supports fuzzy search, and the namespace namespace supports fuzzy search, key use by grep content

4.7 To query the statefulset describe command, use
```shell
xx kssd [statefulset] [namespace] [key]
or
xx kstatefulsetd [statefulset] [namespace] [key]
````
The statefulset name supports fuzzy search, and the namespace namespace supports fuzzy search, key use by grep content

4.8 To query the CR describe command, use
```shell
xx kcrd [cr] [cr key] [namespace] [key]
```
cr type ,cr key supports fuzzy search, and the namespace namespace supports fuzzy search, key use by grep content

#### 5.yaml
5.1 Save the pod yaml command and use
```shell
xx kpy [pod] [namespace] [file]
or
xx kpody [pod] [namespace] [file]
````
Pod name supports fuzzy search, namespace namespace supports fuzzy, file is saved to file name

5.2 Save the ingress yaml command and use
```shell
xx kiy [ingress] [namespace] [file]
or
xx kingressy [ingress] [namespace] [file]
````
ingress name supports fuzzy search, namespace namespace supports fuzzy, file is saved to file name

5.3 Save the service yaml command and use
```shell
xx ksy [service] [namespace] [file]
or
xx kservicey [service] [namespace] [file]
````
The service name supports fuzzy search, the namespace namespace supports fuzzy search, and the file is saved to the file name

5.4 Save the deployment yaml command and use
```shell
xx kdy [deployment] [namespace] [file]
or
xx kdeploymenty [deployment] [namespace] [file]
````
The deployment name supports fuzzy search, the namespace namespace supports fuzzy search, and the file is saved to the file name

5.5 Save the configmap yaml command and use
```shell
xx kcy [configmap] [namespace] [file]
or
xx kconfigmapy [configmap] [namespace] [file]
````
configmap name supports fuzzy search, namespace namespace supports fuzzy, file is saved to file name

5.6 Save the secret yaml command and use
```shell
xx ksecy [secret] [namespace] [file]
or
xx ksecrety [secret] [namespace] [file]
````
The secret name supports fuzzy search, the namespace namespace supports fuzzy search, and the file is saved to the file name

5.7 Save the statefulset yaml command and use
```shell
xx kssy [statefulset] [namespace] [file]
or
xx kstatefulsety [statefulset] [namespace] [file]
````
Statefulset name supports fuzzy search, namespace namespace supports fuzzy, file is saved to filename

5.8 Save the cr yaml command and use
```shell
xx kcry [cr] [cr key] [namespace] [file]
```
cr type, cr key supports fuzzy search, namespace namespace supports fuzzy, file is saved to filename

#### 6.delete
6.1 Delete the pod command, use
```shell
xx kpdel [pod] [namespace]
or
xx kpoddel [pod] [namespace]
````
pod name supports fuzzy search, namespace namespace supports fuzzy search

6.2 Delete the ingress command, use
```shell
xx kidel [ingress] [namespace]
or
xx kingressdel [ingress] [namespace]
````
ingress name supports fuzzy search, namespace namespace supports fuzzy search

6.3 Delete the service command, use
```shell
xx ksdel [service] [namespace]
or
xx kservicedel [service] [namespace]
````
The service name supports fuzzy search, and the namespace namespace supports fuzzy search

6.4 delete the deployment command, use
```shell
xx kddel [deployment] [namespace]
or
xx kdeploymentdel [deployment] [namespace]
````
The deployment name supports fuzzy search, and the namespace namespace supports fuzzy search

6.5 delete the configmap command, use
```shell
xx kcdel [configmap] [namespace]
or
xx kconfigmapdel [configmap] [namespace]
````
configmap name supports fuzzy search, namespace namespace supports fuzzy

6.6 Delete the secret command and use
```shell
xx ksecdel [secret] [namespace]
or
xx ksecretdel [secret] [namespace]
````
The secret name supports fuzzy search, and the namespace namespace supports fuzzy search

6.7 Delete the statefulset command and use
```shell
xx kssdel [statefulset] [namespace]
or
xx kstatefulsetdel [statefulset] [namespace]
````
The statefulset name supports fuzzy search, and the namespace namespace supports fuzzy search

6.8 Delete the cr command and use
```shell
xx kcrdel [cr] [cr key] [namespace]
```
cr type, cr key supports fuzzy search, and the namespace namespace supports fuzzy search

#### 7.apply
7.1 Apply the yaml configuration file command, use
```shell
xx ka [file]
or
xx kapply [file]
````
file yaml configuration file

#### 8.copy
8.1 To copy the file command from the pod container, use
```shell
xx kcopy [pod] [namespace] [srcFile] [saveFile]
````
Pod name supports fuzzy search, namespace namespace supports fuzzy search, srcFile is the path of the file to be copied in the container, saveFile is the local save path

#### 9.edit
9.1 To edit the pod command, use
```shell
xx kpe [pod] [namespace]
or
xx kpode [pod] [namespace]
```
pod name supports fuzzy search, namespace namespace supports fuzzy search

9.2 edit the ingress command, use
```shell
xx kie [ingress] [namespace]
or
xx kingresse [ingress] [namespace]
```
ingress name supports fuzzy search, namespace namespace supports fuzzy search

9.3 edit the service command, use
```shell
xx kse [service] [namespace]
or
xx kservicee [service] [namespace]
```
The service name supports fuzzy search, and the namespace namespace supports fuzzy search

9.4 edit the deployment command, use
```shell
xx kde [deployment] [namespace]
or
xx kdeploymente [deployment] [namespace]
```
The deployment name supports fuzzy search, and the namespace namespace supports fuzzy search

9.5 edit the configmap command, use
```shell
xx kce [configmap] [namespace]
or
xx kconfigmape [configmap] [namespace]
```
configmap name supports fuzzy search, namespace namespace supports fuzzy

9.6 edit the secret command and use
```shell
xx ksece [secret] [namespace]
or
xx ksecrete [secret] [namespace]
```
The secret name supports fuzzy search, and the namespace namespace supports fuzzy search

9.7 Edit the statefulset command and use
```shell
xx ksse [statefulset] [namespace]
or
xx kstatefulsete [statefulset] [namespace]
```
The statefulset name supports fuzzy search, and the namespace namespace supports fuzzy search

9.8 Edit the cr command and use
```shell
xx kcre [cr] [cr key] [namespace]
```
cr type, cr key supports fuzzy search, and the namespace namespace supports fuzzy search
