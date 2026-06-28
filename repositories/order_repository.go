package repositories

import (
	modelos "ecommerce/models"
	"errors"
)

type PedidoRepository struct {
	pedidos []modelos.Pedido
}

func NuevoPedidoRepository() *PedidoRepository {

	return &PedidoRepository{}

}

func (r *PedidoRepository) ObtenerTodos() []modelos.Pedido {

	return r.pedidos

}

func (r *PedidoRepository) Crear(pedido modelos.Pedido) modelos.Pedido {

	pedido.ID = len(r.pedidos) + 1

	r.pedidos = append(r.pedidos, pedido)

	return pedido

}

func (r *PedidoRepository) ObtenerPorID(id int) (*modelos.Pedido, error) {

	for i := range r.pedidos {

		if r.pedidos[i].ID == id {

			return &r.pedidos[i], nil

		}

	}

	return nil, errors.New("pedido no encontrado")

}
