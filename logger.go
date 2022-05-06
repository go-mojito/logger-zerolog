package zerolog

import (
	"github.com/go-mojito/mojito/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger is a wrapper around zerolog.Logger
type Logger struct {
	fields logger.Fields
}

// NewLogger returns a new Logger
func NewLogger() Logger {
	return Logger{
		fields: make(logger.Fields),
	}
}

// NewLogger returns a new Logger
func NewLoggerWithFields(fields logger.Fields) Logger {
	return Logger{
		fields: fields,
	}
}

// Debug will write a debug log
func (l Logger) Debug(msg interface{}) {
	log := log.Debug()
	log = l.addFields(log)
	log.Msg(msg.(string))
}

// Debugf will write a debug log sprintf-style
func (l Logger) Debugf(msg string, values ...interface{}) {
	log := log.Debug()
	log = l.addFields(log)
	log.Msgf(msg, values...)
}

// Error will write a error log
func (l Logger) Error(msg interface{}) {
	log := log.Error()
	log = l.addFields(log)
	log.Msg(msg.(string))
}

// Errorf will write a error log sprintf-style
func (l Logger) Errorf(msg string, values ...interface{}) {
	log := log.Error()
	log = l.addFields(log)
	log.Msgf(msg, values...)
}

// Fatal will write a fatal log
func (l Logger) Fatal(msg interface{}) {
	log := log.Fatal()
	log = l.addFields(log)
	log.Msg(msg.(string))
}

// Fatalf will write a fatal log sprintf-style
func (l Logger) Fatalf(msg string, values ...interface{}) {
	log := log.Fatal()
	log = l.addFields(log)
	log.Msgf(msg, values...)
}

// Field will add a field to a new logger and return it
func (l Logger) Field(name string, val interface{}) logger.Logger {
	f := l.fields.Clone()
	f[name] = val
	return l
}

// Fields will add multiple fields to a new logger and return it
func (l Logger) Fields(fields logger.Fields) logger.Logger {
	f := l.fields.Clone()
	for name, val := range fields {
		f[name] = val
	}
	return NewLoggerWithFields(fields)
}

// Info will write a info log
func (l Logger) Info(msg interface{}) {
	log := log.Info()
	log = l.addFields(log)
	log.Msg(msg.(string))
}

// Infof will write a info log sprintf-style
func (l Logger) Infof(msg string, values ...interface{}) {
	log := log.Info()
	log = l.addFields(log)
	log.Msgf(msg, values...)
}

// Trace will write a trace log
func (l Logger) Trace(msg interface{}) {
	log := log.Trace()
	log = l.addFields(log)
	log.Msg(msg.(string))
}

// Tracef will write a trace log sprintf-style
func (l Logger) Tracef(msg string, values ...interface{}) {
	log := log.Trace()
	log = l.addFields(log)
	log.Msgf(msg, values...)
}

// Warn will write a warn log
func (l Logger) Warn(msg interface{}) {
	log := log.Warn()
	log = l.addFields(log)
	log.Msg(msg.(string))
}

// Warnf will write a warn log sprintf-style
func (l Logger) Warnf(msg string, values ...interface{}) {
	log := log.Warn()
	log = l.addFields(log)
	log.Msgf(msg, values...)
}

func (l Logger) addFields(log *zerolog.Event) *zerolog.Event {
	for field := range l.fields {
		log = log.Interface(field, l.fields[field])
	}
	return log
}
