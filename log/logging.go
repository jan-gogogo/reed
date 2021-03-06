package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger
var Clogger *logrus.Logger

type emptyWriter struct{}

func (ew emptyWriter) Write(p []byte) (int, error) {
	return 0, nil
}

func Init() {
	fileHooker := NewFileRotateHooker(os.Getenv("GOPATH")+"/src/github.com/reed", 60*60*24)

	Logger = logrus.New()
	Logger.Hooks.Add(fileHooker)
	Logger.Out = &emptyWriter{}
	Logger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}
	Logger.Level = logrus.InfoLevel

	Clogger = logrus.New()
	Clogger.Hooks.Add(fileHooker)
	Clogger.Out = &emptyWriter{}
	Clogger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}
	Clogger.Level = logrus.InfoLevel

}
