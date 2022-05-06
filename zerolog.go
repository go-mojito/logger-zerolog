package zerolog

import (
	"os"

	"github.com/go-mojito/mojito"
	zlog "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// AsDefault registers zerolog as the default logger
func AsDefault() {
	err := mojito.Register(func() mojito.Logger {
		return NewLogger()
	}, true)
	if err != nil {
		panic(err)
	}
}

// As registers zerolog under a given name
func As(name string) {
	err := mojito.RegisterNamed(name, func() mojito.Logger {
		return NewLogger()
	}, true)
	if err != nil {
		panic(err)
	}
}

func Pretty() {
	log.Logger = log.Output(zlog.ConsoleWriter{Out: os.Stderr})
}
