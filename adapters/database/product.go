package database

import (
	"database/sql"
	"fmt"

	"github.com/felipefbs/hex-arch/application"
)

type ProductDatabase struct {
	db *sql.DB
}

func NewProductDatabase(db *sql.DB) application.IProductPersistence {
	return &ProductDatabase{db: db}
}

func (repo *ProductDatabase) Get(id string) (application.IProduct, error) {
	stmt, err := repo.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}

	productFound := application.Product{}
	if err := stmt.QueryRow(id).Scan(
		&productFound.ID, &productFound.Name, &productFound.Price, &productFound.Status,
	); err != nil {
		return nil, err
	}

	return &productFound, nil
}

func (repo *ProductDatabase) SaveV1(product application.IProduct) (application.IProduct, error) {
	found, err := repo.Get(product.GetID())
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if found == nil {
		return repo.create(product)
	}

	return repo.update(product)
}

func (repo *ProductDatabase) Save(product application.IProduct) (application.IProduct, error) {
	id := ""
	err := repo.db.QueryRow("SELECT id FROM products WHERE id = $1", product.GetID()).Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return repo.create(product)
	}

	return repo.update(product)
}

func (repo *ProductDatabase) create(product application.IProduct) (application.IProduct, error) {
	result, err := repo.db.Exec(
		"INSERT INTO products (id, name, price, status) VALUES ($1, $2, $3, $4)",
		product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil || rows != 1 {
		return nil, fmt.Errorf("invalid rows affected")
	}

	return product, nil
}

func (repo *ProductDatabase) update(product application.IProduct) (application.IProduct, error) {
	_, err := repo.db.Exec(
		"UPDATE products SET name = $1, price = $2, status = $3 where id = $4",
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID(),
	)

	return product, err
}
