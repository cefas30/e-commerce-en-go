package interfaces

import modelos "ecommerce/models"

// ProRepository define la interfaz para el repositorio de productos
type ProductoRepository interface {
	ObtenerTodos() []modelos.Producto
	ObtenerPorID(id int) (*modelos.Producto, error)
	Crear(producto modelos.Producto) modelos.Producto
	Actualizar(id int, producto modelos.Producto) error
	Eliminar(id int) error
}
