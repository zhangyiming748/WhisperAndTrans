package main

import (
	"WhisperAndTrans/constant"
	"WhisperAndTrans/replace"
	"WhisperAndTrans/sql"
	"WhisperAndTrans/translateShell"
	"WhisperAndTrans/util"
	"fmt"
	"github.com/zhangyiming748/lumberjack"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func init() {
	if !util.IsExistCmd("whisper") {
		log.Fatal("whisper未安装")
	}
	if !util.IsExistCmd("trans") {
		log.Fatal("trans未安装")
	}
}
func main() {
	p := &constant.Param{
		Root:     "/Users/zen/Github/FastYt-dlp/eu",
		Language: "Russian",
		Pattern:  "mp4",
		Model:    "base",
		Location: "/Users/zen/Github/FastYt-dlp",
		Proxy:    "192.168.1.20:8889",
	}
	if root := os.Getenv("root"); root != "" {
		p.SetRoot(root)
	}
	if language := os.Getenv("language"); language != "" {
		p.SetLanguage(language)
	}
	if pattern := os.Getenv("pattern"); pattern != "" {
		p.SetPattern(pattern)
	}
	if model := os.Getenv("model"); model != "" {
		p.SetModel(model)
	}
	if location := os.Getenv("location"); location != "" {
		p.SetLocation(location)
	}
	if proxy := os.Getenv("proxy"); proxy != "" {
		p.SetProxy(proxy)
	}

	c := new(constant.Count)
	defer func() {
		log.Printf("\r从bing获取:%d条\n从google获取:%d条\n从deeplx获取:%d条\n从cache获取:%d条\n", c.GetBing(), c.GetGoogle(), c.GetDeeplx(), c.GetCache())
	}()
	setLog(p)
	sql.Initial(p)
	replace.SetSensitive(p)
	seed := rand.New(rand.NewSource(time.Now().Unix()))
	files, _ := util.GetAllFileInfoFast(p.GetRoot(), p.GetPattern())
	for _, file := range files {
		log.Printf("文件名:%v\n", file)
		//whisper true.mp4 --model base --language English --model_dir /Users/zen/Whisper --output_format srt
		//cmd := exec.Command("whisper", file.FullPath, "--model", level, "--model_dir", location, "--language", language, "--output_dir", root, "--verbose", "True")
		cmd := exec.Command("whisper", file, "--model", p.GetModel(), "--model_dir", p.GetLocation(), "--output_format", "srt", "--prepend_punctuations", ",.?", "--language", p.GetLanguage(), "--output_dir", p.GetRoot(), "--verbose", "True")
		err := util.ExecCommand(cmd)
		if err != nil {
			log.Printf("当前字幕生成错误\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
		} else {
			// todo 翻译字幕
			r := seed.Intn(2000)
			//中间文件名
			srt := strings.Replace(file, p.GetPattern(), "srt", 1)
			tmpname := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), strconv.Itoa(r), ".srt"}, "")
			before := util.ReadByLine(srt)
			after, _ := os.OpenFile(tmpname, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
			for i := 0; i < len(before); i += 4 {
				if i+3 > len(before) {
					continue
				}
				after.WriteString(fmt.Sprintf("%s\n", before[i]))
				after.WriteString(fmt.Sprintf("%s\n", before[i+1]))
				src := before[i+2]

				afterSrc := replace.GetSensitive(src)

				var dst string

				if get, err := sql.GetDatabase().Hash().Get("translations", src); err == nil {
					dst = get.String()
					fmt.Println("find in cache")
					c.SetCache()
				} else {
					dst = translateShell.Translate(afterSrc, p, c)
					var count int
					for replace.Falied(dst) {
						if count > 3 {
							log.Printf("重试三次后依然失败srt=%v\tdst=%v\n", afterSrc, dst)
							dst = replace.Hans(dst)
							break
						}
						log.Printf("查询失败\t重试%v\n", count)
						time.Sleep(1 * time.Second)
						dst = translateShell.Translate(afterSrc, p, c)
						count++
					}
				}
				dst = replace.GetSensitive(dst)
				sql.GetDatabase().Hash().Set("translations", src, dst)
				log.Printf("文件名:%v\n原文:%v\n译文:%v\n", tmpname, src, dst)
				after.WriteString(fmt.Sprintf("%s\n", src))
				after.WriteString(fmt.Sprintf("%s\n", dst))
				after.WriteString(fmt.Sprintf("%s\n", before[i+3]))
				after.Sync()
			}
			origin := strings.Join([]string{strings.Replace(srt, ".srt", "", 1), "_origin", ".srt"}, "")
			exec.Command("cp", srt, origin).CombinedOutput()
			os.Rename(tmpname, srt)
		}
	}
}
func setLog(p *constant.Param) {
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   strings.Join([]string{p.GetRoot(), "WhisperAndTrans.log"}, string(os.PathSeparator)),
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
	}
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
