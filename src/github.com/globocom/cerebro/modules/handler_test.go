package modules

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"
)

type MockedPersistenceClient struct {
	mock.Mock
}

func (mock *MockedPersistenceClient) GetUser(user string) *User {
	return &User{Segments: []string{"Female"}}
}
func (mock *MockedPersistenceClient) EmptyUser() *User {
	return &User{Segments: make([]string, 0)}
}
func (mock *MockedPersistenceClient) Close() {
	return
}

func TestHealthcheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/healthcheck", strings.NewReader("{}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	settings, _ := LoadSettings()
	client := new(MockedPersistenceClient)

	NewHTTPHandler(settings, client).Healthcheck(c)
	expectedCode := 200
	gotCode := rec.Code
	expectedBody := `{"status":"WORKING"}`
	gotBody := rec.Body.String()

	if expectedCode != gotCode {
		t.Errorf("Healthcheck should always be 200 when application is up.")
	}
	if gotBody != expectedBody {
		t.Errorf("Expected Status Working, but was: %s", rec.Body.String())
	}
}

func TestIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("{}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	settings, _ := LoadSettings()
	client := new(MockedPersistenceClient)

	NewHTTPHandler(settings, client).Index(c)
	expectedCode := 200
	gotCode := rec.Code

	if expectedCode != gotCode {
		t.Errorf("Index should always be 200 when application is up.")
	}
}

func TestGetAttribute(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/attribute", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	settings, _ := LoadSettings()
	client := new(MockedPersistenceClient)

	NewHTTPHandler(settings, client).GetAttribute(c)

	fmt.Printf("teste - %s\n", rec.Body.String())
	var attribute Attribute
	json.NewDecoder(rec.Body).Decode(&attribute)

	expectedCode := 200
	gotCode := rec.Code
	expectedName := "age"
	gotName := attribute.Name
	expectedType := "string"
	gotType := attribute.Type

	if expectedCode != gotCode {
		t.Errorf("Healthcheck should always be 200 when application is up.")
	}
	if gotName != expectedName {
		t.Errorf("Expected name age, but was: %s", gotName)
	}
	if gotType != expectedType {
		t.Errorf("Expected type string, but was: %s", gotType)
	}
}
