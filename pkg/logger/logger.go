package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
)

func Init() {
	log := logrus.New()
	log.SetReportCaller(true)
	log.Formatter = &logrus.JSONFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
	}
	log.SetOutput(os.Stdout)
}
