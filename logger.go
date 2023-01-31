package zerolog

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-mojito/mojito/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger is a wrapper around zerolog.Logger
type Logger struct {
	logger zerolog.Logger
	fields logger.Fields
}

// NewLogger returns a new Logger
func NewLogger() *Logger {
	return &Logger{
		logger: log.Logger,
		fields: make(logger.Fields),
	}
}

// NewLogger returns a new Logger
func NewLoggerWithFields(logger zerolog.Logger, fields logger.Fields) *Logger {
	return &Logger{
		logger: logger,
		fields: fields,
	}
}

// SetOutput sets the output destination for the default logger.
func (l *Logger) SetOutput(w io.Writer) error {
	logger := l.logger.Output(w)
	l.logger = logger
	return nil
}

// Field will add a field to a new logger and return it
func (l *Logger) Field(name string, val interface{}) logger.Logger {
	f := l.fields.Clone()
	f[name] = val
	return NewLoggerWithFields(l.logger, f)
}

// Fields will add multiple fields to a new logger and return it
func (l *Logger) Fields(fields logger.Fields) logger.Logger {
	f := l.fields.Clone()
	for name, val := range fields {
		f[name] = val
	}
	return NewLoggerWithFields(l.logger, f)
}

// Trace will write a trace log
func (l *Logger) Trace(msg interface{}) {
	log := l.logger.Trace()
	log = l.addFields(log)
	log.Msg(fmt.Sprintf("%s", msg))
}

// Tracef will write a trace log sprintf-style
func (l *Logger) Tracef(msg string, values ...interface{}) {
	log := l.logger.Trace()
	log = l.addFields(log)
	log.Msgf(msg, values...)
}

// Debug will write a debug log
func (l *Logger) Debug(msg interface{}) {
	log := l.logger.Debug()
	log = l.addFields(log)
	log.Msg(fmt.Sprintf("%s", msg))
}

// Debugf will write a debug log sprintf-style
func (l *Logger) Debugf(msg string, values ...interface{}) {
	log := l.logger.Debug()
	log = l.addFields(log)
	log.Msgf(msg, values...)
}

// Info will write a info log
func (l *Logger) Info(msg interface{}) {
	log := l.logger.Info()
	log = l.addFields(log)
	log.Msg(fmt.Sprintf("%s", msg))
}

// Infof will write a info log sprintf-style
func (l *Logger) Infof(msg string, values ...interface{}) {
	log := l.logger.Info()
	log = l.addFields(log)
	log.Msgf(msg, values...)
}

// Warn will write a warn log
func (l *Logger) Warn(msg interface{}) {
	log := l.logger.Warn()
	log = l.addFields(log)
	log.Msg(fmt.Sprintf("%s", msg))
}

// Warnf will write a warn log sprintf-style
func (l *Logger) Warnf(msg string, values ...interface{}) {
	log := l.logger.Warn()
	log = l.addFields(log)
	log.Msgf(msg, values...)
}

// Error will write a error log
func (l *Logger) Error(msg interface{}) {
	log := l.logger.Error()
	log = l.addFields(log)
	log.Msg(fmt.Sprintf("%s", msg))
}

// Errorf will write a error log sprintf-style
func (l *Logger) Errorf(msg string, values ...interface{}) {
	log := l.logger.Error()
	log = l.addFields(log)
	log.Msgf(msg, values...)
}

// Fatal will write a fatal log
func (l *Logger) Fatal(msg interface{}) {
	log := l.logger.Fatal()
	log = l.addFields(log)
	log.Msg(fmt.Sprintf("%s", msg))
}

// Fatalf will write a fatal log sprintf-style
func (l *Logger) Fatalf(msg string, values ...interface{}) {
	log := l.logger.Fatal()
	log = l.addFields(log)
	log.Msgf(msg, values...)
}

func (l Logger) addFields(log *zerolog.Event) *zerolog.Event {
	for field := range l.fields {
		sanitizedVal := strings.ReplaceAll(fmt.Sprint(l.fields[field]), "\n", "")
		sanitizedVal = strings.ReplaceAll(sanitizedVal, "\r", "")
		log = log.Str(field, fmt.Sprint(sanitizedVal))
	}
	return log
}
