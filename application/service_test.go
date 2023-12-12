package application_test

import (
	"testing"

	mock_application "github.com/felipefbs/hex-arch/application/mocks"
	"github.com/golang/mock/gomock"
)

func TestProductService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockIProduct(ctrl)
	persistence := mock_application.NewMockIProductPersistence(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()
}
