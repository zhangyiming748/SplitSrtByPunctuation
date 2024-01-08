package logic

import (
	"fmt"
	"regexp"
	"testing"
)

// go test -v -run TestBalance
func TestBalance(t *testing.T) {
	str := "这.是一,个，测试字符串。"
	re := regexp.MustCompile(`[,.!?]`)
	matches := re.FindStringIndex(str)

	if matches != nil {
		fmt.Printf("第一个出现的标点符号是：%c\t位置在：%d", str[matches[0]], matches[1])
	} else {
		fmt.Println("没有找到标点符号")
	}
}
