package handler

import "net/http"

func (h *Handler) Router() http.Handler {
	router := http.NewServeMux()

	return router
}