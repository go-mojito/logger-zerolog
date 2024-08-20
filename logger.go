package zerolog

import (
	"io"

	"github.com/go-mojito/mojito/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger is a wrapper around zerolog.Logger
type Logger struct {
	Logger zerolog.Logger
	level  zerolog.Level
}

// NewLogger returns a new Logger
func NewLogger() *Logger {
	return &Logger{
		Logger: log.Logger,
		level:  zerolog.NoLevel,
	}
}

// SetOutput sets the output destination for the default Logger.
func (l *Logger) SetOutput(w io.Writer) error {
	Logger := l.Logger.Output(w)
	l.Logger = Logger
	return nil
}

// Field will add a field to a new Logger and return it
func (l *Logger) Field(name string, val interface{}) logger.Logger {
	return &Logger{
		Logger: l.Logger.With().Interface(name, val).Logger(),
		level:  l.level,
	}
}

// Fields will add multiple fields to a new Logger and return it
func (l *Logger) Fields(fields logger.Fields) logger.Logger {
	return &Logger{
		Logger: l.Logger.With().Fields(fields).Logger(),
		level:  l.level,
	}
}

// Trace will write a trace log
func (l *Logger) Trace() logger.Logger {
	return &Logger{
		Logger: l.Logger,
		level:  zerolog.TraceLevel,
	}
}

// Debug will write a debug log
func (l *Logger) Debug() logger.Logger {
	return &Logger{
		Logger: l.Logger,
		level:  zerolog.DebugLevel,
	}
}

// Info will write a info log
func (l *Logger) Info() logger.Logger {
	return &Logger{
		Logger: l.Logger,
		level:  zerolog.InfoLevel,
	}
}

// Warn will write a warn log
func (l *Logger) Warn() logger.Logger {
	return &Logger{
		Logger: l.Logger,
		level:  zerolog.WarnLevel,
	}
}

// Error will write a error log
func (l *Logger) Error() logger.Logger {
	return &Logger{
		Logger: l.Logger,
		level:  zerolog.ErrorLevel,
	}
}

// Fatal will write a fatal log
func (l *Logger) Fatal() logger.Logger {
	return &Logger{
		Logger: l.Logger,
		level:  zerolog.FatalLevel,
	}
}

// Msg implements logger.Logger.
func (l *Logger) Msg(msg interface{}) {
	l.Msgf("%v", msg)
}

// Msgf implements logger.Logger.
func (l *Logger) Msgf(msg string, values ...interface{}) {
	switch l.level {
	case zerolog.TraceLevel:
		l.Logger.Trace().Msgf(msg, values...)
	case zerolog.DebugLevel:
		l.Logger.Debug().Msgf(msg, values...)
	case zerolog.InfoLevel:
		l.Logger.Info().Msgf(msg, values...)
	case zerolog.WarnLevel:
		l.Logger.Warn().Msgf(msg, values...)
	case zerolog.ErrorLevel:
		l.Logger.Error().Msgf(msg, values...)
	case zerolog.FatalLevel:
		l.Logger.Fatal().Msgf(msg, values...)
	case zerolog.PanicLevel:
		l.Logger.Panic().Msgf(msg, values...)
	case zerolog.NoLevel:
		l.Logger.Log().Msgf(msg, values...)
	}
}
