package app

import "github.com/iktkhor/user-subscriptions-api/internal/app/handler"

type App struct {
	h *handler.Handler
}

func New() *App {
	handler := handler.New()

	return &App{
		h: handler,
	}
}