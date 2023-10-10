package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/stretchr/testify/assert"
)

func mockHttpService() *HttpService {

	svc := HttpService{}

	svc.corsConfig = cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	svc.devMode = os.Getenv("RUN_MODE") == "DEV"
	svc.testMode = os.Getenv("TEST_HTTP") == "TRUE"
	if svc.devMode {
		log.Println("-DEV_MODE enabled-")
	}

	if err := svc.registerRoutes(); err != nil {
		fmt.Println(err)
		panic("failed to register routes")
	}

	return &svc
}

func TestPingAPI(t *testing.T) {
	httpmock := mockHttpService()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/ping", nil)
	httpmock.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}
