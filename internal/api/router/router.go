package router

import (
	"net/http"

	"github.com/alonsofritz/tt-shopee/internal/api/handler"
)

func SetupRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.HealthCheckHandler)

	return mux
}
