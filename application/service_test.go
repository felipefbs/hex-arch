package application_test

import (
	"testing"

	"github.com/felipefbs/hex-arch/application"
	mock_application "github.com/felipefbs/hex-arch/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProductService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockIProduct(ctrl)
	persistence := mock_application.NewMockIProductPersistence(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	svc := application.ProductService{
		Persistence: persistence,
	}

	productFound, err := svc.Get(uuid.NewString())
	assert.Nil(t, err)
	assert.NotEmpty(t, productFound)
}
