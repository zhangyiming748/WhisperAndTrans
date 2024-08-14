package main

import (
	"WhisperAndTrans/constant"
	mylog "WhisperAndTrans/log"
	"WhisperAndTrans/replace"
	"WhisperAndTrans/sql"
	"WhisperAndTrans/translateShell"
	"WhisperAndTrans/util"
	"log"
	"os"
	"os/exec"
)

func init() {
	if !util.IsExistCmd("whisper") {
		log.Fatal("whisper未安装")
	}
	//if !util.IsExistCmd("trans") {
	//	log.Fatal("trans未安装")
	//}
}
func main() {
	p := &constant.Param{
		Root:     "/mnt/e/video/Download",
		Language: "English",
		Pattern:  "mp3",
		Model:    "large-v3",
		Location: "/mnt/c/Users/zen/Downloads",
		Proxy:    "",
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

	mylog.SetLog(p)
	sql.Initial(p)
	replace.SetSensitive(p)
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
			translateShell.Translate(file, p, c)
		}
	}
}
