package base

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"strconv"
	"strings"
)

var cols map[string]string = make(map[string]string)

func init() {
	cols["red"] = "\033[1;31;40m%s\033[0m"
}

/*
* 模糊查找指定列，并在该列中高亮显示, eg:
* str="default       ixx-demo-deployment-5df6465c5b-5xm5r   1/1     Running   2          75d"
* ret = FindAllByKey(str, "ixx", 1)
* "default       [red]ixx[red]-demo-deployment-5df6465c5b-5xm5r   1/1     Running   2          75d"
 */
func FindAllByKey(txt, key string, idx int) string {
	if key == "" || key == "@" {
		return txt
	}
	keyColor := fmt.Sprintf(cols["red"], key)
	ret := ""
	txts := strings.Split(txt, "\n")
	for _, str := range txts {
		strs := strings.Fields(str)
		if len(strs) > idx && strings.Index(strs[idx], key) >= 0 {
			fixColor := strings.ReplaceAll(strs[idx], key, keyColor)
			str = strings.ReplaceAll(str, strs[idx], fixColor)
			ret += str + "\n"
		}
	}
	if ret == "" {
		ret = "Not Found!\n"
	}
	return ret
}

// 过滤模糊的列，并供用户选择，最终返回唯一的选择值
func FindColByKey(txt, key string, idx int) string {
	var ps []string
	txts := strings.Split(txt, "\n")
	for _, str := range txts {
		strs := strings.Fields(str)
		if len(strs) > idx && strings.Index(strs[idx], key) >= 0 {
			ps = append(ps, strs[idx])
		}
	}

	if len(ps) == 0 {
		return ""
	}
	name := ""
	var prompt survey.Prompt
	if len(ps) <= SEL_LIMIT {
		prompt = &survey.Select{
			Message: "Choose a option:",
			Options: ps,
		}
	} else {
		msg := ""
		for i, s := range ps {
			msg += strconv.Itoa(i) + ": " + s + "\n"
		}
		prompt = &survey.Input{
			Message: "The list:\n" + msg + "Input index:",
		}
	}
	survey.AskOne(prompt, &name)
	if len(ps) > SEL_LIMIT {
		if name == "" || name == "q" {
			return ""
		}
		idx, _ := strconv.Atoi(name)
		name = ps[idx]
	}
	return name
}

// 整行模糊查找指定key, 并高亮显示
func FindByKey(ret, key string) string {
	return FindByKeySetColor(ret, key, "red")
}
func FindByKeySetColor(ret, key, col string) string {
	rets := strings.Split(ret, "\n")
	ret = ""
	if cols[col] != "" {
		col = cols[col]
	} else {
		col = cols["red"]
	}
	keyColor := fmt.Sprintf(col, key)
	for _, r := range rets {
		if strings.Index(r, key) >= 0 {
			r = strings.ReplaceAll(r, key, keyColor)
			ret += r + "\n"
		}
	}
	if ret == "" {
		ret = "Not Found!\n"
	}
	return ret
}
