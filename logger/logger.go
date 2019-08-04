package logger

/*
	There is a lot more that can be done with the logging module. This is just a basic POC implementation of the logging module
*/
import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
)

var appLog *logrus.Logger
var webLog *logrus.Logger

// Init initializes the app and web loggers
/*
	Input:
		- none
	Output:
		- none
*/
func Init() {
	standardLogger := logrus.New()
	standardLogger.Formatter = &logrus.JSONFormatter{}
	standardLogger.Level = logrus.InfoLevel
	standardFields := logrus.Fields{
		"appname": "pwshed",
	}
	standardLogger.WithFields(standardFields)

	appLog = standardLogger
	webLog = standardLogger
}

// Info sets custom fields before writing the log
/*
	Input:
		- path - string - path of the api call or "cli" OR function name for the applog
		- message - string - message for the log
	Output:
		- none
*/
func Info(path, message string) {
	_, file, line, _ := runtime.Caller(1)
	pwd := strings.Split(file, "/")
	file = pwd[len(pwd)-1]
	webLog.WithFields(logrus.Fields{"path": path, "time": time.Now()}).Info(message)
	appLog.WithFields(logrus.Fields{"path": fmt.Sprintf("%s - L%d", file, line), "time": time.Now()}).Info(message)
}

// Warn sets custom fields before writing the log
/*
	Input:
		- path - string - path of the api call or "cli"
		- message - string - message for the log
	Output:
		- none
*/
func Warn(path, message string) {
	_, file, line, _ := runtime.Caller(1)
	pwd := strings.Split(file, "/")
	file = pwd[len(pwd)-1]
	webLog.WithFields(logrus.Fields{"path": path, "time": time.Now()}).Warn(message)
	appLog.WithFields(logrus.Fields{"path": fmt.Sprintf("%s - L%d", file, line), "time": time.Now()}).Warn(message)
}

// Error sets custom fields before writing the log
/*
	Input:
		- path - string - path of the api call or "cli"
		- message - string - message for the log
	Output:
		- none
*/
func Error(path, message string) {
	_, file, line, _ := runtime.Caller(1)
	pwd := strings.Split(file, "/")
	file = pwd[len(pwd)-1]
	webLog.WithFields(logrus.Fields{"path": path, "time": time.Now()}).Error(message)
	appLog.WithFields(logrus.Fields{"path": fmt.Sprintf("%s - L%d", file, line), "time": time.Now()}).Error(message)
}
