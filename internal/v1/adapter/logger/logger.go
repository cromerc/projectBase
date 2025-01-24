package logger

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/cromerc/projectBase/internal/v1/domain"

	"github.com/cromerc/projectBase/internal/v1/service"
	"github.com/rs/zerolog"
)

type logger struct {
	log *zerolog.Logger
	ctx context.Context
}

// Ctx will set the context used by the logger.
func (l *logger) Ctx(ctx context.Context) domain.Logger {
	l.ctx = ctx
	return l
}

// Debug will log info useful for debugging the program as it advances.
func (l *logger) Debug(args ...any) domain.Logger {
	for _, arg := range args {
		l.log.Debug().Ctx(l.ctx).Msg(fmt.Sprint(arg))
	}
	return l
}

// Info will log general information as the program advances.
func (l *logger) Info(args ...any) domain.Logger {
	for _, arg := range args {
		l.log.Info().Ctx(l.ctx).Msg(fmt.Sprintf("%+v", arg))
	}
	return l
}

// Warn will log warnings that can be recovered from.
func (l *logger) Warn(args ...any) domain.Logger {
	for _, arg := range args {
		l.log.Warn().Ctx(l.ctx).Msg(fmt.Sprint(arg))
	}
	return l
}

// Error will log sever errors.
func (l *logger) Error(args ...any) domain.Logger {
	for _, arg := range args {
		l.log.Error().Ctx(l.ctx).Msg(fmt.Sprint(arg))
	}
	return l
}

// Fatal will log the messages and then end the process.
func (l *logger) Fatal(args ...any) domain.Logger {
	for i, arg := range args {
		// The last arg should call fatal since it will kill the process
		if (len(args) - 1) == i {
			l.log.Fatal().Ctx(l.ctx).Msg(fmt.Sprint(arg))
		} else {
			l.log.Error().Ctx(l.ctx).Msg(fmt.Sprint(arg))
		}
	}
	return l
}

// New returns a configured logging instance to be used
func New(logConfig service.Log) domain.Logger {
	var logLevel zerolog.Level
	switch logConfig.Level {
	case "debug":
		logLevel = zerolog.DebugLevel
	case "info":
		logLevel = zerolog.InfoLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	case "none":
		logLevel = zerolog.NoLevel
	default:
		logLevel = zerolog.InfoLevel
	}

	zl := zerolog.New(os.Stdout).With().Timestamp().Logger()

	if logConfig.Pretty {
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano}
		zl = zerolog.New(output).With().Timestamp().Logger()
	}

	zerolog.SetGlobalLevel(logLevel)
	zerolog.TimeFieldFormat = time.RFC3339Nano

	l := logger{
		log: &zl,
		ctx: context.Background(),
	}

	return &l
}
