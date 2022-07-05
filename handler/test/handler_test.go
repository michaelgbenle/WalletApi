package test

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	mockdatabase "github.com/michaelgbenle/WalletApi/database/mocks"
	"github.com/michaelgbenle/WalletApi/handler"
	"github.com/michaelgbenle/WalletApi/models"
	"github.com/michaelgbenle/WalletApi/router"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestGetCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	//defer ctrl.Finish()
	mockDB := mockdatabase.NewMockDB(ctrl)
	h := handler.Handler{DB: mockDB}
	route, _ := router.SetupRouter(&h)
	var accountNos = "1187654311"
	customer := models.Customer{
		Name:       "Bella",
		AccountNos: "1187654311",
		Balance:    0,
	}
	bodyJSON, err := json.Marshal(&customer)
	if err != nil {
		t.Fail()
	}
	t.Run("Testing for get customer", func(t *testing.T) {
		mockDB.EXPECT().Getcustomer(accountNos).Return(nil, errors.New("account number should be 10 digits"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/customer?accountNos=1187654311", strings.NewReader(string(bodyJSON)))

		route.ServeHTTP(rw, req)
		assert.Contains(t, rw.Body.String(), "user not found")
		assert.Equal(t, rw.Code, 400)

	})
	t.Run("Testing for get customer", func(t *testing.T) {
		mockDB.EXPECT().Getcustomer(accountNos).Return(&customer, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/customer?accountNos=1187654311", strings.NewReader(string(bodyJSON)))
		route.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotContains(t, w.Body.String(), string(bodyJSON))

	})
}
func TestGetTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockdatabase.NewMockDB(ctrl)
	h := handler.Handler{DB: mockDB}
	route, _ := router.SetupRouter(&h)
	transactions := &[]models.Transaction{
		{
			Model: gorm.Model{ID: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
			CustomerId: 2,
			AccountNos: "1187654311",
			Type: "credit",
			Success: true,
		},
		{
			Model: gorm.Model{ID: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
			CustomerId: 2,
			AccountNos: "1187654311",
			Type: "credit",
			Success: true,

		},
		{
			"ID": 3,
			"CreatedAt": "2022-07-05T09:53:19.243563+01:00",
			"UpdatedAt": "2022-07-05T09:53:19.25113+01:00",
			"DeletedAt": null,
			"customer_id": 2,
			"accountNos": "1187654311",
			"type": "debit",
			"success": true
		},
		{
			"ID": 4,
			"CreatedAt": "2022-07-05T10:07:39.939194+01:00",
			"UpdatedAt": "2022-07-05T10:07:39.944188+01:00",
			"DeletedAt": null,
			"customer_id": 2,
			"accountNos": "1187654311",
			"type": "debit",
			"success": true
		},
		{
			"ID": 5,
			"CreatedAt": "2022-07-05T11:37:06.41214+01:00",
			"UpdatedAt": "2022-07-05T11:37:06.418631+01:00",
			"DeletedAt": null,
			"customer_id": 2,
			"accountNos": "1187654311",
			"type": "debit",
			"success": true
		}
	}

}
