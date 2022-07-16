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
	defer ctrl.Finish()
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

	t.Run("Testing for error", func(t *testing.T) {
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
		AccountNos: credit.AccountNos,
		Type:       "credit",
		Success:    true,
	}
	bodyJSON, err := json.Marshal(transaction)
	if err != nil {
		t.Fail()
	}
	t.Run("Testing for error", func(t *testing.T) {
		mockDB.EXPECT().Creditwallet(credit).Return(nil, errors.New("errors exist"))
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/credit", strings.NewReader(string(bodyJSON)))

		route.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusInternalServerError, rw.Code)
		assert.Contains(t, rw.Body.String(), "unable to credit wallet")
	})

	t.Run("Testing for success", func(t *testing.T) {
		mockDB.EXPECT().Creditwallet(credit).Return(&transaction, nil)
		rw := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPatch, "/credit", strings.NewReader(string(bodyJSON)))

		route.ServeHTTP(rw, req)
		assert.Equal(t, http.StatusOK, rw.Code)
		assert.Contains(t, rw.Body.String(), string(bodyJSON))
	})
}

func TestDebitWallet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockdatabase.NewMockDB(ctrl)
	h := handler.Handler{DB: mockDB}
	route, _ := router.SetupRouter(&h)

	customer := models.Customer{
		Model:      gorm.Model{ID: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}},
		Name:       "Rose",
		AccountNos: "1187654311",
		Balance:    3000,
	}
	debit := &models.Money{
		AccountNos: "1187654311",
		Amount:     0,
	}

	transaction := &models.Transaction{
		Model: gorm.Model{
			ID: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}},
		CustomerId: customer.ID,
		AccountNos: debit.AccountNos,
		Type:       "debit",
		Success:    false,
	}
	transaction1 := models.Transaction{
		Model: gorm.Model{
			ID: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}},
		CustomerId: customer.ID,
		AccountNos: debit.AccountNos,
		Type:       "debit",
		Success:    true,
	}
	bodyJSON, _ := json.Marshal(transaction1)

	t.Run("Testing for error", func(t *testing.T) {
		mockDB.EXPECT().Getcustomer(debit.AccountNos).Return(&customer, nil)
		mockDB.EXPECT().CreateTransaction(transaction).AnyTimes()
		mockDB.EXPECT().Debitwallet(debit).Return(nil, errors.New("unable to debit wallet"))
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PATCH", "/debit", strings.NewReader(string(bodyJSON)))
		route.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "unable to debit wallet")
	})

	t.Run("Testing for success", func(t *testing.T) {
		mockDB.EXPECT().Getcustomer(debit.AccountNos).Return(&customer, nil)
		mockDB.EXPECT().CreateTransaction(transaction).AnyTimes()
		mockDB.EXPECT().Debitwallet(debit).Return(&transaction1, nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PATCH", "/debit", strings.NewReader(string(bodyJSON)))
		route.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), string(bodyJSON))
	})

}

func TestAddCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockdatabase.NewMockDB(ctrl)
	h := handler.Handler{DB: mockDB}
	route, _ := router.SetupRouter(&h)

	customer := models.Customer{
		Model: gorm.Model{
			ID: 1, CreatedAt: time.Time{}, UpdatedAt: time.Time{}, DeletedAt: gorm.DeletedAt{}},
		Name:       "Bella",
		AccountNos: "1187654311",
		Balance:    0,
	}
	bodyJSON, err := json.Marshal(&customer)
	if err != nil {
		t.Fail()
	}

	t.Run("Testing for add customer", func(t *testing.T) {
		//mockDB.EXPECT().Getcustomer(customer.AccountNos).Return(&customer, err)
		mockDB.EXPECT().Addcustomer(&customer).Return(nil)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/addcustomer", strings.NewReader(string(bodyJSON)))
		route.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "unable to bind json")

	})

}
