package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"technical-test/internal/app"
)

// Handler defines the HTTP handler for the API.
type Handler struct {
	app *app.App
}

// NewHandler creates a new HTTP handler.
func NewHandler(app *app.App) *Handler {
	return &Handler{app: app}
}

// GetDelegationsResponse represents the response structure for the GetDelegations endpoint
type GetDelegationsResponse struct {
	Data []DelegationResponse `json:"data"`
}

// DelegationResponse represents the structure of a delegation in the response
type DelegationResponse struct {
	Timestamp string `json:"timestamp"`
	Amount    string `json:"amount"`
	Delegator string `json:"delegator"`
	Level     string `json:"level"`
}

// GetDelegations handles requests to retrieve delegations
func (h *Handler) GetDelegations(w http.ResponseWriter, r *http.Request) {
	yearStr := r.URL.Query().Get("year")
	delegations, err := h.app.Repo.FindAll(yearStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response GetDelegationsResponse
	for _, d := range delegations {
		response.Data = append(response.Data, DelegationResponse{
			Timestamp: d.Timestamp, // Keep as string
			Amount:    strconv.FormatInt(d.Amount, 10),
			Delegator: d.Sender.Address,
			Level:     strconv.FormatInt(d.Level, 10),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
