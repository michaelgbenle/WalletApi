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
		mockDB.EXPECT().Getcustomer(accountNos).Return(nil, errors.New("errors exist"))
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
	var accountNos = "1187654311"
	transactions := []models.Transaction{
		{
			Model:      gorm.Model{ID: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
			CustomerId: 2,
			AccountNos: "1187654311",
			Type:       "credit",
			Success:    true,
		},
		{
			Model:      gorm.Model{ID: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
			CustomerId: 2,
			AccountNos: "1187654311",
			Type:       "credit",
			Success:    true,
		},
		{
			Model:      gorm.Model{ID: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
			CustomerId: 2,
			AccountNos: "1187654311",
			Type:       "credit",
			Success:    true,
		},
	}
	bodyJSON, err := json.Marshal(transactions)
	if err != nil {
		t.Fail()
	}
	t.Run("Testing for all transactions", func(t *testing.T) {
		mockDB.EXPECT().Gettransaction(accountNos).Return(&transactions, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/transactions?accountNos=1187654311", strings.NewReader(string(bodyJSON)))

		route.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)

	})

}

func TestCreditWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockdatabase.NewMockDB(ctrl)
	h := handler.Handler{DB: mockDB}
	route, _ := router.SetupRouter(&h)

	credit := &models.Money{
		AccountNos: "1187654311",
		Amount:     0,
	}

	transaction := models.Transaction{
		AccountNos: "1187654311",
		Type:       "credit",
		Success:    true,
	}
	bodyJSON, err := json.Marshal(transaction)
	if err != nil {
		t.Fail()
	}

	mockDB.EXPECT().Creditwallet(credit).Return(&transaction, nil)
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPatch, "/credit", strings.NewReader(string(bodyJSON)))

	route.ServeHTTP(rw, req)
	assert.Equal(t, http.StatusOK, rw.Code)
	assert.NotContains(t, rw.Body.String(), transaction)

}

func TestDebitWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockdatabase.NewMockDB(ctrl)
	h := handler.Handler{DB: mockDB}
	route, _ := router.SetupRouter(&h)

	debit := &models.Money{
		AccountNos: "1187654311",
		Amount:     0,
	}
	transaction := models.Transaction{
		AccountNos: "1187654311",
		Type:       "credit",
		Success:    false,
	}
	customer := models.Customer{

		Name:       "bella",
		AccountNos: "1187654311",
		Balance:    0,
	}
	bodyJSON, err := json.Marshal(transaction)
	if err != nil {
		t.Fail()
	}
	mockDB.EXPECT().Getcustomer(debit.AccountNos).Return(&customer, nil)
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/debit", strings.NewReader(string(bodyJSON)))
	route.ServeHTTP(rw, req)

	mockDB.EXPECT().CreateTransaction(transaction).Return(&transaction, nil)
	rw = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/debit", strings.NewReader(string(bodyJSON)))

	assert.Equal(t, http.StatusOK, rw.Code)
	assert.NotContains(t, rw.Body.String(), transaction)

}
