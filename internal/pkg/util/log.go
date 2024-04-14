package util

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"os"
)

func SetupLog() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	l := zerolog.New(os.Stdout).With().Timestamp().Caller().Stack().Logger()
	log.Logger = l
	zerolog.DefaultContextLogger = &l
}
