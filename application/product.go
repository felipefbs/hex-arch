package application

import (
	"errors"

	"github.com/google/uuid"
)

type IProduct interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	Disabled = "disabled"
	Enabled  = "enabled"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
}

func NewProduct(name string, price ...float64) *Product {
	product := &Product{
		ID:     uuid.NewString(),
		Name:   name,
		Price:  0,
		Status: Disabled,
	}

	if len(price) != 0 {
		return product
	}

	product.Price = price[0]

	return product
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = Disabled
	}

	if p.Status != Enabled && p.Status != Disabled {
		return false, errors.New("the status must be enabled or disabled")
	}

	if p.Price < 0 {
		return false, errors.New("the price must be greater or equal zero")
	}

	_, err := uuid.Parse(p.ID)
	if err != nil {
		return false, errors.New("invalid id")
	}

	if p.Name == "" {
		return false, errors.New("invalid name")
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price <= 0 {
		p.Status = Disabled

		return errors.New("the price must be greater than zero to enable the product")
	}

	p.Status = Enabled

	return nil
}

func (p *Product) Disable() error {
	if p.Price > 0 {
		return errors.New("the price must be equal zero to disable the product")
	}

	p.Status = Disabled

	return nil
}

func (p *Product) GetID() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
