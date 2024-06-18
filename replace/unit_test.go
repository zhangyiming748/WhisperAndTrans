package replace

import "testing"

func TestHans(t *testing.T) {
	str := "Hello, \u001B你好！123abc"
	ret := Hans(str)

	t.Log(ret)
}
