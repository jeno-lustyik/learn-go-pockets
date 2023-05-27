package pocketlog

import "fmt"

// Level represents an available logging level
type Level byte

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

func (l *Logger) Debugf(format string, args ...any) {

	_, _ = fmt.Fprintf(l.output, format+"\n", args...)

}

func (l *Logger) Infof(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}
