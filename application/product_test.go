package application_test

import (
	"testing"

	"github.com/felipefbs/hex-arch/application"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProductStatus(t *testing.T) {
	t.Run("Enabling product", func(t *testing.T) {
		product := application.Product{
			ID:     uuid.NewString(),
			Name:   "Vibrador Ultra Master v2",
			Price:  10,
			Status: application.Disabled,
		}

		err := product.Enable()
		assert.Nil(t, err)
		assert.Equal(t, application.Enabled, product.Status)
	})

	t.Run("Enabling product but getting price error", func(t *testing.T) {
		product := application.Product{
			ID:     uuid.NewString(),
			Name:   "Vibrador Ultra Master v2",
			Price:  0,
			Status: application.Disabled,
		}

		err := product.Enable()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "the price must be greater than zero to enable the product")
		assert.NotEqual(t, application.Enabled, product.Status)
	})

	t.Run("Disabling product", func(t *testing.T) {
		product := application.Product{
			ID:     uuid.NewString(),
			Name:   "Vibrador Ultra Master v2",
			Price:  0,
			Status: application.Disabled,
		}

		err := product.Disable()
		assert.Nil(t, err)
		assert.Equal(t, application.Disabled, product.Status)
	})

	t.Run("Disabling product but getting price error", func(t *testing.T) {
		product := application.Product{
			ID:     uuid.NewString(),
			Name:   "Vibrador Ultra Master v2",
			Price:  10,
			Status: application.Disabled,
		}

		err := product.Disable()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "the price must be equal zero to disable the product")
		assert.NotEqual(t, application.Enabled, product.Status)
	})
}

func TestProductValidation(t *testing.T) {
	t.Run("Valid Product", func(t *testing.T) {
		product := application.Product{
			ID:     uuid.NewString(),
			Name:   "Mouse ultra master",
			Price:  50.51,
			Status: application.Enabled,
		}

		isValid, err := product.IsValid()
		assert.Nil(t, err)
		assert.True(t, isValid)
	})

	t.Run("Invalid Product - ID", func(t *testing.T) {
		product := application.Product{
			ID:     "uuid",
			Name:   "Mouse ultra master",
			Price:  50.51,
			Status: application.Enabled,
		}

		isValid, err := product.IsValid()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "invalid id")
		assert.False(t, isValid)
	})

	t.Run("Invalid Product - Name", func(t *testing.T) {
		product := application.Product{
			ID:     uuid.NewString(),
			Name:   "",
			Price:  50.51,
			Status: application.Enabled,
		}

		isValid, err := product.IsValid()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "invalid name")
		assert.False(t, isValid)
	})

	t.Run("Invalid Product - Price", func(t *testing.T) {
		product := application.Product{
			ID:     uuid.NewString(),
			Name:   "Mouse ultra master",
			Price:  -20.6,
			Status: application.Enabled,
		}

		isValid, err := product.IsValid()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "the price must be greater or equal zero")
		assert.False(t, isValid)
	})

	t.Run("Invalid Product - Status", func(t *testing.T) {
		product := application.Product{
			ID:     uuid.NewString(),
			Name:   "Mouse ultra master",
			Price:  20.6,
			Status: "INVALID",
		}

		isValid, err := product.IsValid()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "the status must be enabled or disabled")
		assert.False(t, isValid)
	})
}
