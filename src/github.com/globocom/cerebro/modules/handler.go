package modules

import (
	"net/http"

	"github.com/labstack/echo"
)

// HTTPHandler base handler
type HTTPHandler struct {
	config Settings
	client PersistenceClient
}

// GET /
func (h *HTTPHandler) Index(context echo.Context) error {
	return context.JSON(http.StatusOK, Version{Version: VERSION})
}

// GET /healthcheck
func (h *HTTPHandler) Healthcheck(context echo.Context) error {
	return context.JSON(http.StatusOK, Healthcheck{Status: "WORKING"})
}

// NewHTTPHandler initializes handle object
func NewHTTPHandler(config Settings, client PersistenceClient) *HTTPHandler {
	return &HTTPHandler{
		config: config,
		client: client,
	}
}
