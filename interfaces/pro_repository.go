package interfaces
import "ecommerce/modelos"
// ProRepository define la interfaz para el repositorio de productos
type ProRepository interface {
	// GetProductoByID obtiene un producto por su ID
	GetProductoByID(id int) (*modelos.Producto, error)
	// SaveProducto guarda un nuevo producto en el repositorio
	SaveProducto(producto *modelos.Producto) error	
	// UpdateProducto actualiza un producto existente en el repositorio
	UpdateProducto(producto *modelos.Producto) error	
	// DeleteProducto elimina un producto del repositorio por su ID
	DeleteProducto(id int) error	
	// ListProductos devuelve una lista de todos los productos en el repositorio
	ListProductos() ([]*modelos.Producto, error)
}	
	
