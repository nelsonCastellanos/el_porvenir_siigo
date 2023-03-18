package controllers_test

import (
	"github.com/gin-gonic/gin"
	"github.com/nelsonCastellanos/golang-api-mongo/cmd/api/controllers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHealthCheckController_HandlePing(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	var params []gin.Param
	u := url.Values{}
	MockJsonGet(ctx, params, u)
	controller := controllers.NewHealthCheckController()

	controller.HandlePing(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)
	body := `{"message":"pong"}`
	assert.Equal(t, body, w.Body.String())

}

// mock getrequest
func MockJsonGet(c *gin.Context, params gin.Params, u url.Values) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")

	// set path params
	c.Params = params

	// set query params
	c.Request.URL.RawQuery = u.Encode()
}
