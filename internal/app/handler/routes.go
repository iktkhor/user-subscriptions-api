package handler

import "net/http"

func (h *Handler) NewRouter() http.Handler {
	router := http.NewServeMux()

	return router
}