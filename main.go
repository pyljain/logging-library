package main

import (
	"logging-clone/pkg/mylogger"
	"time"
)

func main() {
	logger := mylogger.NewMyLogger().
		WithString().
		WithStdIO()

	logger.Info("hello", mylogger.WithLabel("app", "r2d2"))
	logger.Debug("hello 1")

	time.Sleep(5 * time.Second)
}
