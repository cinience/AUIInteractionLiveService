package log

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"path/filepath"
)

var (
	accessLogger *logrus.Logger
)

func NewAccessLogger(path string) {
	accessLogger = logrus.New()
	accessLogFilePath := filepath.Join(filepath.Dir(path), "access.log")
	accessLogFile := &lumberjack.Logger{
		Filename:   accessLogFilePath,
		MaxSize:    512, // Max megabytes before log is rotated
		MaxBackups: 5,   // Max number of old log files to keep
		MaxAge:     2,   // Max number of days to retain log files
		Compress:   false,
	}
	accessLogger.SetOutput(accessLogFile)
	accessLogger.SetFormatter(&logrus.JSONFormatter{})

}
func Accessf(fields map[string]interface{}) {
	accessLogger.WithFields(fields).Info()
}
