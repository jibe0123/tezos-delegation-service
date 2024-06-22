package api

import (
	"net/http"
	"technical-test/internal/app"
)

// NewRouter creates a new HTTP router.
func NewRouter(app *app.App) http.Handler {
	mux := http.NewServeMux()
	handler := NewHandler(app)
	mux.HandleFunc("/xtz/delegations", handler.GetDelegations)
	return mux
}
