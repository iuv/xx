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
	case "n","ns","namespace": namespace(a1)
	case "e","exec": bash(a1, a2, a3)
	case "l","log": logs(a1, a2, a3)
	case "p","pod": getResource("pod", a1, a2)
	case "d","deployment": getResource("deployment", a1, a2)
	case "i","ingress": getResource("ingress", a1, a2)
	case "s","service": getResource("service", a1, a2)
	case "c","configmap": getResource("configmap", a1, a2)
	case "sec","secret": getResource("secret", a1, a2)
	case "ss","statefulset": getResource("statefulset", a1, a2)
	case "pd","podd": resourceDescribe("pod", a1, a2)
	case "dd","deploymentd": resourceDescribe("deployment", a1, a2)
	case "id","ingressd": resourceDescribe("ingress", a1, a2)
	case "sd","serviced": resourceDescribe("service", a1, a2)
	case "cd","configmapd": resourceDescribe("configmap", a1, a2)
	case "secd","secretd": resourceDescribe("secret", a1, a2)
	case "ssd","statefulsetd": resourceDescribe("statefulset", a1, a2)
	case "py","pody": resourceYaml("pod", a1, a2, a3)
	case "dy","deploymenty": resourceYaml("deployment", a1, a2, a3)
	case "iy","ingressy": resourceYaml("ingress", a1, a2, a3)
	case "sy","servicey": resourceYaml("service", a1, a2, a3)
	case "cy","configmapy": resourceYaml("configmap", a1, a2, a3)
	case "secy","secrety": resourceYaml("secret", a1, a2, a3)
	case "ssy","statefulsety": resourceYaml("statefulset", a1, a2, a3)
	case "pdel","poddel": resourceDel("pod", a1, a2)
	case "ddel","deploymentdel": resourceDel("deployment", a1, a2)
	case "idel","ingressdel": resourceDel("ingress", a1, a2)
	case "sdel","servicedel": resourceDel("service", a1, a2)
	case "secdel","secretdel": resourceDel("secret", a1, a2)
	case "ssdel","statefulsetdel": resourceDel("statefulset", a1, a2)
	case "pe","pode": resourceEdit("pod", a1, a2)
	case "de","deploymente": resourceEdit("deployment", a1, a2)
	case "ie","ingresse": resourceEdit("ingress", a1, a2)
	case "se","servicee": resourceEdit("service", a1, a2)
	case "ce","configmape": resourceEdit("configmap", a1, a2)
	case "sece","secrete": resourceEdit("secret", a1, a2)
	case "sse","statefulsete": resourceEdit("statefulset", a1, a2)
	case "a","apply": apply(a1)
	case "copy": cp(a1, a2, a3, a4)
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
// 查询资源yaml基本方法
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
// 查询资源详情基本方法
func resourceDescribe(res, key, namespace string)  {
	key, namespace = getResAndNamespace(res, key, namespace)
	if key != "" {
		ret := base.Execp("kubectl describe "+res+" "+ key +" -n "+namespace)
		fmt.Print(ret)
	}
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
		if name=="" || name=="q" {
			return ""
		}
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
			if name=="" || name=="q" {
				return ""
			}
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

// 编辑资源基本方法
func resourceEdit(res, key, namespace string)  {
	key, namespace = getResAndNamespace(res, key, namespace)
	if key != "" {
		base.Run("kubectl edit "+res+" "+ key +" -n "+namespace)
	}
}