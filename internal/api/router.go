package api

import (
	"github.com/gin-gonic/gin"
	"technical-test/internal/app"
)

func NewRouter(app *app.App) *gin.Engine {
	router := gin.Default()
	handler := NewHandler(app)
	router.GET("/xtz/delegations", handler.GetDelegations)
	return router
}
