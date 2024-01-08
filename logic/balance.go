package logic

import (
	"fmt"
	"strings"
)

func Balance(s []string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
		if strings.HasSuffix(s[i], ",") || strings.HasSuffix(s[i], ".") || strings.HasSuffix(s[i], "?") || strings.HasSuffix(s[i], "!") || strings.HasSuffix(s[i], "...") {
			fmt.Printf("完整的句子:%s\n", s[i])
		} else {
			fmt.Printf("不完整的句子:%s\n", s[i])
		}
	}
}
