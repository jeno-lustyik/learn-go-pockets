package pocketlog

// Logger is used to log information
type Logger struct {
	threshold Level

	// fill later
}

func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}
