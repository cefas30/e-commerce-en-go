package servicios

import (
	"ecommerce/modelos"
	"ecommerce/utils"
	"fmt"
)

// validar disponibilidad de un producto
func ValidarDisponibilidad(producto *modelos.Producto, cantidad int) error {
	if producto.GetStock() < cantidad {
		return utils.ErrorProductoSinStock
	}
	return nil
}

// validar cantidad de productos a comprar
func ValidarCantidad(cantidad int) error {
	if cantidad <= 0 {
		return utils.ErrCantidadInvalida
	}
	return nil
}

// calcular total de la compra
func CalcularTotal(producto *modelos.Producto, cantidad int) float64 {
	return producto.GetPrecio() * float64(cantidad)
}

// mostrar detalles de la compra
func MostrarDetallesCompra(producto *modelos.Producto, cantidad int) {
	total := CalcularTotal(producto, cantidad)
	fmt.Printf("Producto: %s\nCantidad: %d\nPrecio Unitario: %.2f\nTotal: %.2f\n",
		producto.GetNombre(), cantidad, producto.GetPrecio(), total)
}
