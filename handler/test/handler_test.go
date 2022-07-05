package test

import (
	"github.com/golang/mock/gomock"
	mockdatabase "github.com/michaelgbenle/WalletApi/database/mocks"
	"testing"
)

func TestGetCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockDB := mockdatabase.NewMockDB(ctrl)

}
