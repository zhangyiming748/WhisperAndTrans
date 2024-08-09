package replace

import (
	"WhisperAndTrans/constant"
	"WhisperAndTrans/sql"
	"WhisperAndTrans/util"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var Sensitive = map[string]string{}

func GetSensitive(str string) string {
	for k, v := range Sensitive {
		if strings.Contains(str, k) {
			str = strings.Replace(str, k, v, -1)
			log.Printf("替换生效\tbefore:%v\tafter:%v\t替换之后的完整句子:%v\n", k, v, str)
		}
	}
	return str
}

func SetSensitive(p *constant.Param) {
	fp1 := strings.Join([]string{p.GetRoot(), "sensitive.txt"}, string(os.PathSeparator))
	fp2 := "sensitive.txt"
	lines := []string{}
	if util.IsExist(fp1) {
		log.Printf("从视频目录%v中加载敏感词\n", fp1)
		lines = readByLine(fp1)
	}
	if util.IsExist(fp2) {
		log.Printf("从程序目录%v中加载敏感词\n", fp2)
		lines = readByLine(fp1)
	} else {
		log.Println("没有找到敏感词文件")
	}
	for _, line := range lines {
		before := strings.Split(line, ":")[0]
		after := strings.Split(line, ":")[1]
		log.Printf("敏感词:\tbefore:%v\tafter:%v\t", before, after)
		Sensitive[before] = after
		set, err := sql.GetDatabase().Hash().Set("sensitive", before, after)
		if err != nil {
			log.Println("敏感词写入数据库失败")
		} else if set {
			log.Printf("写入数据库成功\n")
		}
	}
}

/*
所有符号替换为空格
*/
func space(str string) string {
	str = strings.Replace(str, "。", " ", -1)
	str = strings.Replace(str, "，", " ", -1)
	str = strings.Replace(str, "《", " ", -1)
	str = strings.Replace(str, "》", " ", -1)
	str = strings.Replace(str, "【", " ", -1)
	str = strings.Replace(str, "】", " ", -1)
	str = strings.Replace(str, "（", " ", -1)
	str = strings.Replace(str, "）", " ", -1)
	str = strings.Replace(str, "「", " ", -1)
	str = strings.Replace(str, "」", " ", -1)
	str = strings.Replace(str, "+", " ", -1)
	str = strings.Replace(str, ".", " ", 1)
	str = strings.Replace(str, ",", " ", -1)
	str = strings.Replace(str, "(", " ", -1)
	str = strings.Replace(str, ")", " ", -1)
	str = strings.Replace(str, "(", " ", -1)
	str = strings.Replace(str, ")", " ", -1)
	str = strings.Replace(str, "(", " ", -1)
	str = strings.Replace(str, ")", " ", -1)
	str = strings.Replace(str, "(", " ", -1)
	str = strings.Replace(str, ")", " ", -1)
	str = strings.Replace(str, "_", " ", -1)
	str = strings.Replace(str, "`", " ", -1)
	str = strings.Replace(str, "·", " ", -1)
	str = strings.Replace(str, "、", " ", -1)
	str = strings.Replace(str, "！", " ", -1)
	str = strings.Replace(str, "|", " ", -1)
	str = strings.Replace(str, "｜", " ", -1)
	str = strings.Replace(str, ":", " ", -1)
	str = strings.Replace(str, " ", " ", -1)
	str = strings.Replace(str, "&", " ", -1)
	str = strings.Replace(str, "？", " ", -1)
	str = strings.Replace(str, "(", " ", -1)
	str = strings.Replace(str, ")", " ", -1)
	str = strings.Replace(str, "-", " ", -1)
	str = strings.Replace(str, " ", " ", -1)
	str = strings.Replace(str, "“", " ", -1)
	str = strings.Replace(str, "”", " ", -1)
	str = strings.Replace(str, "--", " ", -1)
	str = strings.Replace(str, "_", " ", -1)
	str = strings.Replace(str, "：", " ", -1)
	str = strings.Replace(str, "\n", "", -1)
	return str
}
func Hans(input string) string {
	//input := "Hello, 你好！123abc"
	input = space(input)
	done := ""
	//reg := regexp.MustCompile(`\p{Han}|\d|[a-zA-Z]|\s`)
	reg := regexp.MustCompile(`\p{Han}|\d|\s`)
	matches := reg.FindAllString(input, -1)
	for _, match := range matches {
		//fmt.Printf("%d,%s", i, match)
		done = strings.Join([]string{done, match}, "")
	}
	done = remove331x220(done)
	return done
}

/*
golang 实现 从字符串中找到以331开头220结尾的子字符串 删除后返回新的字符串
*/
func remove331x220(s string) string {
	re := regexp.MustCompile(`331.*?220`)
	return re.ReplaceAllString(s, "")
}
func readByLine(fp string) []string {
	lines := []string{}
	fi, err := os.Open(fp)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		log.Println("按行读文件出错")
		return []string{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lines = append(lines, string(a))
	}
	return lines
}
