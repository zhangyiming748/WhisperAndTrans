package translateShell

import (
	"testing"
)

// go test -v -run TestTranslate
func TestDeepXl(t *testing.T) {
	ret := DeepXl("hello,world")
	t.Log(ret)
}
