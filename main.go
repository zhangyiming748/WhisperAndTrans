package main

import (
	"WhisperAndTrans/constant"
	"WhisperAndTrans/util"
	"log"
	"os/exec"
)

func main() {
	p := constant.Param{
		Root:     "/Users/zen/Github/FastYt-dlp/Hitomi",
		Language: "ja",
		Pattern:  "mp4",
		Model:    "base",
		Location: "/Users/zen/Github/FastYt-dlp",
	}
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
		}
	}
}
