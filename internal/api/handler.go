package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"technical-test/internal/domain"
	"technical-test/internal/service"
)

// Handler handles HTTP requests for delegations
type Handler struct {
	svc *service.DelegationService
}
type jsonResponse struct {
	Data []domain.Delegation `json:"data"`
}

// NewHandler creates a new Handler instance
func NewHandler(svc *service.DelegationService) *Handler {
	return &Handler{svc: svc}
}

// GetDelegations handles requests to retrieve delegations
func (h *Handler) GetDelegations(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	var year int
	var err error
	if yearStr != "" {
		year, err = strconv.Atoi(yearStr)
		if err != nil {
			http.Error(w, "Invalid year parameter", http.StatusBadRequest)
			return
		}
	}

	delegations, err := h.svc.GetDelegations(year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := jsonResponse{
		Data: delegations,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
