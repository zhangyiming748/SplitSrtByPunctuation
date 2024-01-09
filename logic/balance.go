package logic

import (
	"fmt"
	"github.com/zhangyiming748/SplitSrtByPunctuation/replace"
	"log/slog"
	"regexp"
	"strings"
)

func Balance(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		//fmt.Println(s[i])
		if strings.HasSuffix(s[i], ",") || strings.HasSuffix(s[i], ".") || strings.HasSuffix(s[i], "?") || strings.HasSuffix(s[i], "!") || strings.HasSuffix(s[i], "...") {
			fmt.Printf("完整的句子:%s\n", s[i])
		} else {
			fmt.Printf("不完整的句子:%s\n", s[i])
			// 找下一句的标点符号
			c := GetFirst(s[i+1])
			first := fmt.Sprintf("%c", c)
			prefix := strings.Split(s[i+1], first)[0]
			s[i] = strings.Join([]string{s[i], " ", prefix, first}, "")
			s[i] = replace.ChinesePunctuation(s[i])
			s[i+1] = strings.Replace(s[i+1], strings.Join([]string{prefix, first}, ""), "", 1)
			s[i+1] = replace.ChinesePunctuation(s[i+1])
			// 如果第二行开头是空格 删掉这个空格
			if strings.HasPrefix(s[i+1], " ") {
				s[i+1] = s[i+1][1:]
			}
			slog.Debug("分割后", slog.String("前半部分", prefix), slog.String("组成完整的一句话", s[i]), slog.String("删除后的下一句话", s[i+1]))
		}
	}
	return s
}

/*
获取第一个出现的标点符号
*/
func GetFirst(str string) uint8 {
	//str := "这是一个测试字符串。"
	re := regexp.MustCompile(`[,.?!]`)
	matches := re.FindStringIndex(str)
	if matches != nil {
		fmt.Printf("第一个出现的标点符号是：%c\n", str[matches[0]])
		return str[matches[0]]
	} else {
		fmt.Println("没有找到标点符号")
		return 0
	}
}
