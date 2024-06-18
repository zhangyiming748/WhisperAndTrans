package translateShell

import (
	"WhisperAndTrans/util"
	"encoding/json"
)

/*
curl --location 'https://api.deeplx.org/translate' \
--header 'Content-Type: application/json' \

	--data '{
	    "text": "Hello, world!",
	    "source_lang": "EN",
	    "target_lang": "ZH"
	}'
*/
type ans struct {
	Code         int      `json:"code"`
	Id           int64    `json:"id"`
	Data         string   `json:"data"`
	Alternatives []string `json:"alternatives"`
}

func DeepXl(src string) string {
	uri := "https://api.deeplx.org/translate"
	headers := map[string]string{
		"content-type": "application/json",
	}
	data := map[string]string{
		"text":        src,
		"source_lang": "auto",
		"target_lang": "ZH",
	}
	b, _ := util.HttpPostJson(headers, data, uri)
	var a ans
	json.Unmarshal(b, &a)
	return a.Data
}
