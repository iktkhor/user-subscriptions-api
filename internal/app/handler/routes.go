package handler

import "net/http"

func Router() http.Handler {
	router := http.NewServeMux()

	return router
}