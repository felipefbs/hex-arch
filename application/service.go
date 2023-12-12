package application

type IProductService interface {
	Get(id string) (IProduct, error)
	Create(name string, price float64) (IProduct, error)
	Enable(product IProduct) (IProduct, error)
	Disable(product IProduct) (IProduct, error)
}

type ProductService struct {
	Persistence IProductPersistence
}

func (svc *ProductService) Get(id string) (IProduct, error) {
	product, err := svc.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (svc *ProductService) Create(name string, price float64) (IProduct, error) {
	product := NewProduct(name, price)

	err := product.IsValid()
	if err != nil {
		return nil, err
	}

	result, err := svc.Persistence.Save(product)
	if err != nil {
		return nil, err
	}

	return result, nil
}
