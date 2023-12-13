package application

type ProductService struct {
	Persistence IProductPersistence
}

func NewProductService(repo IProductPersistence) IProductService {
	return &ProductService{Persistence: repo}
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

	return svc.Persistence.Save(product)
}

func (svc *ProductService) Enable(product IProduct) (IProduct, error) {
	err := product.Enable()
	if err != nil {
		return nil, err
	}

	return svc.Persistence.Save(product)
}

func (svc *ProductService) Disable(product IProduct) (IProduct, error) {
	err := product.Disable()
	if err != nil {
		return nil, err
	}

	return svc.Persistence.Save(product)
}
