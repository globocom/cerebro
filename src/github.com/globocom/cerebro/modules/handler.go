package modules

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
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
	return context.JSON(http.StatusOK, Status{Status: "WORKING"})
}

// GET /attribute/attributeName
func (h *HTTPHandler) GetAttribute(context echo.Context) error {
	name := context.Param("name")
	a, err := h.client.GetAttribute(name)
	if err != nil {
		log.Warn(err)
		return context.JSON(http.StatusNotFound, Status{Status: "MISS", Err: err.Error()})
	}
	return context.JSON(http.StatusOK, Attribute{Name: a.Name, Type: a.Type})
}

// POST /attribute
func (h *HTTPHandler) PostAttribute(context echo.Context) error {
	decoder := json.NewDecoder(context.Request().Body)
	var a Attribute
	err := decoder.Decode(&a)
	if err != nil {
		log.Warn(err)
		return context.JSON(http.StatusNotFound, Status{Status: "MISS", Err: err.Error()})
	}
	err = h.client.AddAttribute(a.Name, a.Type)
	if err != nil {
		log.Warn(err)
		return context.JSON(http.StatusNotFound, Status{Status: "MISS", Err: err.Error()})
	}
	return context.JSON(http.StatusOK, Status{Status: "CREATED"})
}

// DELETE /attribute/attributeName
func (h *HTTPHandler) DeleteAttribute(context echo.Context) error {
	name := context.Param("name")
	err := h.client.DeleteAttribute(name)
	if err != nil {
		log.Warn(err)
		return context.JSON(http.StatusNotFound, Status{Status: "MISS", Err: err.Error()})
	}
	return context.JSON(http.StatusOK, Status{Status: "DELETED"})
}

// PUT /attribute/attributeName
func (h *HTTPHandler) UpdateAttribute(context echo.Context) error {
	name := context.Param("name")
	var a Attribute
	a.Name = name

	decoder := json.NewDecoder(context.Request().Body)
	err := decoder.Decode(&a)
	if err != nil {
		log.Warn(err)
		return context.JSON(http.StatusNotFound, Status{Status: "MISS", Err: err.Error()})
	}

	err = h.client.UpdateAttribute(a.Name, a.Type)
	if err != nil {
		log.Warn(err)
		return context.JSON(http.StatusNotFound, Status{Status: "MISS", Err: err.Error()})
	}
	return context.JSON(http.StatusOK, Status{Status: "UPDATED"})
}

// NewHTTPHandler initializes handle object
func NewHTTPHandler(config Settings, client PersistenceClient) *HTTPHandler {
	return &HTTPHandler{
		config: config,
		client: client,
	}
}
