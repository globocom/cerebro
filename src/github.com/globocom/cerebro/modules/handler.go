package modules

import (
	"encoding/json"
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

// GET /attribute
func (h *HTTPHandler) GetAllAttributes(context echo.Context) error {
	return context.JSON(http.StatusNotFound, `"status":"RETORNA O NOME DE TODOS OS ATRIBUTOS"`)
}

// GET /attribute/attributeName
func (h *HTTPHandler) GetAttribute(context echo.Context) error {
	name := context.Param("name")
	if name == "fakeAttribute" {
		return context.JSON(http.StatusOK, Attribute{Name: name, Type: "int"})
	}
	return context.JSON(http.StatusNotFound, `"status":"MISS"`)
}

// POST /attribute
func (h *HTTPHandler) PostAttribute(context echo.Context) error {
	decoder := json.NewDecoder(context.Request().Body)
	var a Attribute
	err := decoder.Decode(&a)
	if err != nil {
		panic(err)
	}
	return context.JSON(http.StatusOK, Attribute{Name: a.Name, Type: a.Type})
}

// PUT /attribute
func (h *HTTPHandler) UpdateAttribute(context echo.Context) error {
	decoder := json.NewDecoder(context.Request().Body)
	var a Attribute
	a.Name = "fakeAttribute"
	a.Type = "int"
	err := decoder.Decode(&a)
	if err != nil {
		panic(err)
	}
	return context.JSON(http.StatusOK, Attribute{Name: a.Name, Type: a.Type})
}

// DELETE /attribute/attributeName
func (h *HTTPHandler) DeleteAttribute(context echo.Context) error {
	name := context.Param("name")
	var a Attribute
	a.Name = "fakeAttribute"
	a.Type = "string"
	if a.Name == name {
		a = Attribute{}
	}
	return context.JSON(http.StatusOK, Attribute{Name: a.Name, Type: a.Type})
}

// NewHTTPHandler initializes handle object
func NewHTTPHandler(config Settings, client PersistenceClient) *HTTPHandler {
	return &HTTPHandler{
		config: config,
		client: client,
	}
}
