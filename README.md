# MyLogger

MyLogger is a flexible, queue-based logging library for Go applications that supports multiple output formats and destinations.

## Features

- Multiple output formats (JSON and plain text)
- Configurable log destinations (stdout and file)
- Queue-based logging with configurable batch size and flush intervals
- Multiple log levels (Debug, Info, Warning, Error, Panic)
- Support for structured logging with key-value pairs
- Thread-safe implementation

## Installation

```bash
go get github.com/pyljain/logging-library
```

## Quick Start

```go
package main

import (
    "github.com/pyljain/logging-library/pkg/mylogger"
    "time"
)

func main() {
    // Create a new logger with string format and stdout output
    logger := mylogger.NewMyLogger().
        WithString().
        WithStdIO()

    // Log with additional context
    logger.Info("hello", mylogger.WithLabel("app", "myapp"))
    
    // Simple log message
    logger.Debug("debug message")
}
```

## Configuration Options

### Output Format

```go
// JSON format
logger.WithJSON()

// String format (default)
logger.WithString()
```

### Output Destination

```go
// Write to stdout (default)
logger.WithStdIO()

// Write to file
logger.WithLogFile("path/to/logfile.log")
```

### Log Levels

Available log levels (from lowest to highest):
- Debug (10)
- Info (20)
- Warning (30)
- Error (40)
- Panic (50)

Set minimum log level:
```go
logger.WithMinLogLevel(mylogger.LogLevelDebug)
```

### Queue Configuration

```go
// Set maximum queue size before flush
logger.WithMaxQueueDepth(1000)

// Set queue flush interval
logger.WithQueueFlushInterval(5 * time.Second)
```

## Log Methods

```go
// Basic logging
logger.Debug("debug message")
logger.Info("info message")
logger.Warning("warning message")
logger.Error("error message")
logger.Panic("panic message")

// Logging with context
logger.Info("user logged in", 
    mylogger.WithLabel("user_id", "123"),
    mylogger.WithLabel("ip", "192.168.1.1"),
)
```

## Output Examples

### String Format
```
2024-03-14 10:15:30    Info    hello    app:myapp
```

### JSON Format
```json
{
    "timestamp": "2024-03-14T10:15:30Z",
    "level": "Info",
    "message": "hello",
    "elements": {
        "app": "myapp"
    }
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT
