package mylogger

import "time"

func (ml *MyLogger) WithJSON() *MyLogger {
	ml.format = "json"
	return ml
}

func (ml *MyLogger) WithString() *MyLogger {
	ml.format = "string"
	return ml
}

func (ml *MyLogger) WithStdIO() *MyLogger {
	ml.destination = "stdio"
	return ml
}

func (ml *MyLogger) WithLogFile(location string) *MyLogger {
	ml.destination = "file"
	ml.logFile = location
	return ml
}

func (ml *MyLogger) WithMinLogLevel(level LogLevel) *MyLogger {
	ml.minLogLevel = level
	return ml
}

func (ml *MyLogger) WithMaxQueueDepth(depth int) *MyLogger {
	ml.queue.maxSize = depth
	return ml
}

func (ml *MyLogger) WithQueueFlushInterval(interval time.Duration) *MyLogger {
	ml.queue.defaultFlushInterval = interval
	return ml
}
