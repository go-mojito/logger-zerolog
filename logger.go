package zerolog

import (
	"fmt"
	"io"

	"github.com/go-mojito/mojito/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger is a wrapper around zerolog.Logger
type Logger struct {
	logger zerolog.Logger
}

// NewLogger returns a new Logger
func NewLogger() *Logger {
	return &Logger{
		logger: log.Logger.With().Logger(),
	}
}

// SetOutput sets the output destination for the current logger.
func (l *Logger) SetOutput(w io.Writer) error {
	logger := l.logger.Output(w)
	l.logger = logger
	return nil
}

// Field will add a field to a new logger and return it
func (l *Logger) Field(name string, val interface{}) logger.Logger {
	return &Logger{
		logger: l.logger.With().Interface(name, val).Logger(),
	}
}

// Fields will add multiple fields to a new logger and return it
func (l *Logger) Fields(fields logger.Fields) logger.Logger {
	return &Logger{
		logger: l.logger.With().Fields(fields).Logger(),
	}
}

// Trace will write a trace log
func (l *Logger) Trace(msg interface{}) {
	l.logger.Trace().Msg(fmt.Sprintf("%s", msg))
}

// Tracef will write a trace log sprintf-style
func (l *Logger) Tracef(msg string, values ...interface{}) {
	l.logger.Trace().Msgf(msg, values...)
}

// Debug will write a debug log
func (l *Logger) Debug(msg interface{}) {
	l.logger.Debug().Msg(fmt.Sprintf("%s", msg))
}

// Debugf will write a debug log sprintf-style
func (l *Logger) Debugf(msg string, values ...interface{}) {
	l.logger.Debug().Msgf(msg, values...)
}

// Info will write a info log
func (l *Logger) Info(msg interface{}) {
	l.logger.Info().Msg(fmt.Sprintf("%s", msg))
}

// Infof will write a info log sprintf-style
func (l *Logger) Infof(msg string, values ...interface{}) {
	l.logger.Info().Msgf(msg, values...)
}

// Warn will write a warn log
func (l *Logger) Warn(msg interface{}) {
	l.logger.Warn().Msg(fmt.Sprintf("%s", msg))
}

// Warnf will write a warn log sprintf-style
func (l *Logger) Warnf(msg string, values ...interface{}) {
	l.logger.Warn().Msgf(msg, values...)
}

// Error will write a error log
func (l *Logger) Error(msg interface{}) {
	l.logger.Error().Msg(fmt.Sprintf("%s", msg))
}

// Errorf will write a error log sprintf-style
func (l *Logger) Errorf(msg string, values ...interface{}) {
	l.logger.Error().Msgf(msg, values...)
}

// Fatal will write a fatal log
func (l *Logger) Fatal(msg interface{}) {
	l.logger.Fatal().Msg(fmt.Sprintf("%s", msg))
}

// Fatalf will write a fatal log sprintf-style
func (l *Logger) Fatalf(msg string, values ...interface{}) {
	l.logger.Fatal().Msgf(msg, values...)
}
