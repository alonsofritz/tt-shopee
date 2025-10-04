package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alonsofritz/tt-shopee/internal/domain/model"
	"github.com/alonsofritz/tt-shopee/internal/dto"
	"github.com/alonsofritz/tt-shopee/internal/service"
)

type TicketHandler struct {
	ticketService *service.TicketService
}

func NewTicketHandler(ticketService *service.TicketService) *TicketHandler {
	return &TicketHandler{
		ticketService: ticketService,
	}
}

func (h *TicketHandler) ProcessTicketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var payload dto.TicketRequestDTO
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var results []map[string]interface{}

	ticket := model.Ticket{
		ShowID: payload.ShowID,
		UserID: payload.UserID,
	}

	err := h.ticketService.ProcessTicket(ticket)
	if err != nil {
		results = append(results, map[string]interface{}{
			"show_id": payload.ShowID,
			"user_id": payload.UserID,
			"status":  "error",
			"error":   err.Error(),
		})
	} else {
		results = append(results, map[string]interface{}{
			"show_id": payload.ShowID,
			"user_id": payload.UserID,
			"status":  "ok",
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(results)
}
