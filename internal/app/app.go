package app

import "github.com/fernandogiovanini/backhome/internal/logger"

type App struct{}

func New() *App {
	return &App{}
}

// Fatal calls logger.Fatal but it should
// call a.logger.Fatal() when we do DI on New
func (a *App) Fatal(msg string, args ...any) {
	logger.Fatalf(msg, args...)
}
