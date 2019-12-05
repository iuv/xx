package k8s

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"io/ioutil"
	"strconv"
	"strings"
	"xx/base"
)

func K8s(cmd,a1,a2,a3,a4 string){
	switch cmd {
	case "n": namespace(a1)
	case "p": pod(a1, a2)
	case "e": bash(a1, a2, a3)
	case "l": logs(a1, a2, a3)
	case "d": deployments(a1, a2)
	case "i": ingress(a1, a2)
	case "s": service(a1, a2)
	case "pd": podDesc(a1, a2)
	case "dd": deploymentDesc(a1, a2)
	case "id": ingressDesc(a1, a2)
	case "sd": serviceDesc(a1, a2)
	case "py": podYaml(a1, a2, a3)
	case "dy": deploymentYaml(a1, a2, a3)
	case "iy": ingressYaml(a1, a2, a3)
	case "sy": serviceYaml(a1, a2, a3)
	case "pdel": podDel(a1, a2)
	case "ddel": deploymentDel(a1, a2)
	case "idel": ingressDel(a1, a2)
	case "sdel": serviceDel(a1, a2)
	case "a": apply(a1)
	case "cp": cp(a1, a2, a3, a4)
	default: fmt.Println("command not found: k"+cmd)
	}
}
// k8s cp 复制文件
func cp(pod, namespace, srcFile, saveFile string)  {
	pod, namespace = getResAndNamespace("pod", pod, namespace)
	if pod == ""{
		return
	}
	ret := base.Execp("kubectl cp "+namespace+"/"+pod+":"+srcFile+" "+saveFile)
	fmt.Print(ret)
}
// apply yml 文件
func apply(file string)  {
	ret := base.Execp("kubectl apply -f "+file)
	fmt.Print(ret)
}
// service删除
func serviceDel(key, namespace string)  {
	resourceDel("service", key, namespace)
}
// ingress删除
func ingressDel(key, namespace string)  {
	resourceDel("ingress", key, namespace)
}
// deployment删除
func deploymentDel(key, namespace string)  {
	resourceDel("deployment", key, namespace)
}
// pod删除
func podDel(key, namespace string)  {
	resourceDel("pod", key, namespace)
}
// 删除资源基本方法
func resourceDel(res, key, namespace string)  {
	key, namespace = getResAndNamespace(res, key, namespace)
	if key != "" {
		prompt := &survey.Confirm{
			Message:  "Are you sure delete "+res+": "+key+" in namespace:"+namespace,
		}
		yesOrNo := false
		survey.AskOne(prompt, &yesOrNo)
		if yesOrNo{
			ret := base.Execp("kubectl delete "+res+" "+ key +" -n "+namespace)
			fmt.Print(ret)
		}
	}
}
// service yaml
func serviceYaml(key, namespace, file string)  {
	resourceYaml("service", key, namespace, file)
}
// ingress yaml
func ingressYaml(key, namespace, file string)  {
	resourceYaml("ingress", key, namespace, file)
}
// deployment yaml
func deploymentYaml(key, namespace, file string)  {
	resourceYaml("deployment", key, namespace, file)
}
// pod yaml
func podYaml(key, namespace, file string)  {
	resourceYaml("pod", key, namespace, file)
}
// 查询资源详情基本方法
func resourceYaml(res, key, namespace, file string)  {
	key, namespace = getResAndNamespace(res ,key, namespace)
	if key != "" {
		ret := base.Execp("kubectl get "+res+" "+key+" -n "+namespace+" -o yaml")
		if file == "" {
			fmt.Print(ret)
		} else {
			err := ioutil.WriteFile(file, []byte(ret), 0666)
			if err != nil{
				fmt.Print(err)
			}
		}
	}
}
// service详情
func serviceDesc(key, namespace string)  {
	resourceDescribe("service", key, namespace)
}
// ingress详情
func ingressDesc(key, namespace string)  {
	resourceDescribe("ingress", key, namespace)
}
// deployment详情
func deploymentDesc(key, namespace string)  {
	resourceDescribe("deployment", key, namespace)
}
// pod详情
func podDesc(key, namespace string)  {
	resourceDescribe("pod", key, namespace)
}
// 查询资源详情基本方法
func resourceDescribe(res, key, namespace string)  {
	key, namespace = getResAndNamespace(res, key, namespace)
	if key != "" {
		ret := base.Execp("kubectl describe "+res+" "+ key +" -n "+namespace)
		fmt.Print(ret)
	}
}

// 获取service
func service(key,namespace string)  {
	getResource("service", key, namespace)
}
// 获取ingress
func ingress(key,namespace string)  {
	getResource("ingress", key, namespace)
}

// 获取deployments
func deployments(key,namespace string)  {
	getResource("deployment", key, namespace)
}

// 查询控制台日志
func logs(pod, namespace, lines string)  {
	if lines == ""{
		lines = "100"
	}
	pod,namespace = getResAndNamespace("pod", pod, namespace)
	if pod != "" {
		base.Run("kubectl logs --tail "+lines+" -f "+pod+" -n "+namespace)
	}
}
// 进入bash
func bash(pod, namespace, sh string)  {
	if sh == "" {
		sh = "bash"
	}
	pod,namespace = getResAndNamespace("pod", pod, namespace)
	if pod != "" {
		base.Run("kubectl exec -it "+pod+" -n "+namespace+" -- "+sh)
	}
}
// 获取pod
func pod(key,namespace string)  {
	getResource("pod", key, namespace)
}

// 模糊查询namespace
func findNamespace(key string) string{
	ret := base.Exec("kubectl get namespace")
	rets := strings.Split(ret, "\n")
	var ns []string
	for _,s := range rets {
		strs := strings.Fields(s)
		if len(strs) == 0 {
			continue
		}
		s := strs[0]
		if key == "" || strings.Index(s, key) >=0 {
			ns = append(ns, s)
		}
	}
	if len(ns) == 0{
		fmt.Println("Namespace:"+key+" Not Found")
		return ""
	} else if len(ns) == 1{
		return ns[0]
	}
	name := ""
	var prompt survey.Prompt
	if len(ns)<=base.SEL_LIMIT{
		prompt = &survey.Select{
			Message:       "Choose a option:",
			Options:       ns,
		}
	} else {
		msg := ""
		for i,s := range ns {
			msg += strconv.Itoa(i)+": "+s+"\n"
		}
		prompt = &survey.Input{
			Message:  "The list:\n"+msg+"Input index:",
		}
	}
	survey.AskOne(prompt, &name)
	if len(ns)>base.SEL_LIMIT {
		idx,_ := strconv.Atoi(name)
		name = ns[idx]
	}
	return name
}

// 模糊过滤pod名称，并供用户选择，最终返回唯一的选择值
// 如果idx == 1 说明是在全局查询  要返回namespace
func findResByKey(txt,key string, idx int) string{
	var ps []string
	txts := strings.Split(txt, "\n")[1:]
	for _,str := range txts {
		strs := strings.Fields(str)
		if len(strs)>idx && (key == "@" || key == "" || strings.Index(strs[idx], key) >=0) {
			if idx == 1 {
				ps = append(ps, "pod:"+strs[idx]+" namespace:"+strs[0])
			} else {
				ps = append(ps, strs[idx])
			}
		}
	}

	name := ""
	if len(ps) == 0{
		return ""
	} else if len(ps) == 1 {
		name = ps[0]
	} else {
		var prompt survey.Prompt
		if len(ps)<=base.SEL_LIMIT{
			prompt = &survey.Select{
				Message:       "Choose a option:",
				Options:       ps,
			}
		} else {
			msg := ""
			for i,s := range ps {
				msg += strconv.Itoa(i)+": "+s+"\n"
			}
			prompt = &survey.Input{
				Message:  "The list:\n"+msg+"Input index:",
			}
		}
		survey.AskOne(prompt, &name)
		if len(ps)>base.SEL_LIMIT {
			idx,_ := strconv.Atoi(name)
			name = ps[idx]
		}
	}
	if idx == 1 {
		name = strings.ReplaceAll(name, "pod:", "")
		name = strings.ReplaceAll(name, "namespace:", "")
	}
	return name
}

// 获取namespace
func namespace(key string) {
	ret := base.Execp("kubectl get namespace")
	if key != ""{
		ret = base.FindByKey(ret, key)
	}
	fmt.Print(ret)
}
// 获取resource
func getResource(res,key,namespace string)  {
	// namespace为空则查询所有，不为空则查询指定【模糊】namespace
	ret := ""
	if namespace != ""{
		namespace = findNamespace(namespace)
		if namespace == ""{
			return
		}
		ret = base.Execp("kubectl get "+res+" -n "+namespace)
		ret = base.FindAllByKey(ret, key, 0)
	} else {
		ret = base.Execp("kubectl get "+res+" --all-namespaces")
		ret = base.FindAllByKey(ret, key, 1)
	}
	fmt.Print(ret)
}
// 查询指定的Res及namespace
func getResAndNamespace(res, key, namespace string) (string,string) {
	// namespace为空则查询所有，不为空则查询指定【模糊】namespace
	ret := ""
	if namespace != "" && namespace != "@"{
		namespace = findNamespace(namespace)
		if namespace == ""{
			return "",""
		}
		ret = base.Execp("kubectl get "+res+" -n "+namespace)
		ret = findResByKey(ret, key, 0)
	} else {
		ret = base.Execp("kubectl get "+res+" --all-namespaces")
		ret = findResByKey(ret, key, 1)
		if ret != "" {
			rets := strings.Fields(ret)
			ret = rets[0]
			namespace = rets[1]
		}
	}
	if ret == "" {
		fmt.Print(res+" Not Found!\n")
		return "",""
	}
	return ret,namespace
}
