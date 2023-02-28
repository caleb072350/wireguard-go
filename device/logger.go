package device

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Logger struct {
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
}

// Log levels for use with NewLogger.
const (
	LogLevelSilent = iota
	LogLevelError
	LogLevelInfo
	LogLevelDebug
)

// Function for use in Logger for discarding logged lines.
func DiscardLogf(format string, args ...any) {}

// NewLogger constructs a Logger that writes to stdout.
// It logs at the specified log level and above.
// It decorates log lines with the log level, date, time and prepend.
func NewLogger(level int, prepend string) *Logger {
	output := os.Stdout
	logger := new(Logger)

	logErr, logInfo, logDebug := func() (io.Writer, io.Writer, io.Writer) {
		if level >= LogLevelDebug {
			return output, output, output
		}
		if level >= LogLevelInfo {
			return output, output, ioutil.Discard
		}
		if level >= LogLevelError {
			return output, ioutil.Discard, ioutil.Discard
		}
		return ioutil.Discard, ioutil.Discard, ioutil.Discard
	}()

	logger.Debug = log.New(logDebug,
		"DEBUG: "+prepend,
		log.Ldate|log.Ltime,
	)
	logger.Info = log.New(logInfo,
		"INFO: "+prepend,
		log.Ldate|log.Ltime,
	)
	logger.Error = log.New(logErr,
		"ERROR: "+prepend,
		log.Ldate|log.Ltime,
	)
	return logger
}
