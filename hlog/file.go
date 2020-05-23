package hlog

import (
	"fmt"
	"os"
	"path"
	"time"
)

// LogFile print log to file
type LogFile struct {
	logLevel    uint
	filePath    string
	fileName    string
	fileObj     *os.File
	fileErrObj  *os.File
	maxFileSize int64
	logChan     chan *logMsg
}

type logMsg struct {
	logLevel  uint
	msg       string
	funcName  string
	fileName  string
	timestamp string
	line      int
}

// NewLogFile create and return a LogFile struct
func NewLogFile(level uint, filePath, fileName string, maxFileSize int64, maxSize int64) (logFile *LogFile) {
	logFile = &LogFile{
		logLevel:    level,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxFileSize,
		logChan:     make(chan *logMsg, maxSize),
	}
	err := logFile.initFile()
	if err != nil {
		panic(err)
	}
	go logFile.writeLogBackground()
	return
}

func (l *LogFile) initFile() error {
	fullPath := path.Join(l.filePath, l.fileName)
	fileObj, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.FileMode(0777))
	if err != nil {
		return err
	}
	fileErrObj, err := os.OpenFile(fullPath+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	l.fileObj = fileObj
	l.fileErrObj = fileErrObj
	return nil
}

func (l *LogFile) closeFile() {
	l.fileObj.Close()
	l.fileErrObj.Close()
}

func (l *LogFile) checkSize(file *os.File) (bool, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return false, err
	}
	size := fileInfo.Size()
	if size >= l.maxFileSize {
		return true, nil
	}
	return false, nil
}

func (l *LogFile) spliteFileBySize(file *os.File) (*os.File, error) {
	ok, err := l.checkSize(l.fileObj)
	if err != nil {
		return nil, err
	}
	if ok {
		// 1. backup
		fileInfo, err := file.Stat()
		if err != nil {
			return nil, err
		}
		fileName := fileInfo.Name()
		nowStr := time.Now().Format("20060102150405")
		logName := path.Join(l.filePath, fileName)
		newLogName := fmt.Sprintf("%s%s.bak", nowStr, logName)
		os.Rename(logName, newLogName)
		// 2. close file
		file.Close()
		// 3. open a new file
		fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return nil, err
		}
		// 4. return new fileobject
		return fileObj, nil
	}
	return file, nil
}

func (l *LogFile) print(level uint, format string, a ...interface{}) {
	if level >= l.logLevel {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		log := &logMsg{
			logLevel:  level,
			msg:       msg,
			funcName:  funcName,
			fileName:  fileName,
			line:      lineNo,
			timestamp: now.Format("2006-01-02 15:04:05"),
		}
		select {
		case l.logChan <- log:
		default: //把日志扔了
		}
		/* newFileObj, err := l.spliteFileBySize(l.fileObj)
		if err != nil {
			panic(err)
		}
		l.fileObj = newFileObj
		fmt.Fprintf(l.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), parseIntToStr(level), fileName, funcName, lineNo, msg)
		if level >= ERROR {
			newFileObj, err := l.spliteFileBySize(l.fileErrObj)
			if err != nil {
				panic(err)
			}
			l.fileErrObj = newFileObj
			fmt.Fprintf(l.fileErrObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), parseIntToStr(level), fileName, funcName, lineNo, msg)
		} */
	}
}

func (l *LogFile) writeLogBackground() {
	for {
		select {
		case log := <-l.logChan:
			newFileObj, err := l.spliteFileBySize(l.fileObj)
			if err != nil {
				panic(err)
			}
			l.fileObj = newFileObj
			fmt.Fprintf(l.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", log.timestamp, parseIntToStr(log.logLevel), log.fileName, log.funcName, log.line, log.msg)
			if log.logLevel >= ERROR {
				newFileObj, err := l.spliteFileBySize(l.fileErrObj)
				if err != nil {
					panic(err)
				}
				l.fileErrObj = newFileObj
				fmt.Fprintf(l.fileErrObj, "[%s] [%s] [%s:%s:%d] %s\n", log.timestamp, parseIntToStr(log.logLevel), log.fileName, log.funcName, log.line, log.msg)
			}
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// Debug prints the information of debug level
func (l *LogFile) Debug(format string, a ...interface{}) {
	l.print(DEBUG, format, a...)
}

// Trace prints the information of trace level
func (l *LogFile) Trace(format string, a ...interface{}) {
	l.print(TRACE, format, a...)
}

// Info prints the information of info level
func (l *LogFile) Info(format string, a ...interface{}) {
	l.print(INFO, format, a...)
}

// Warning prints the information of warning level
func (l *LogFile) Warning(format string, a ...interface{}) {
	l.print(WARNING, format, a...)
}

// Error prints the information of error level
func (l *LogFile) Error(format string, a ...interface{}) {
	l.print(ERROR, format, a...)
}

// Fatal prints the information of fatal level
func (l *LogFile) Fatal(format string, a ...interface{}) {
	l.print(FATAL, format, a...)
}
