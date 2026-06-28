package validators

import (
	"errors"

	modelos "ecommerce/models"
)

func ValidateProduct(product modelos.Producto) error {

	if product.Nombre == "" {
		return errors.New("el nombre es obligatorio")
	}

	if product.Precio <= 0 {
		return errors.New("el precio debe ser mayor que cero")
	}

	if product.Stock < 0 {
		return errors.New("el stock no puede ser negativo")
	}

	return nil

}
