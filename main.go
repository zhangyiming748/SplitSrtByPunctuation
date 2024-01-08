package main

import (
	"fmt"
	"github.com/zhangyiming748/GetFileInfo"
	"github.com/zhangyiming748/SplitSrtByPunctuation/logic"
	"github.com/zhangyiming748/SplitSrtByPunctuation/util"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"strings"
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
	//seed := rand.New(rand.NewSource(time.Now().Unix()))
	//r := seed.Intn(2000)
	//中间文件名
	//tmpname := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), strconv.Itoa(r), ".srt"}, "")
	before := util.ReadByLine(srt)
	srts := []string{}
	for i := 0; i < len(before); i += 4 {
		if i+3 > len(before) {
			continue
		}
		//after.WriteString(fmt.Sprintf("%s\n", before[i]))
		//after.WriteString(fmt.Sprintf("%s\n", before[i+1]))
		src := before[i+2]
		srts = append(srts, src)

		//slog.Info("", slog.String("文件名", tmpname), slog.String("原文", src), slog.String("译文", dst))
		//after.WriteString(fmt.Sprintf("%s\n", src))
		//after.WriteString(fmt.Sprintf("%s\n", dst))
		//fresh = append(fresh, dst)

		//after.WriteString(fmt.Sprintf("%s\n", before[i+3]))
		//after.Sync()
	}
	n := logic.Balance(srts)

	for _, v := range n {
		fmt.Println(v)
	}
	//origin := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), "_origin", ".srt"}, "")
	//exec.Command("cp", srt, origin).CombinedOutput()
	//os.Rename(tmpname, srt)

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
