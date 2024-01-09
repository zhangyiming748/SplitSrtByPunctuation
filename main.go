package main

import (
	"fmt"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/SplitSrtByPunctuation/logic"
	"github.com/zhangyiming748/SplitSrtByPunctuation/util"
	"io"
	"log/slog"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

/*
todo
1. 第二段字幕有可能只有一个结束标点，如果直接上移动，这行没文字，考虑下移
2. 每次移动后想办法重读文件
3. 结束标点很有可能包含，。？！,.?!
*/
func init() {
	setLog()
}

var fresh []string

func main() {
	root := "/mnt/d/git/SplitSrtByPunctuation"
	files := GetFileInfo.GetAllFileInfo(root, "srt")
	for _, file := range files {
		if strings.Contains(file.PurgeName, "origin") {
			continue
		}
		trans(file.FullPath)
	}
}
func trans(srt string) {
	seed := rand.New(rand.NewSource(time.Now().Unix()))
	r := seed.Intn(2000)
	//中间文件名
	tmpname := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), strconv.Itoa(r), ".srt"}, "")
	before := util.ReadByLine(srt)
	after, _ := os.OpenFile(tmpname, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	srts := []string{}
	for i := 0; i < len(before); i += 4 {
		if i+3 > len(before) {
			continue
		}
		//after.WriteString(fmt.Sprintf("%s\n", before[i]))
		//after.WriteString(fmt.Sprintf("%s\n", before[i+1]))
		src := before[i+2]
		srts = append(srts, src)
	}
	n := logic.Balance(srts)
	count := 0
	for j := 0; j < len(before); j += 4 {
		if j+3 > len(before) {
			continue
		}
		after.WriteString(fmt.Sprintf("%s\n", before[j]))
		after.WriteString(fmt.Sprintf("%s\n", before[j+1]))
		after.WriteString(fmt.Sprintf("%s\n", n[count]))
		count++
		after.WriteString(fmt.Sprintf("%s\n", before[j+3]))
	}
	for _, v := range n {
		fmt.Println(v)
	}
	origin := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), "_origin", ".srt"}, "")
	exec.Command("cp", srt, origin).CombinedOutput()
	err := os.Rename(tmpname, srt)
	if err != nil {
		slog.Error("重命名出错")
	}

}
func setLog() {
	opt := slog.HandlerOptions{ // 自定义option
		AddSource: true,
		Level:     slog.LevelDebug, // slog 默认日志级别是 info
	}
	file := "Process.log"
	logf, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0770)
	if err != nil {
		panic(err)
	}
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(logf, os.Stdout), &opt))
	slog.SetDefault(logger)
}

func cpdatabase() {
	folderPath := "/data"
	_, err := os.Stat(folderPath)

	if os.IsNotExist(err) {
		fmt.Println("文件夹不存在")
	} else if err != nil {
		fmt.Println("发生错误：", err)
	} else {
		fmt.Println("文件夹存在")
		exec.Command("cp", "trans", "/data").Run()
	}
}
