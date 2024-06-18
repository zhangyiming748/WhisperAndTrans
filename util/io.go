package util

/*
获取目录下符合条件的所有文件基础信息
*/
import (
	"os"
	"path/filepath"
	"strings"
)

func GetAllFileInfoFast(dir, pattern string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), pattern) {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil
}
