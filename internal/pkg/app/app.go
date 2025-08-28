package app

import (
	"log"
	"net/http"

	"github.com/iktkhor/user-subscriptions-api/internal/app/handler"
)

type App struct {
	h *handler.Handler
}

func New() *App {
	handler := handler.New()

	return &App{
		h: handler,
	}
}


func (a *App) Run() error {
	router := a.h.Router()

	//addr := fmt.Sprintf("%s:%d", a.cfg.Host, a.cfg.Port)
    server := http.Server{
        Addr:    ":8080",
        Handler: router,
    }

	//fmt.Printf("Server listening on %s in %s mode\n", addr, a.cfg.Env)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal()
	}

	return nil
}
