package modelos

// Carrito representa un carrito de compras
type ItemCarrito struct {
	Producto *Producto
	Cantidad int
}

// estructura para representar un carrito de compras, con un producto y su cantidad.
type Carrito struct {
	idUsuario int
	items     []ItemCarrito
}

//constructor para crear un nuevo carrito de compras para un usuario específico.
func NuevoCarrito(idUsuario int) *Carrito {
	return &Carrito{
		idUsuario: idUsuario,
		items:     []ItemCarrito{},
	}
}

//getters para acceder a los campos del carrito.
func (c Carrito) GetItems() []ItemCarrito {

	return c.items
}

// agregar un producto al carrito, si el producto ya existe en el carrito, se actualiza la cantidad.
func (c *Carrito) AgregarProducto(producto *Producto, cantidad int) {
	item := ItemCarrito{
		Producto: producto,
		Cantidad: cantidad,
	}
	c.items = append(c.items, item)
}

// eliminar un producto del carrito por su ID, si el producto no se encuentra, no se hace nada.
func (c *Carrito) EliminarProducto(idProducto int) {
	for i, item := range c.items {
		if item.Producto.GetID() == idProducto {
			c.items = append(c.items[:i], c.items[i+1:]...)
			break
		}
	}
}

// calcular el total del carrito sumando el precio de cada producto multiplicado por su cantidad.
func (c Carrito) CalcularTotal(items []ItemCarrito) float64 {
	total := 0.0
	for _, item := range items {
		total += item.Producto.GetPrecio() * float64(item.Cantidad)
	}
	return total
}
