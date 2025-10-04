package router

import (
	"net/http"

	"github.com/alonsofritz/tt-shopee/internal/api/handler"
	"github.com/alonsofritz/tt-shopee/internal/service"
)

func SetupRouter(ticketService *service.TicketService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.HealthCheckHandler)

	mux.HandleFunc("/tickets", handler.NewTicketHandler(ticketService).ProcessTicketHandler)

	return mux
}
