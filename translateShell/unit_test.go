package translateShell

import (
	"WhisperAndTrans/constant"
	"WhisperAndTrans/sql"
	"fmt"
	"github.com/zhangyiming748/DeepLX"
	"io"
	"net/http"
	"strings"
	"testing"
)

// go test -v -run TestTranslate
//func TestDeepXl(t *testing.T) {
//	ret := DeepXl("hello,world")
//	t.Log(ret)
//}

func TestWeb(t *testing.T) {

	url := "http://192.168.1.6:1188/translate"
	method := "POST"

	payload := strings.NewReader(`{
    "text": "Hello, world!",
    "source_lang": "auto",
    "target_lang": "ZH"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
func TestDeepLX(t *testing.T) {
	lx, err := DeepLx.TranslateByDeepLX("auto", "zh", "hello world", "")
	if err != nil {
		return
	} else {
		t.Log(lx)
	}
}

// go test -v -run TestTrans
func TestTrans(t *testing.T) {
	f := "/mnt/e/video/cod2.srt"
	p := constant.Param{
		Root:     "/mnt/e/video",
		Language: "",
		Pattern:  "srt",
		Model:    "",
		Location: "",
		Proxy:    "192.168.1.20:8889",
	}
	c := constant.Count{
		Bing:   0,
		Google: 0,
		Deeplx: 0,
		Cache:  0,
	}
	sql.SetLevelDB(&p)
	Trans(f, &p, &c)
}
