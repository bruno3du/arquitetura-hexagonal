package cli

import "github.com/bruno3du/hexagonal/application"

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {
	var result = ""
	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return "", err
		}
		result = "Product: " + product.GetName() + " has been created with success"
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		_, err = service.Enable(product)
		if err != nil {
			return result, err
		}
		result = "Product: " + product.GetName() + " has been enabled with success"
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		_, err = service.Disable(product)
		if err != nil {
			return result, err
		}
		result = "Product: " + product.GetName() + " has been disabled with success"
	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		result = "Product: " + res.GetName() + " has been found with success"
	}

	return result, nil
}
