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
	logger zerolog.Context
}

// NewLogger returns a new Logger
func NewLogger() *Logger {
	return &Logger{
		logger: log.Logger.With(),
	}
}

// SetOutput sets the output destination for the current logger.
func (l *Logger) SetOutput(w io.Writer) error {
	logger := l.logger.Logger().Output(w).With()
	l.logger = logger
	return nil
}

// Field will add a field to a new logger and return it
func (l *Logger) Field(name string, val interface{}) logger.Logger {
	return &Logger{
		logger: l.logger.Interface(name, val),
	}
}

// Fields will add multiple fields to a new logger and return it
func (l *Logger) Fields(fields logger.Fields) logger.Logger {
	return &Logger{
		logger: l.logger.Fields(map[string]interface{}(fields)),
	}
}

// Trace will write a trace log
func (l *Logger) Trace(msg interface{}) {
	l.toLogger().Trace().Msg(fmt.Sprintf("%s", msg))
}

// Tracef will write a trace log sprintf-style
func (l *Logger) Tracef(msg string, values ...interface{}) {
	l.toLogger().Trace().Msgf(msg, values...)
}

// Debug will write a debug log
func (l *Logger) Debug(msg interface{}) {
	l.toLogger().Debug().Msg(fmt.Sprintf("%s", msg))
}

// Debugf will write a debug log sprintf-style
func (l *Logger) Debugf(msg string, values ...interface{}) {
	l.toLogger().Debug().Msgf(msg, values...)
}

// Info will write a info log
func (l *Logger) Info(msg interface{}) {
	l.toLogger().Info().Msg(fmt.Sprintf("%s", msg))
}

// Infof will write a info log sprintf-style
func (l *Logger) Infof(msg string, values ...interface{}) {
	l.toLogger().Info().Msgf(msg, values...)
}

// Warn will write a warn log
func (l *Logger) Warn(msg interface{}) {
	l.toLogger().Warn().Msg(fmt.Sprintf("%s", msg))
}

// Warnf will write a warn log sprintf-style
func (l *Logger) Warnf(msg string, values ...interface{}) {
	l.toLogger().Warn().Msgf(msg, values...)
}

// Error will write a error log
func (l *Logger) Error(msg interface{}) {
	l.toLogger().Error().Msg(fmt.Sprintf("%s", msg))
}

// Errorf will write a error log sprintf-style
func (l *Logger) Errorf(msg string, values ...interface{}) {
	l.toLogger().Error().Msgf(msg, values...)
}

// Fatal will write a fatal log
func (l *Logger) Fatal(msg interface{}) {
	l.toLogger().Fatal().Msg(fmt.Sprintf("%s", msg))
}

// Fatalf will write a fatal log sprintf-style
func (l *Logger) Fatalf(msg string, values ...interface{}) {
	l.toLogger().Fatal().Msgf(msg, values...)
}

func (l *Logger) toLogger() *zerolog.Logger {
	logger := l.logger.Logger()
	return &logger
}
