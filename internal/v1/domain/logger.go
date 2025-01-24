package domain

import "context"

type Logger interface {
	Ctx(ctx context.Context) Logger
	Debug(args ...any) Logger
	Info(args ...any) Logger
	Warn(args ...any) Logger
	Error(args ...any) Logger
	Fatal(args ...any) Logger
}
