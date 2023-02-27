package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"grabaciones/functional_test/internal/domain"
	"grabaciones/functional_test/internal/seller"
	"grabaciones/functional_test/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type Resp struct {
	Data domain.Seller `json:"data"`
}

type respList struct {
	Data []domain.Seller `json:"data"`
}

type errorRespone struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

var sellersStore = []domain.Seller{
	{
		ID:          1,
		CID:         123,
		CompanyName: "Digital House",
		Address:     "Av. Siempreviva 123",
		Telephone:   "123456789",
	},
	{
		ID:          2,
		CID:         456,
		CompanyName: "Google",
		Address:     "Av. Siempreviva 456",
		Telephone:   "987654321",
	},
}

var sellerStore = domain.Seller{
	ID:          2,
	CID:         456,
	CompanyName: "Google",
	Address:     "Av. Siempreviva 456",
	Telephone:   "987654321",
}

func createRequestResponse(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, httptest.NewRecorder()
}

func createServerTest(myMock *mocks.DbMock) *gin.Engine {

	repo := mocks.NewRepositorySeller(myMock)
	service := seller.NewService(repo)
	sellerHandler := NewSeller(service)

	rm := gin.Default()

	sr := rm.Group("/api/v1")

	sr.POST("/sellers", sellerHandler.Create())
	sr.GET("/sellers", sellerHandler.GetAll())
	sr.GET("/sellers/:id", sellerHandler.Get())

	return rm
}

//ARRANGE ACT ASSERT
func TestCreateOK(t *testing.T) {

	//arrange
	server := createServerTest(&mocks.DbMock{
		SellerMocked: sellersStore,
		Err:          nil},
	)
	sjson, _ := json.Marshal(sellerStore)

	//act
	req, rr := createRequestResponse(http.MethodPost, "/api/v1/sellers", string(sjson))

	server.ServeHTTP(rr, req)

	resp := Resp{
		Data: domain.Seller{},
	}

	json.Unmarshal(rr.Body.Bytes(), &resp)

	//assert
	assert.Equal(t, sellerStore, resp.Data)
	assert.Equal(t, 201, rr.Code)
}

func TestGetNotFound(t *testing.T) {

	errorExpect := errorRespone{
		Status:  404,
		Code:    "not_found",
		Message: "seller not found",
	}

	server := createServerTest(&mocks.DbMock{
		SellerMocked: sellersStore,
		Err:          errors.New("seller not found"),
	})

	req, rr := createRequestResponse(http.MethodGet, "/api/v1/sellers/30", "")

	server.ServeHTTP(rr, req)

	errResp := errorRespone{}
	json.Unmarshal(rr.Body.Bytes(), &errResp)

	assert.Equal(t, 404, rr.Code)
	assert.Equal(t, errorExpect, errResp)
}
