package log

import "WhisperAndTrans/constant"
import (
	"github.com/zhangyiming748/lumberjack"
)

func SetLog(p *constant.Param) {
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
