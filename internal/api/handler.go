package api

import (
	"net/http"
	"strconv"
	"technical-test/internal/app"
	"time"

	"github.com/gin-gonic/gin"
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
func (h *Handler) GetDelegations(c *gin.Context) {
	yearStr := c.Query("year")

	if err := validateYear(yearStr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year parameter"})
		return
	}

	delegations, err := h.app.Repo.FindAll(yearStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	c.JSON(http.StatusOK, response)
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
