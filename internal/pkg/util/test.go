package util

import (
	"context"
	"github.com/rs/zerolog/log"
	"os"
	"testing"
)

// GoToProjectDir move to root folder. Must only use in tests. Don't use in production
func GoToProjectDir() {
	if err := os.Chdir(GetProjectRoot()); err != nil {
		panic(err)
	}
}

func NewTestCtx(t testing.TB) context.Context {
	ctx := context.Background()
	return log.Ctx(ctx).With().Str("test_name", t.Name()).Logger().WithContext(ctx)
}
