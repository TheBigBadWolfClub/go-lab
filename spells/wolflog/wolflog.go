package wolflog

import (
	"fmt"
	"io"
	"log"
)

type Level int8

const (
	INFO Level = 1 << iota
	DEBUG
	ERROR
	TRACE
	WARN
)

type Wolflogger interface {
	Debug(format string, v ...interface{})
	Error(format string, v ...interface{})
	Info(format string, v ...interface{})
	Trace(format string, v ...interface{})
	Warn(format string, v ...interface{})
}

type wolflog struct {
	log   *log.Logger
	level Level
}

func New(level Level) *wolflog {
	logger := log.Default()
	logger.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return &wolflog{
		log:   logger,
		level: level,
	}
}

func (w *wolflog) Output(out io.Writer) {
	w.log.SetOutput(out)
}

func (w *wolflog) Level(l Level) {
	w.level = l
}

func (w *wolflog) writef(l Level, format string, v ...interface{}) {
	w.log.SetPrefix(l.String() + ": ")
	_ = w.log.Output(3, fmt.Sprintf(format, v...))
}

func (w *wolflog) Debug(format string, v ...interface{}) {
	w.writef(DEBUG, format, v...)
}

func (w *wolflog) Info(format string, v ...interface{}) {
	w.writef(INFO, format, v...)
}

func (w *wolflog) Warn(format string, v ...interface{}) {
	w.writef(WARN, format, v...)

}

func (w *wolflog) Error(format string, v ...interface{}) {
	w.writef(ERROR, format, v...)
}

func (w *wolflog) Trace(format string, v ...interface{}) {
	w.writef(TRACE, format, v...)
}

func (l Level) String() string {
	switch l {
	case DEBUG:
		return "Debug"
	case ERROR:
		return "Error"
	case INFO:
		return "Info"
	case TRACE:
		return "Trace"
	case WARN:
		return "Warn"
	default:
		return "unknown"
	}
}
