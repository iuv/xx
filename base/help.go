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
  1. Query the namespace command, use
    xx kn [keyword]
    or
    xx kns [keyword]
    or
    xx knamespace [keyword]
    keyword supports fuzzy search

  2. To query the pod command, use
    xx kp [keyword] [namespace]
    or
    xx kpod [keyword] [namespace]
    keyword Fuzzy matching pod, if you want to query all namespace namespaces support fuzzy matching

  3. Enter the pod bash command and use
    xx ke [pod] [namespace] [sh]
    or
    xx kexe [pod] [namespace] [sh]
    Log in to pod bash, the pod pod name supports fuzzy search, the namespace to which the namespace belongs supports fuzzy search, sh defaults to bash, and there are special ones that can be passed in (under /bin/ directory)

  4. To query the pod log command, use
    xx kl [pod] [namespace] [lines]
    or
    xx klog [pod] [namespace] [lines]
    Query the pod log, the pod name supports fuzzy search, the namespace to which the namespace belongs supports fuzzy search, lines is the number of output lines, the default is 100 lines

  5. Query the deployments command, use
    xx kd [deployment] [namespace]
    or
    xx kdeployment [deployment] [namespace]
    Deployment name supports ambiguity, namespace namespace supports ambiguity

  6. Query the ingress command, use
    xx ki [ingress] [namespace]
    or
    xx kingress [ingress] [namespace]
    ingress name supports ambiguity, namespace namespace supports ambiguity

  7. Query the service command, use
    xx ks [service] [namespace]
    or
    xx kservice [service] [namespace]
    Service name supports ambiguity, namespace namespace supports ambiguity

  8. Query the configmap command, use
    xx kc [configmap] [namespace]
    or
    xx kconfigmap [configmap] [namespace]
    configmap names support ambiguity, namespace namespaces support ambiguity

  9. To query the secret command, use
    xx ksec [secret] [namespace]
    or
    xx ksecret [secret] [namespace]
    The secret name supports ambiguity, and the namespace namespace supports ambiguity

  10. To query the statefulset command, use
    xx kss [statefulset] [namespace]
    or
    xx kstatefulset [statefulset] [namespace]
    Statefulset names support ambiguity, namespace namespaces support ambiguity

  11. To query the pod describe command, use
    xx kpd [pod] [namespace]
    or
    xx kpodd [pod] [namespace]
    pod name supports fuzzy search, namespace namespace supports fuzzy search

  12. To query the ingress describe command, use
    xx kid [ingress] [namespace]
    or
    xx kingressd [ingress] [namespace]
    ingress name supports fuzzy search, namespace namespace supports fuzzy search

  13. To query the service describe command, use
    xx ksd [service] [namespace]
    or
    xx kserviced [service] [namespace]
    Service name supports fuzzy search, namespace

  14. To query the deployment describe command, use
    xx kdd [deployment] [namespace]
    or
    xx kdeploymentd [deployment] [namespace]
    The deployment name supports fuzzy search, and the namespace namespace supports fuzzy search

  15. To query the configmap describe command, use
    xx kcd [configmap] [namespace]
    or
    xx kconfigmapd [configmap] [namespace]
    configmap name supports fuzzy search, namespace namespace supports fuzzy

  16. To query the secret describe command, use
    xx ksecd [secret] [namespace]
    or
    xx ksecretd [secret] [namespace]
    The secret name supports fuzzy search, and the namespace namespace supports fuzzy search

  17. To query the statefulset describe command, use
    xx kssd [statefulset] [namespace]
    or
    xx kstatefulsetd [statefulset] [namespace]
    The statefulset name supports fuzzy search, and the namespace namespace supports fuzzy search

  18. Save the pod yaml command and use
    xx kpy [pod] [namespace] [file]
    or
    xx kpody [pod] [namespace] [file]
    Pod name supports fuzzy search, namespace namespace supports fuzzy, file is saved to file name

  19. Save the ingress yaml command and use
    xx kiy [ingress] [namespace] [file]
    or
    xx kingressy [ingress] [namespace] [file]
    ingress name supports fuzzy search, namespace namespace supports fuzzy, file is saved to file name

  20. Save the service describe command and use
    xx ksy [service] [namespace] [file]
    or
    xx kservicey [service] [namespace] [file]
    The service name supports fuzzy search, the namespace namespace supports fuzzy search, and the file is saved to the file name

  21. Save the deployment yaml command and use
    xx kdy [deployment] [namespace] [file]
    or
    xx kdeploymenty [deployment] [namespace] [file]
    The deployment name supports fuzzy search, the namespace namespace supports fuzzy search, and the file is saved to the file name

  22. Save the configmap yaml command and use
    xx kcy [configmap] [namespace] [file]
    or
    xx kconfigmapy [configmap] [namespace] [file]
    configmap name supports fuzzy search, namespace namespace supports fuzzy, file is saved to file name

  23. Save the secret yaml command and use
    xx ksecy [secret] [namespace] [file]
    or
    xx ksecrety [secret] [namespace] [file]
    The secret name supports fuzzy search, the namespace namespace supports fuzzy search, and the file is saved to the file name

  24. Save the statefulset yaml command and use
    xx kssy [statefulset] [namespace] [file]
    or
    xx kstatefulsety [statefulset] [namespace] [file]
    Statefulset name supports fuzzy search, namespace namespace supports fuzzy, file is saved to filename

  25. To delete the pod command, use
    xx kpdel [pod] [namespace]
    or
    xx kpoddel [pod] [namespace]
    pod name supports fuzzy search, namespace namespace supports fuzzy search

  26, delete the ingress command, use
    xx kidel [ingress] [namespace]
    or
    xx kingressdel [ingress] [namespace]
    ingress name supports fuzzy search, namespace namespace supports fuzzy search

  27, delete the service command, use
    xx ksdel [service] [namespace]
    or
    xx kservicedel [service] [namespace]
    The service name supports fuzzy search, and the namespace namespace supports fuzzy search

  28, delete the deployment command, use
    xx kddel [deployment] [namespace]
    or
    xx kdeploymentdel [deployment] [namespace]
    The deployment name supports fuzzy search, and the namespace namespace supports fuzzy search

  29, delete the configmap command, use
    xx kcdel [configmap] [namespace]
    or
    xx kconfigmapdel [configmap] [namespace]
    configmap name supports fuzzy search, namespace namespace supports fuzzy

  30. Delete the secret command and use
    xx ksecdel [secret] [namespace]
    or
    xx ksecretdel [secret] [namespace]
    The secret name supports fuzzy search, and the namespace namespace supports fuzzy search

  31. Delete the statefulset command and use
    xx kssdel [statefulset] [namespace]
    or
    xx kstatefulsetdel [statefulset] [namespace]
    The statefulset name supports fuzzy search, and the namespace namespace supports fuzzy search

  32. Apply the yaml configuration file command, use
    xx ka [file]
    or
    xx kapply [file]
    file yaml configuration file

  33. To copy the file command from the pod container, use
    xx kcopy [pod] [namespace] [srcFile] [saveFile]
    Pod name supports fuzzy search, namespace namespace supports fuzzy search, srcFile is the path of the file to be copied in the container, saveFile is the local save path
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
