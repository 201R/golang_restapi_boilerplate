package logger

import (
	"os"
	"path/filepath"

	"github.com/201R/go_api_boilerplate/packages/setting"
	"github.com/sirupsen/logrus"
)

type Level int

var logger logrus.Logger

func Setup() (func(), error) {
	c := setting.LoggerSetting
	logger.SetLevel(logrus.Level(c.Level))
	logger.SetFormatter(getFormatter(c.Format))

	var file *os.File
	if c.Output != "" {
		switch c.Output {
		case "stdout":
			logger.SetOutput(os.Stdout)
		case "stderr":
			logger.SetOutput(os.Stderr)
		case "file":
			if name := c.OutputFile; name != "" {
				_ = os.MkdirAll(filepath.Dir(name), 0777)

				f, err := os.OpenFile(name, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
				if err != nil {
					return nil, err
				}
				logger.SetOutput(f)
				file = f
			}
		}
	}

	return func() {
		if file != nil {
			file.Close()
		}
	}, nil
}

func getFormatter(format string) logrus.Formatter {
	switch format {
	case "json":
		return new(logrus.JSONFormatter)
	default:
		return new(logrus.TextFormatter)
	}
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	logger.Debugln(v)
}

// Debugf output logs at debug level
func Debugf(s string, v ...interface{}) {
	logger.Debugf(s, v...)
}

// Info output logs at info level
func Info(v ...interface{}) {
	logger.Infoln(v)
}

// Infof output logs at info level
func Infof(s string, v ...interface{}) {
	logger.Infof(s, v...)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	logger.Warnln(v)
}

// Warnf output logs at warn level
func Warnf(s string, v ...interface{}) {
	logger.Warnf(s, v...)
}

// Error output logs at error level
func Error(v ...interface{}) {
	logger.Errorln(v)
}

// Errorf output logs at error level
func Errorf(s string, v ...interface{}) {
	logger.Errorf(s, v...)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	logger.Fatalln(v)
}

// Fatalf output logs at fatal level
func Fatalf(s string, v ...interface{}) {
	logger.Fatalf(s, v...)
}
