package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
)

type LogLevel uint8

func (l LogLevel) String() string {
	s := ""
	if l&LevelDebug != 0 {
		s += "DEBUG "
	}
	if l&LevelInfo != 0 {
		s += "INFO "
	}
	if l&LevelWarn != 0 {
		s += "WARN "
	}
	if l&LevelError != 0 {
		s += "ERROR "
	}
	if s == "" {
		return "NONE"
	}
	return s[:len(s)-1] // Remove trailing space
}

const (
	LevelDebug LogLevel = 1 << iota
	LevelInfo
	LevelWarn
	LevelError

	LevelAll  = LevelDebug | LevelInfo | LevelWarn | LevelError
	LevelNone = LogLevel(0)
)

type Logger interface {
	SetVerboseLevel(levels LogLevel)
	GetVerboseLevel() LogLevel

	MapPrefixes(prefixes map[LogLevel]string)
	MapStyles(styles map[LogLevel]string)
	MapWriters(writers map[LogLevel]io.Writer)

	Debug(args ...any)
	Debugf(format string, args ...any)
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
}

type logger struct {
	verboseLevels LogLevel
	prefixes      map[LogLevel]string
	styles        map[LogLevel]string
	writers       map[LogLevel]io.Writer
	mu            sync.Mutex
}

func NewLogger() Logger {
	return &logger{
		verboseLevels: LevelNone,
		prefixes:      make(map[LogLevel]string),
		writers:       make(map[LogLevel]io.Writer),
		mu:            sync.Mutex{},
	}
}

func NewLoggerWith(verboseLevels LogLevel, prefixes map[LogLevel]string, styles map[LogLevel]string, writers map[LogLevel]io.Writer) Logger {
	return &logger{
		verboseLevels: verboseLevels,
		prefixes:      prefixes,
		styles:        styles,
		writers:       writers,
		mu:            sync.Mutex{},
	}
}

func (l *logger) SetVerboseLevel(levels LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.verboseLevels = levels
}

func (l *logger) GetVerboseLevel() LogLevel {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.verboseLevels
}

func (l *logger) MapPrefixes(prefixes map[LogLevel]string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.prefixes = prefixes
}

func (l *logger) MapStyles(styles map[LogLevel]string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.styles = styles
}

func (l *logger) MapWriters(writers map[LogLevel]io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.writers = writers
}

func (l *logger) Debug(args ...any) {
	l.print(LevelDebug, args...)
}

func (l *logger) Debugf(format string, args ...any) {
	l.print(LevelDebug, fmt.Sprintf(format, args...))
}

func (l *logger) Info(args ...any) {
	l.print(LevelInfo, args...)
}

func (l *logger) Infof(format string, args ...any) {
	l.print(LevelInfo, fmt.Sprintf(format, args...))
}

func (l *logger) Warn(args ...any) {
	l.print(LevelWarn, args...)
}

func (l *logger) Warnf(format string, args ...any) {
	l.print(LevelWarn, fmt.Sprintf(format, args...))
}

func (l *logger) Error(args ...any) {
	l.print(LevelError, args...)
}

func (l *logger) Errorf(format string, args ...any) {
	l.print(LevelError, fmt.Sprintf(format, args...))
}

func (l *logger) print(level LogLevel, args ...any) {
	if level == LevelNone || (l.verboseLevels&level) == 0 {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	writer := l.getWriter(level)
	prefix := l.getPrefix(level)
	style := l.getStyle(level)
	output := fmt.Sprint(args...)

	if style != "" {
		output = fmt.Sprintf(style, output)
	}

	fmt.Fprintf(writer, "%s%s\n", prefix, output)
}

func (l *logger) getWriter(level LogLevel) io.Writer {
	if writer, exists := l.writers[level]; exists {
		return writer
	}
	switch level {
	case LevelInfo, LevelWarn:
		return os.Stdout

	case LevelError:
		return os.Stderr

	default:
		return io.Discard
	}
}

func (l *logger) getStyle(level LogLevel) string {
	if style, exists := l.styles[level]; exists {
		return style
	}
	return ""
}

func (l *logger) getPrefix(level LogLevel) string {
	if prefix, exists := l.prefixes[level]; exists {
		return prefix
	}
	return ""
}
