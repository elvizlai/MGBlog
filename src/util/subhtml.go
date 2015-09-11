/**
 * Created by elvizlai on 2015/8/25 14:21
 * Copyright © PubCloud
 */
package util
import (
	"errors"
	"regexp"
	"strings"
)

//整个程序分为3大部分
//1、正则过滤
//2、筛选固定长度
//3、标签检查及补全

func SubHtml(content string, limit int) string {
	if len(content) < limit {
		return content
	}

	content = strings.Replace(content, "\n", "", -1)

	//-->正则过滤获取content中的所有标签
	all := regexp.MustCompile(`<[\S\s]+?>`)
	allResult := all.FindAllString(content, -1) //获取到的所有标签

	//	for _, v := range allResult {
	//		fmt.Println(v)
	//	}

	//所有标签都提取正确，无需检验以上内容

	needRemove := regexp.MustCompile(`<[\S\s]+?/>`)//筛选项目，匹配诸如<br/>之类的标签

	//筛选固定长度-->limit
	var i, totalLength int
	cpOfContent := content

	//计算所需要截取的长度
	for strLength := 0; i < len(allResult); i++ {
		index := strings.Index(content, allResult[i])//标签索引 -- 这部分是按照byte算的
		lenOfIndex := strings.Count(allResult[i], "") - 1//标签长度 -- 按照rune计算

		//fmt.Println(allResult[i], index, lenOfIndex)

		//如果index==0，则表明为头标签，需要截取长度等于标签长度，字符串长度不需要计算
		if index == 0 {
			totalLength += lenOfIndex
		}else {
			sub := []byte(content)[:index]
			//fmt.Println(string(sub))
			strLength += len([]rune(string(sub)))
			totalLength = len([]rune(string(sub))) + lenOfIndex + totalLength
			content = strings.Replace(content, string(sub), "", 1)
		}

		if strLength >= limit {
			break
		}

		content = strings.Replace(content, allResult[i], "", 1)
	}

	//截取后的字符串
	if totalLength == 0 {
		cpOfContent = string([]rune(cpOfContent)[:limit])
	}else {
		cpOfContent = string([]rune(cpOfContent)[:totalLength])
	}

	//-->修复标签，如果i==len(allResult)，则无需修复标签
	if i != len(allResult) {
		allResult = allResult[:i + 1]

		//移除不需要补全的标签
		for i := 0; i < len(allResult); i++ {
			if len(needRemove.FindStringIndex(allResult[i])) == 2 {
				allResult = append(allResult[:i], allResult[i + 1:]...)
				i--
			}
			if allResult[i] == "<br>" || strings.HasPrefix(allResult[i], "<img") {
				allResult = append(allResult[:i], allResult[i + 1:]...)
				i--
			}
		}

		//使用栈来维护标签是否配对
		stack := NewStack(len(allResult))

		for i := 0; i < len(allResult); i++ {
			if strings.HasPrefix(allResult[i], "</") {
				stack.Pop()
			}else {
				stack.Push(allResult[i])
			}
		}


		//将所有不配对的补齐
		for {
			if stack.Len() != 0 {
				str, _ := stack.Pop()
				if str == "<br>" || strings.Contains(str, "<img") {
					continue
				}

				index := strings.Index(str, " ")//??
				if index == -1 {
					index = len(str) - 1
				}
				cpOfContent += "</" + string([]byte(str)[1:index]) + ">"
			}else {
				break
			}
		}
	}

	return cpOfContent
}

type Stack struct {
	st  []string
	len int
	cap int
}

func NewStack(cap int) *Stack {
	st := make([]string, 0, cap)
	return &Stack{st, 0, cap}
}

func (this *Stack)Push(p string) {
	this.st = append(this.st, p)
	this.len = len(this.st)
	this.cap = cap(this.st)
}

func (this *Stack)Pop() (string, error) {
	if this.len == 0 {
		return "", errors.New("Can't pop an empty stack")
	}
	this.len -= 1
	out := this.st[this.len]
	this.st = this.st[:this.len]
	return out, nil
}

func (this *Stack)Len() int {
	return this.len
}

func (this *Stack)Cap() int {
	return this.cap
}