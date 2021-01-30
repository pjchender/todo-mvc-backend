package logger

import (
	"os"
	"regexp"
)

// logger package is used with logrus to show logger status, e.g., error, warning, on Google Cloud
// https://medium.com/p/239baf3b1ac2/
type LogWriter struct {
	levelRegex *regexp.Regexp
}

const (
	LevelError   = "error"
	LevelWarning = "warning"
	LevelFatal   = "fatal"
	LevelPanic   = "panic"
)

func NewLogWriter() *LogWriter {
	levelRegex := regexp.MustCompile("level=([a-z]+)")
	return &LogWriter{
		levelRegex: levelRegex,
	}
}

func (l *LogWriter) toLogLevel(p []byte) string {
	level := l.levelRegex.FindString(string(p))
	if level != "" {
		return level
	}
	return ""
}

func (l *LogWriter) Write(p []byte) (int, error) {
	level := l.toLogLevel(p)
	isAlert := containsString([]string{LevelError, LevelWarning, LevelFatal, LevelPanic}, level)
	if isAlert {
		return os.Stderr.Write(p)
	}
	return os.Stdout.Write(p)
}

func containsString(elements []string, target string) bool {
	for _, str := range elements {
		if str == target {
			return true
		}
	}
	return false
}
