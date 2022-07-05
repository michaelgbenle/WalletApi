package test

import (
	"github.com/golang/mock/gomock"
	mockdatabase "github.com/michaelgbenle/WalletApi/database/mocks"
	"github.com/michaelgbenle/WalletApi/handler"
	"github.com/michaelgbenle/WalletApi/router"
	"testing"
)

func TestGetCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mockdatabase.NewMockDB(ctrl)
	h := handler.Handler{DB: mockDB}
	route, _ := router.SetupRouter(h)

}
