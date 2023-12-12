package application_test

import (
	"testing"

	"github.com/felipefbs/hex-arch/application"
	mock_application "github.com/felipefbs/hex-arch/application/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestProductService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := mock_application.NewMockIProduct(ctrl)
	persistence := mock_application.NewMockIProductPersistence(ctrl)

	svc := application.ProductService{
		Persistence: persistence,
	}

	t.Run("Get product", func(t *testing.T) {
		persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

		productFound, err := svc.Get(uuid.NewString())
		assert.Nil(t, err)
		assert.NotEmpty(t, productFound)
	})

	t.Run("Save product", func(t *testing.T) {
		persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

		productSaved, err := svc.Create("testProduct", 10.0)
		assert.Nil(t, err)
		assert.NotEmpty(t, productSaved)
	})

	t.Run("Enable product", func(t *testing.T) {
		persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
		product.EXPECT().Enable().Return(nil).AnyTimes()

		productSaved, err := svc.Enable(product)
		assert.Nil(t, err)
		assert.NotEmpty(t, productSaved)
	})

	t.Run("Disable product", func(t *testing.T) {
		persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
		product.EXPECT().Disable().Return(nil).AnyTimes()

		productSaved, err := svc.Disable(product)
		assert.Nil(t, err)
		assert.NotEmpty(t, productSaved)
	})
}
