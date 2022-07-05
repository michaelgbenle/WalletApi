package test

import (
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	mockdatabase "github.com/michaelgbenle/WalletApi/database/mocks"
	"github.com/michaelgbenle/WalletApi/handler"
	"github.com/michaelgbenle/WalletApi/models"
	"github.com/michaelgbenle/WalletApi/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockdatabase.NewMockDB(ctrl)
	h := handler.Handler{DB: mockDB}
	route, _ := router.SetupRouter(h)
	var accountNos uint = 1187654311
	customer := models.Customer{
		Name:       "Bella",
		AccountNos: "",
		Balance:    1187654311,
	}
	bodyJSON, err := json.Marshal(customer)
	if err != nil {
		t.Fail()
	}
	t.Run("Testing for all product", func(t *testing.T) {
		mockDB.EXPECT().Getcustomer().Return(customer)
		rw := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet,
			"/api/v1/products",
			strings.NewReader(string(bodyJSON)))
		if err != nil {
			fmt.Printf("errrr here %v \n", err)
			return
		}
		route.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), string(bodyJSON))
	})
}
