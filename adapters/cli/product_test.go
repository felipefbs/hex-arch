package cli_test

import (
	"errors"
	"testing"

	"github.com/felipefbs/hex-arch/adapters/cli"
	"github.com/felipefbs/hex-arch/application"
	mock_application "github.com/felipefbs/hex-arch/application/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestProductCLI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := application.Product{
		ID:     uuid.NewString(),
		Name:   "product test",
		Price:  10.0,
		Status: application.Enabled,
	}

	t.Run("Create product", func(t *testing.T) {
		productMock := mock_application.NewMockIProduct(ctrl)
		svc := mock_application.NewMockIProductService(ctrl)

		productMock.EXPECT().GetName().Return(product.Name).AnyTimes()
		productMock.EXPECT().GetPrice().Return(product.Price).AnyTimes()
		productMock.EXPECT().String().Return("product.String()").AnyTimes()

		svc.EXPECT().Create(product.Name, product.Price).Return(productMock, nil).AnyTimes()

		result, err := cli.Run(svc, "create", productMock)
		assert.Nil(t, err)
		assert.Contains(t, result, "created")
	})

	t.Run("Failed to create product", func(t *testing.T) {
		productMock := mock_application.NewMockIProduct(ctrl)
		svc := mock_application.NewMockIProductService(ctrl)

		productMock.EXPECT().GetName().Return(product.Name).AnyTimes()
		productMock.EXPECT().GetPrice().Return(product.Price).AnyTimes()

		svc.EXPECT().Create(product.Name, product.Price).Return(nil, errors.New("AAAAAAAAaa")).AnyTimes()

		result, err := cli.Run(svc, "create", productMock)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})

	t.Run("Enable product", func(t *testing.T) {
		productMock := mock_application.NewMockIProduct(ctrl)
		svc := mock_application.NewMockIProductService(ctrl)

		productMock.EXPECT().GetID().Return(product.ID).AnyTimes()
		svc.EXPECT().Get(product.ID).Return(productMock, nil).AnyTimes()
		svc.EXPECT().Enable(productMock).Return(productMock, nil).AnyTimes()
		productMock.EXPECT().String().Return("product.String()").AnyTimes()

		result, err := cli.Run(svc, "enable", productMock)
		assert.Nil(t, err)
		assert.Contains(t, result, "enabled")
	})

	t.Run("Disable product", func(t *testing.T) {
		productMock := mock_application.NewMockIProduct(ctrl)
		svc := mock_application.NewMockIProductService(ctrl)

		productMock.EXPECT().GetID().Return(product.ID).AnyTimes()
		svc.EXPECT().Get(product.ID).Return(productMock, nil).AnyTimes()
		svc.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()
		productMock.EXPECT().String().Return("product.String()").AnyTimes()

		result, err := cli.Run(svc, "disable", productMock)
		assert.Nil(t, err)
		assert.Contains(t, result, "disabled")
	})

	t.Run("Find product", func(t *testing.T) {
		productMock := mock_application.NewMockIProduct(ctrl)
		svc := mock_application.NewMockIProductService(ctrl)

		productMock.EXPECT().GetID().Return(product.ID).AnyTimes()
		svc.EXPECT().Get(product.ID).Return(productMock, nil).AnyTimes()
		svc.EXPECT().Disable(productMock).Return(productMock, nil).AnyTimes()
		productMock.EXPECT().String().Return("product.String()").AnyTimes()

		result, err := cli.Run(svc, "find", productMock)
		assert.Nil(t, err)
		assert.Contains(t, result, "found")
	})
}
