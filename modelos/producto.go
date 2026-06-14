package modelos

import "errors"

// producto representa un producto del e-commerce
type Producto struct {
	ID     int
	Nombre string
	Precio float64
	stock  int
}

//contructor para crear un nuevo producto
func NuevoProducto(id int, nombre string, precio float64, stock int) *Producto {
	return &Producto{
		ID:     id,
		Nombre: nombre,
		Precio: precio,
		stock:  stock,
	}
}

// getter para obtener el precio del producto
func (p Producto) GetID() int {
	return p.ID
}
func (p Producto) GetNombre() string {
	return p.Nombre
}
func (p Producto) GetPrecio() float64 {
	return p.Precio
}
func (p Producto) GetStock() int {
	return p.stock
}

// setter para actualizar el precio del producto
func (p *Producto) SetPrecio(precio float64) error {
	if precio < 0 {
		return errors.New("el precio no puede ser negativodebe ser mayo a cero")
	}
	p.Precio = precio
	return nil
}
func (p *Producto) SetStock(stock int) error {
	if stock < 0 {
		return errors.New("el stock no puede ser negativo debe ser mayo a cero")
	}
	p.stock = stock
	return nil
}
