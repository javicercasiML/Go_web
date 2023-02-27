package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"rest3/internal/domain"
	"rest3/internal/product"
	"rest3/pkg/store"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	// create router
	_ = os.Setenv("TOKEN", "123456")
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// create handler
	storage := store.NewStorage("./products1.json")
	repo := product.NewRepository(storage)
	service := product.NewService(repo)
	handler := NewProduct(service)

	// Define the routes
	products := router.Group("/products")
	{
		products.GET("", handler.Get())
		products.POST("", handler.Create())
		products.GET("/:id", handler.GetById())
		products.GET("/search", handler.GetSearch())
		products.DELETE(":id", handler.Delete())
		products.PATCH(":id", handler.Patch())
		products.PUT(":id", handler.Put())
	}
	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	request := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("token", "123456")

	return request, httptest.NewRecorder()
}

func Test_get(t *testing.T) {

	//arange
	server := createServer()
	request, response := createRequestTest("GET", "/products/", "")
	received := []domain.Product{}

	//act
	server.ServeHTTP(response, request)

	//Assert
	//assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, http.StatusOK, response.Code)
	err := json.Unmarshal(response.Body.Bytes(), &received)
	assert.Nil(t, err)
}
