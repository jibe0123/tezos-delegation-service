package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"technical-test/internal/api"
	"technical-test/internal/domain"
	"technical-test/internal/repository"
	"technical-test/internal/service"
	"testing"
	"time"
)

func TestGetDelegations(t *testing.T) {
	repo := repository.NewMemoryRepository()
	svc := service.NewDelegationService(repo)
	handler := api.NewHandler(svc)

	// Seed the repository with test data
	now := time.Now()
	lastYear := now.AddDate(-1, 0, 0)
	svc.SaveDelegation("tz1a1SAaXRt9yoGMx29rh9FsBF4UzmvojdTL", now, 1000, 12345)
	svc.SaveDelegation("KT1JejNYjmQYh8yw95u5kfQDRuxJcaUPjUnf", lastYear, 2000, 67890)

	req, err := http.NewRequest("GET", "/xtz/delegations", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.GetDelegations).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response struct {
		Data []domain.Delegation `json:"data"`
	}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if len(response.Data) != 2 {
		t.Fatalf("expected 2 delegations, got %d", len(response.Data))
	}
}

func TestGetDelegationsByYear(t *testing.T) {
	repo := repository.NewMemoryRepository()
	svc := service.NewDelegationService(repo)
	handler := api.NewHandler(svc)

	// Seed the repository with test data
	now := time.Now()
	lastYear := now.AddDate(-1, 0, 0)
	svc.SaveDelegation("tz1a1SAaXRt9yoGMx29rh9FsBF4UzmvojdTL", now, 1000, 12345)
	svc.SaveDelegation("KT1JejNYjmQYh8yw95u5kfQDRuxJcaUPjUnf", lastYear, 2000, 67890)

	req, err := http.NewRequest("GET", "/xtz/delegations?year="+strconv.Itoa(now.Year()), nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.GetDelegations).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response struct {
		Data []domain.Delegation `json:"data"`
	}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("could not decode response: %v", err)
	}

	if len(response.Data) != 1 {
		t.Fatalf("expected 1 delegation, got %d", len(response.Data))
	}

	if response.Data[0].Timestamp.Year() != now.Year() {
		t.Fatalf("expected year %d, got %d", now.Year(), response.Data[0].Timestamp.Year())
	}
}
