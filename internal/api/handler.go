package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"technical-test/internal/app"
	"time"
)

// Handler represents the HTTP handler for the API.
type Handler struct {
	app *app.App
}

// NewHandler creates a new Handler instance.
func NewHandler(app *app.App) *Handler {
	return &Handler{
		app: app,
	}
}

// GetDelegationsResponse represents the response for the GetDelegations handler.
type GetDelegationsResponse struct {
	Data []DelegationResponse `json:"data"`
}

// DelegationResponse represents a single delegation in the response.
type DelegationResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Amount    string    `json:"amount"`
	Delegator string    `json:"delegator"`
	Level     string    `json:"level"`
}

// GetDelegations handles requests to retrieve delegations.
func (h *Handler) GetDelegations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	yearStr := r.URL.Query().Get("year")

	if err := validateYear(yearStr); err != nil {
		http.Error(w, "Invalid year parameter", http.StatusBadRequest)
		return
	}

	delegations, err := h.app.Repo.FindAll(yearStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := GetDelegationsResponse{
		Data: []DelegationResponse{}, // Init with an empty array
	}
	for _, d := range delegations {
		response.Data = append(response.Data, DelegationResponse{
			Timestamp: d.Timestamp,
			Amount:    strconv.FormatInt(d.Amount, 10),
			Delegator: d.Sender.Address,
			Level:     strconv.FormatInt(int64(d.Level), 10),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// validateYear validates the year parameter.
func validateYear(yearStr string) error {
	if yearStr == "" {
		return nil
	}
	if _, err := strconv.Atoi(yearStr); err != nil {
		return err
	}
	return nil
}
