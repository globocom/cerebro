package modules

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HandlerTestSuite struct {
	suite.Suite
	client *http.Client
}

func (suite *HandlerTestSuite) SetupSuite() {
	go Init()
	suite.client = &http.Client{
		Timeout: 1 * time.Second,
	}
}

func (suite *HandlerTestSuite) TestHealthcheck() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8088/healthcheck", nil)
	assert.Nil(suite.T(), err)
	resp, err := suite.client.Do(req)
	assert.Nil(suite.T(), err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(suite.T(), err)
	suite.Equal(200, resp.StatusCode)
	suite.Equal(`{"status":"WORKING"}
`, string(body))
}

func (suite *HandlerTestSuite) TestIndex() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8088/", nil)
	assert.Nil(suite.T(), err)
	resp, err := suite.client.Do(req)
	assert.Nil(suite.T(), err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(suite.T(), err)
	suite.Equal(200, resp.StatusCode)
	suite.Equal(fmt.Sprintf(`{"version":"%s"}
`, VERSION), string(body))
}

func (suite *HandlerTestSuite) TestGetAttribute() {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8088/attribute/fakeAttribute", nil)
	assert.Nil(suite.T(), err)
	resp, err := suite.client.Do(req)
	assert.Nil(suite.T(), err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(suite.T(), err)
	suite.Equal(200, resp.StatusCode)
	suite.Equal(`{"name":"fakeAttribute","type":"string"}
`, string(body))
}

func TestHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HandlerTestSuite))
}
