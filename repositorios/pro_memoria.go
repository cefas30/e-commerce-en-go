package repositorios

import (
	"ecommerce/interfaces"
	"ecommerce/modelos"
	"ecommerce/utils"
	"errors"
)

// ProMemoria es una implementación en memoria del repositorio de productos
type ProMemoria struct {
	productos map[int]*modelos.Producto
}

// NewProMemoria crea una nueva instancia de ProMemoria
func NewProMemoria() *ProMemoria {
	return &ProMemoria{
		productos: make(map[int]*modelos.Producto),
	}
}

// GetProductoByID obtiene un producto por su ID
func (r *ProMemoria) GetProductoByID(id int) (*modelos.Producto, error) {
	producto, exists := r.productos[id]
	if !exists {
		return nil, utils.ErrorProductoNoEncontrado
	}
	return producto, nil
}

// SaveProducto guarda un nuevo producto en el repositorio
func (r *ProMemoria) SaveProducto(producto *modelos.Producto) error {
	if _, exists := r.productos[producto.ID]; exists {
		return errors.New("el producto con este ID ya existe")
	}
	r.productos[producto.ID] = producto
	return nil
}

// UpdateProducto actualiza un producto existente en el repositorio
func (r *ProMemoria) UpdateProducto(producto *modelos.Producto) error {
	if _, exists := r.productos[producto.ID]; !exists {
		return utils.ErrorProductoNoEncontrado
	}
	r.productos[producto.ID] = producto
	return nil
}

// DeleteProducto elimina un producto del repositorio por su ID
func (r *ProMemoria) DeleteProducto(id int) error {
	if _, exists := r.productos[id]; !exists {
		return utils.ErrorProductoNoEncontrado
	}
	delete(r.productos, id)
	return nil
}

// ListProductos devuelve una lista de todos los productos en el repositorio
func (r *ProMemoria) ListProductos() ([]*modelos.Producto, error) {
	productos := make([]*modelos.Producto, 0, len(r.productos))
	for _, producto := range r.productos {
		productos = append(productos, producto)
	}
	return productos, nil
}

// Aseguramos que ProMemoria implementa la interfaz ProRepository
var _ interfaces.ProRepository = (*ProMemoria)(nil)
