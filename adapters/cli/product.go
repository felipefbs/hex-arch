package cli

import (
	"fmt"

	"github.com/felipefbs/hex-arch/application"
)

func Run(service application.IProductService, action string, product application.IProduct) (string, error) {
	result := ""

	switch action {
	case "create":
		productCreated, err := service.Create(product.GetName(), product.GetPrice())
		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("created product: %v", productCreated.String())

	case "enable":
		productFound, err := service.Get(product.GetID())
		if err != nil {
			return "", err
		}

		if _, err = service.Enable(productFound); err != nil {
			return "", err
		}

		result = fmt.Sprintf("enabled product: %v", productFound.String())

	case "disable":
		productFound, err := service.Get(product.GetID())
		if err != nil {
			return "", err
		}

		if _, err = service.Disable(productFound); err != nil {
			return "", err
		}

		result = fmt.Sprintf("disabled product: %v", productFound.String())

	case "find":
		productFound, err := service.Get(product.GetID())
		if err != nil {
			return "", err
		}

		result = fmt.Sprintf("found product: %v", productFound.String())
	}

	return result, nil
}
