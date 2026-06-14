package servicios

import (
	"ecommerce/modelos"
	"ecommerce/utils"
)

// funcion de crear pedido
func CrearPedido(id int, cliente *modelos.Usuario, carrito *modelos.Carrito) (*modelos.Pedido, error) {
	items := carrito.GetItems()
	if len(items) == 0 {
		return nil, utils.ErrPedidoVacío
	}
	if len(items) > 1 {
		return nil, utils.ErrPedidoConMultiplesItems
	}

	item := items[0]
	pedido := modelos.NuevoPedido(id, cliente, item.Producto, item.Cantidad)
	return pedido, nil

}

// funcion Actualizar Inventario
func ActualizarInventario(items []modelos.ItemCarrito) error {
	for _, item := range items {
		nuevoStock := item.Producto.GetStock() - item.Cantidad
		err := item.Producto.SetStock(nuevoStock)
		if err != nil {
			return err
		}
	}
	return nil
}
