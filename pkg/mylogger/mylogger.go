package mylogger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type MyLogger struct {
	destination string
	logFile     string
	format      string
	minLogLevel LogLevel
	queue       *queue
}

func NewMyLogger() *MyLogger {
	l := &MyLogger{
		destination: "stdio",
		// logFile:      "mylogger.log",
		format:      "string",
		minLogLevel: LogLevelInfo,
	}

	l.queue = newQueue(1000, 1*time.Second, l.write)
	go l.queue.run()

	return l
}

type LogElement struct {
	Label string
	Value interface{}
}

func WithLabel(label string, value interface{}) LogElement {
	return LogElement{
		Label: label,
		Value: value,
	}
}

func (ml *MyLogger) Info(message string, elements ...LogElement) {
	ml.Log(LogLevelInfo, message, elements...)
}

func (ml *MyLogger) Debug(message string, elements ...LogElement) {
	ml.Log(LogLevelDebug, message, elements...)
}

func (ml *MyLogger) Error(message string, elements ...LogElement) {
	ml.Log(LogLevelError, message, elements...)
}

func (ml *MyLogger) Panic(message string, elements ...LogElement) {
	ml.Log(LogLevelPanic, message, elements...)
}

func (ml *MyLogger) Warning(message string, elements ...LogElement) {
	ml.Log(LogLevelWarning, message, elements...)
}

func (ml *MyLogger) Log(level LogLevel, message string, elements ...LogElement) {
	if level < ml.minLogLevel {
		return
	}

	elementsMap := map[string]interface{}{}
	for _, e := range elements {
		elementsMap[e.Label] = e.Value
	}

	ml.queue.add(&queueElement{
		Timestamp: time.Now(),
		Level:     level,
		Message:   message,
		Elements:  elementsMap,
	})
}

func (ml *MyLogger) getWriter() io.Writer {
	if ml.destination == "stdio" {
		return os.Stdout
	} else if ml.destination == "file" {
		f, err := os.OpenFile(ml.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("failed to open log file: %v", err)
		}
		return f
	}
	return nil
}

func (ml *MyLogger) write(queueElements []*queueElement) {

	writer := ml.getWriter()

	for _, qe := range queueElements {
		if ml.format == "string" {
			base := fmt.Sprintf("%s\t%s\t%s", qe.Timestamp.GoString(), qe.Level.String(), qe.Message)
			for k, v := range qe.Elements {
				base = fmt.Sprintf("%s\t%s:%v", base, k, v)
			}
			writer.Write([]byte(base + "\n"))
		} else if ml.format == "json" {
			err := jsoniter.NewEncoder(writer).Encode(&qe)
			if err != nil {
				log.Fatalf("failed to encode json: %v", err)
			}
		}
	}
}
