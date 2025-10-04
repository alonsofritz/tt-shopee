package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alonsofritz/tt-shopee/internal/dto"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "apllication/json")
	w.WriteHeader(http.StatusOK)

	response := dto.HealthResponseDto{
		Status:  "OK",
		Message: "Server UP!",
	}

	_ = json.NewEncoder(w).Encode(response)
}
