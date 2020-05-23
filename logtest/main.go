package main

import (
	"go_learning/hlog"
)

func main() {
	logFile := hlog.NewLogFile(hlog.WARNING, "./", "huangwneyu.log", 256, 100)
	for {
		logFile.Debug("abc")
		logFile.Error("abc")
		logFile.Warning("abc")
	}
}
