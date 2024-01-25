package base

import (
	"fmt"
)

var title string = " xx Simplified Command Tool " + Version
var describe = `
This tool simplifies the use of common shell, docker, and kubernetes commands
If you have usage problems, please visit https://github.com/iuv/xx to submit issues,
Use help as follows:
`
var shell string = `
shell:
  1. xx ip [port] Get local ip and public network ip (if there is an external network)
  Example result:
  Local IP: 172.16.112.12
  HTTP Server: http://172.16.112.12
  HTTP Server: http://172.16.112.12:8080
  Public Network IP: 8.8.8.8
  optional output
  HTTP Server: http://172.16.112.12:[port]

  2. xx ps [str] Get the process, fuzzy search according to str, and highlight
`
var docker string = `
docker (the "@" placeholder is used when the parameter is empty and subsequent parameters are required):
  1. Run the docker command, use
    xx dr [imageName] [containerName] [port]
    or
    xx drun [imageName] [containerName] [port]
    By default, the background process is used to start the imageName image name supports fuzzy search, containerName sets the container name, port is the port number of the mapping
    Support "8080:8080" and "8080" two methods, "8080" will be automatically completed as "8080:8080"

  2. To query the docker container log command, use
    xx dl [dockername] [lines]
    or
    xx dlog [dockername] [lines]
    Query container output log, dockername supports image/container name fuzzy search, lines is the number of output lines, default 100 lines

  3. Enter the docker bash command and use
    xx de [dockername]
    or
    xx dexec [dockername]
    Enter the container bash, dockername supports image/container name fuzzy search

  4. Start the docker container command, use
    xx ds [dockername]
    or
    xx dstart [dockername]
    Start the container, dockername supports image/container name fuzzy search

  5. Restart the docker command and use
    xx drs [dockername]
    or
    xx drestart [dockername]
    Restart the container, dockername supports image/container name fuzzy search

  6. Stop the docker command and use
    xx dk [dockername]
    or
    xx dstop [dockername]
    Stop the container, dockername supports image/container name fuzzy search

  7. Find the docker image and use
    xx di [imageName]
    or
    xx dimages [imageName]
    Find images, imageName supports fuzzy search

  8. Pull the docker image and use
    xx dpl [imageName]
    or
    xx dpull [imageName]
    Pull image, imageName, image full path

  9. Push the docker image and use
    xx dph [imageName]
    or
    xx dpush [imageName]
    Push image, imageName supports fuzzy search

  10. Tag the docker image and use
    xx dt [imageName] [tagname]
    or
    xx dtag [imageName] [tagname]
    Image tagging, imageName supports fuzzy search, tagname is the name of the tag that needs to be tagged

  11, docker view all containers, use
    xx dps [dockername]
    View all containers (running and stopped), dockername supports image/container name fuzzy search

  12. Docker deletes the image and the container started using the image, using
    xx drm [imageName]
    Delete the image and all containers started using the image, imageName supports fuzzy search

  13. Docker local-containers copy files to each other, use
    # Copy the files in the container to the local
    xx dc [dockerName]:[filePath] [localPath]
    # Copy local files to the container
    xx dc [localPath] [dockerName]:[filePath]
    or
    # Copy the files in the container to the local
    xx dcp [dockerName]:[filePath] [localPath]
    # Copy local files to the container
    xx dcp [localPath] [dockerName]:[filePath]
    Docker local-containers copy files to each other, dockerName container name supports fuzzy search filePath container file/folder path localPath local file path eg: xx dc mysql:/tmp/a.sql .
    
  14. Docker saves the image as a local file and uses
    xx dsa [imageName] [fileName]
    or
    xx dsave [imageName] [fileName]
    Docker saves the image as a local file, imageName supports fuzzy search for the file name saved by fileName
    
  15. Docker imports images from local files, using
    xx dlo [fileName]
    or
    xx dload [fileName]
    Docker imports images from local files, fileName is the file name to be imported
    
  16. Docker saves the running container as an image, using
    xx dco [dockerName] [imageName]
    or
    xx dcommit [dockerName] [imageName]
    Docker saves the running container as an image, dockerName container name supports fuzzy search imageName saved image name
    
  17, docker view the image creation history, use
    xx dh [imageName]
    or
    xx dhistory [imageName]
    docker view the image creation history, imageName image name supports fuzzy search
  18. Docker builds the image (execute in the directory where the Dockerfile is located)
xx db [imageName]
	or
xx dbuild [imageName]
    Docker builds an image and executes it in the directory where the Dockerfile is located. imageName is the image name
`
var k8s = `
k8s (the "@" placeholder is used when the parameter is empty and subsequent parameters are required):
===================1.exec========================
  1.1 Enter the pod bash command and use
    xx ke [pod] [namespace] [sh]
    or
    xx kexe [pod] [namespace] [sh]
    Log in to pod bash, the pod pod name supports fuzzy search, the namespace to which the namespace belongs supports fuzzy search, sh defaults to bash, and there are special ones that can be passed in (under /bin/ directory)

===================2.log========================
  2.1 To query the pod log command, use
    xx kl [pod] [namespace] [lines]
    or
    xx klog [pod] [namespace] [lines]
    Query the pod log, the pod name supports fuzzy search, the namespace to which the namespace belongs supports fuzzy search, lines is the number of output lines, the default is 100 lines

===================3.query========================
  3.1 Query the namespace command, use
    xx kn [keyword]
    or
    xx kns [keyword]
    or
    xx knamespace [keyword]
    keyword supports fuzzy search

  3.2 To query the pod command, use
    xx kp [keyword] [namespace]
    or
    xx kpod [keyword] [namespace]
    keyword Fuzzy matching pod, if you want to query all namespace namespaces support fuzzy matching

  3.3 Query the deployments command, use
    xx kd [deployment] [namespace]
    or
    xx kdeployment [deployment] [namespace]
    Deployment name supports ambiguity, namespace namespace supports ambiguity

  3.4 Query the ingress command, use
    xx ki [ingress] [namespace]
    or
    xx kingress [ingress] [namespace]
    ingress name supports ambiguity, namespace namespace supports ambiguity

  3.5 Query the service command, use
    xx ks [service] [namespace]
    or
    xx kservice [service] [namespace]
    Service name supports ambiguity, namespace namespace supports ambiguity

  3.6 Query the configmap command, use
    xx kc [configmap] [namespace]
    or
    xx kconfigmap [configmap] [namespace]
    configmap names support ambiguity, namespace namespaces support ambiguity

  3.7 To query the secret command, use
    xx ksec [secret] [namespace]
    or
    xx ksecret [secret] [namespace]
    The secret name supports ambiguity, and the namespace namespace supports ambiguity

  3.8 To query the statefulset command, use
    xx kss [statefulset] [namespace]
    or
    xx kstatefulset [statefulset] [namespace]
    Statefulset names support ambiguity, namespace namespaces support ambiguity

  3.9 To query the CR command, use
    xx kcr [cr] [cr key] [namespace]
    cr type ,cr key support ambiguity, namespace namespaces support ambiguity

===================4.describe========================
  4.1 To query the pod describe command, use
    xx kpd [pod] [namespace] [key]
    or
    xx kpodd [pod] [namespace] [key]
    pod name supports fuzzy search, namespace namespace supports fuzzy search, key use by grep content

  4.2 To query the ingress describe command, use
    xx kid [ingress] [namespace] [key]
    or
    xx kingressd [ingress] [namespace] [key]
    ingress name supports fuzzy search, namespace namespace supports fuzzy search, key use by grep content

  4.3 To query the service describe command, use
    xx ksd [service] [namespace] [key]
    or
    xx kserviced [service] [namespace] [key]
    Service name supports fuzzy search, namespace namespace supports fuzzy search, key use by grep content

  4.4 To query the deployment describe command, use
    xx kdd [deployment] [namespace] [key]
    or
    xx kdeploymentd [deployment] [namespace] [key
    The deployment name supports fuzzy search, and the namespace namespace supports fuzzy search, key use by grep content

  4.5 To query the configmap describe command, use
    xx kcd [configmap] [namespace] [key]
    or
    xx kconfigmapd [configmap] [namespace] [key]
    configmap name supports fuzzy search, namespace namespace supports fuzzy, key use by grep content

  4.6 To query the secret describe command, use
    xx ksecd [secret] [namespace] [key]
    or
    xx ksecretd [secret] [namespace] [key]
    The secret name supports fuzzy search, and the namespace namespace supports fuzzy search, key use by grep content

  4.7 To query the statefulset describe command, use
    xx kssd [statefulset] [namespace] [key]
    or
    xx kstatefulsetd [statefulset] [namespace] [key]
    The statefulset name supports fuzzy search, and the namespace namespace supports fuzzy search, key use by grep content

  4.8 To query the CR describe command, use
    xx kcrd [cr] [cr key] [namespace] [key]
    cr type ,cr key supports fuzzy search, and the namespace namespace supports fuzzy search, key use by grep content

===================5.yaml========================
  5.1 Save the pod yaml command and use
    xx kpy [pod] [namespace] [file]
    or
    xx kpody [pod] [namespace] [file]
    Pod name supports fuzzy search, namespace namespace supports fuzzy, file is saved to file name

  5.2 Save the ingress yaml command and use
    xx kiy [ingress] [namespace] [file]
    or
    xx kingressy [ingress] [namespace] [file]
    ingress name supports fuzzy search, namespace namespace supports fuzzy, file is saved to file name

  5.3 Save the service describe command and use
    xx ksy [service] [namespace] [file]
    or
    xx kservicey [service] [namespace] [file]
    The service name supports fuzzy search, the namespace namespace supports fuzzy search, and the file is saved to the file name

  5.4 Save the deployment yaml command and use
    xx kdy [deployment] [namespace] [file]
    or
    xx kdeploymenty [deployment] [namespace] [file]
    The deployment name supports fuzzy search, the namespace namespace supports fuzzy search, and the file is saved to the file name

  5.5 Save the configmap yaml command and use
    xx kcy [configmap] [namespace] [file]
    or
    xx kconfigmapy [configmap] [namespace] [file]
    configmap name supports fuzzy search, namespace namespace supports fuzzy, file is saved to file name

  5.6 Save the secret yaml command and use
    xx ksecy [secret] [namespace] [file]
    or
    xx ksecrety [secret] [namespace] [file]
    The secret name supports fuzzy search, the namespace namespace supports fuzzy search, and the file is saved to the file name

  5.7 Save the statefulset yaml command and use
    xx kssy [statefulset] [namespace] [file]
    or
    xx kstatefulsety [statefulset] [namespace] [file]
    Statefulset name supports fuzzy search, namespace namespace supports fuzzy, file is saved to filename

  5.8 Save the cr yaml command and use
    xx kcry [cr] [cr key] [namespace] [file]
    cr type, cr key supports fuzzy search, namespace namespace supports fuzzy, file is saved to filename

===================6.delete========================
  6.1 To delete the pod command, use
    xx kpdel [pod] [namespace]
    or
    xx kpoddel [pod] [namespace]
    pod name supports fuzzy search, namespace namespace supports fuzzy search

  6.2 delete the ingress command, use
    xx kidel [ingress] [namespace]
    or
    xx kingressdel [ingress] [namespace]
    ingress name supports fuzzy search, namespace namespace supports fuzzy search

  6.3 delete the service command, use
    xx ksdel [service] [namespace]
    or
    xx kservicedel [service] [namespace]
    The service name supports fuzzy search, and the namespace namespace supports fuzzy search

  6.4 delete the deployment command, use
    xx kddel [deployment] [namespace]
    or
    xx kdeploymentdel [deployment] [namespace]
    The deployment name supports fuzzy search, and the namespace namespace supports fuzzy search

  6.5 delete the configmap command, use
    xx kcdel [configmap] [namespace]
    or
    xx kconfigmapdel [configmap] [namespace]
    configmap name supports fuzzy search, namespace namespace supports fuzzy

  6.6 Delete the secret command and use
    xx ksecdel [secret] [namespace]
    or
    xx ksecretdel [secret] [namespace]
    The secret name supports fuzzy search, and the namespace namespace supports fuzzy search

  6.7 Delete the statefulset command and use
    xx kssdel [statefulset] [namespace]
    or
    xx kstatefulsetdel [statefulset] [namespace]
    The statefulset name supports fuzzy search, and the namespace namespace supports fuzzy search

  6.8 Delete the cr command and use
    xx kcrdel [cr] [cr key] [namespace]
    cr type, cr key supports fuzzy search, and the namespace namespace supports fuzzy search

===================7.apply========================
  7.1 Apply the yaml configuration file command, use
    xx ka [file]
    or
    xx kapply [file]
    file yaml configuration file

===================8.copy========================
  8.1 To copy the file command from the pod container, use
    xx kcopy [pod] [namespace] [srcFile] [saveFile]
    Pod name supports fuzzy search, namespace namespace supports fuzzy search, srcFile is the path of the file to be copied in the container, saveFile is the local save path

===================9.edit========================
  9.1 To edit the pod command, use
    xx kpe [pod] [namespace]
    or
    xx kpode [pod] [namespace]
    pod name supports fuzzy search, namespace namespace supports fuzzy search

  9.2 edit the ingress command, use
    xx kie [ingress] [namespace]
    or
    xx kingresse [ingress] [namespace]
    ingress name supports fuzzy search, namespace namespace supports fuzzy search

  9.3 edit the service command, use
    xx kse [service] [namespace]
    or
    xx kservicee [service] [namespace]
    The service name supports fuzzy search, and the namespace namespace supports fuzzy search

  9.4 edit the deployment command, use
    xx kde [deployment] [namespace]
    or
    xx kdeploymente [deployment] [namespace]
    The deployment name supports fuzzy search, and the namespace namespace supports fuzzy search

  9.5 edit the configmap command, use
    xx kce [configmap] [namespace]
    or
    xx kconfigmape [configmap] [namespace]
    configmap name supports fuzzy search, namespace namespace supports fuzzy

  9.6 edit the secret command and use
    xx ksece [secret] [namespace]
    or
    xx ksecrete [secret] [namespace]
    The secret name supports fuzzy search, and the namespace namespace supports fuzzy search

  9.7 Edit the statefulset command and use
    xx ksse [statefulset] [namespace]
    or
    xx kstatefulsete [statefulset] [namespace]
    The statefulset name supports fuzzy search, and the namespace namespace supports fuzzy search

  9.8 Edit the cr command and use
    xx kcre [cr] [cr key] [namespace]
    cr type, cr key supports fuzzy search, and the namespace namespace supports fuzzy search

`

func Help(key string) {
	fmt.Println(title)
	fmt.Print(describe)
	switch key {
	case "shell":
		fmt.Print(shell)
	case "docker":
		fmt.Print(docker)
	case "k8s":
		fmt.Print(k8s)
	default:
		fmt.Print(shell)
		fmt.Print(docker)
		fmt.Print(k8s)
	}
}
