package hlog

import (
	"fmt"
	"os"
	"time"
)

// LogConsole is used to be a receiver of method
type LogConsole struct {
	logLevel uint
}

func (l LogConsole) print(level uint, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Fprintf(os.Stdout, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"),
		parseIntToStr(level), fileName, funcName, lineNo, msg)
}

// NewLogConsole return a log struct
// level depends on the parameter "level"
func NewLogConsole(level uint) LogConsole {
	return LogConsole{
		level,
	}
}

// NewLogConsoleNull return a log struct
// default level is debug
func NewLogConsoleNull() LogConsole {
	return LogConsole{0}
}

// Debug prints the information of debug level
func (l LogConsole) Debug(format string, a ...interface{}) {
	if l.logLevel <= DEBUG {
		l.print(DEBUG, format, a...)
	}
}

// Trace prints the information of trace level
func (l LogConsole) Trace(format string, a ...interface{}) {
	if l.logLevel <= TRACE {
		l.print(TRACE, format, a...)
	}
}

// Info prints the information of info level
func (l LogConsole) Info(format string, a ...interface{}) {
	if l.logLevel <= INFO {
		l.print(INFO, format, a...)
	}
}

// Warning prints the information of warning level
func (l LogConsole) Warning(format string, a ...interface{}) {
	if l.logLevel <= WARNING {
		l.print(WARNING, format, a...)
	}
}

// Error prints the information of error level
func (l LogConsole) Error(format string, a ...interface{}) {
	if l.logLevel <= ERROR {
		l.print(ERROR, format, a...)
	}
}

// Fatal prints the information of fatal level
func (l LogConsole) Fatal(format string, a ...interface{}) {
	if l.logLevel <= FATAL {
		l.print(FATAL, format, a...)
	}
}
