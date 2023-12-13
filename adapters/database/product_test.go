package database_test

import (
	"database/sql"
	"testing"

	"github.com/felipefbs/hex-arch/adapters/database"
	"github.com/felipefbs/hex-arch/application"
	_ "github.com/glebarez/go-sqlite"
	"github.com/stretchr/testify/assert"
)

var (
	product1 = application.Product{ID: "0b75ad26-0a25-459e-810f-0a0cbcc7d677", Name: "test product 1", Price: 10.20, Status: application.Enabled}
	product2 = application.Product{ID: "fd207a43-102a-48e4-b6b9-6816d8e0df2b", Name: "test product 2", Price: 0, Status: application.Disabled}
)

func bootstrapTest() *sql.DB {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS products (id string, name string, price float64, status string)")
	if err != nil {
		panic(err)
	}

	db.Exec("INSERT INTO products (id, name, price, status) VALUES ($1,$2,$3,$4)", product1.ID, product1.Name, product1.Price, product1.Status)
	db.Exec("INSERT INTO products (id, name, price, status) VALUES ($1,$2,$3,$4)", product2.ID, product2.Name, product2.Price, product2.Status)

	return db
}

func TestProductDatabase(t *testing.T) {
	db := bootstrapTest()
	defer db.Close()

	productDB := database.NewProductDatabase(db)

	t.Run("Get product", func(t *testing.T) {
		foundProduct, err := productDB.Get(product1.ID)
		assert.Nil(t, err)
		assert.Equal(t, product1.ID, foundProduct.GetID())
	})

	t.Run("Saves a new product", func(t *testing.T) {
		product3 := application.NewProduct("product3", 10)
		savedProduct, err := productDB.Save(product3)
		assert.Nil(t, err)
		assert.Equal(t, product3.ID, savedProduct.GetID())
	})

	t.Run("Saves an already existing product", func(t *testing.T) {
		product2.Price = 20
		product2.Status = application.Enabled
		savedProduct, err := productDB.Save(&product2)
		assert.Nil(t, err)

		foundProduct, err := productDB.Get(savedProduct.GetID())
		assert.Nil(t, err)

		assert.Equal(t, product2.ID, foundProduct.GetID())
		assert.Equal(t, product2.Price, foundProduct.GetPrice())
		assert.Equal(t, product2.Status, foundProduct.GetStatus())
	})
}
