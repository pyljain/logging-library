package mylogger

type LogLevel int

const (
	LogLevelPanic   LogLevel = 50
	LogLevelError   LogLevel = 40
	LogLevelWarning LogLevel = 30
	LogLevelInfo    LogLevel = 20
	LogLevelDebug   LogLevel = 10
)

func (l LogLevel) String() string {
	switch l {
	case LogLevelPanic:
		return "Panic"
	case LogLevelError:
		return "Error"
	case LogLevelWarning:
		return "Warning"
	case LogLevelInfo:
		return "Info"
	case LogLevelDebug:
		return "Debug"
	default:
		return "Unknown"
	}
}
