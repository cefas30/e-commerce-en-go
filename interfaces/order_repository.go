package interfaces

import modelos "ecommerce/models"

// PedidoRepository define el contrato para persistir pedidos.
type PedidoRepository interface {
	Guardar(pedido *modelos.Pedido) error
	ObtenerPorID(id int) (*modelos.Pedido, error)
	ListarPedidos() ([]*modelos.Pedido, error)
}
