package repositorios

import (
	"ecommerce/modelos"
	"ecommerce/utils"
)

type PeddMemoria struct {
	pedidos []*modelos.Pedido
}

func NewPedMemoria() *PeddMemoria {
	return &PeddMemoria{
		pedidos: make([]*modelos.Pedido, 0),
	}
}

// Guardar almacena un pedido en memoria.
func (r *PeddMemoria) Guardar(pedido *modelos.Pedido) error {
	r.pedidos = append(r.pedidos, pedido)
	return nil
}

// ObtenerPorID busca un pedido por su ID.
func (r *PeddMemoria) ObtenerPorID(
	id int,
) (*modelos.Pedido, error) {

	for _, pedido := range r.pedidos {

		if pedido.GetID() == id {

			return pedido, nil
		}
	}

	return nil,
		utils.ErrPedidoNoExiste
}

// ListarPedidos devuelve todos los pedidos almacenados.
func (r *PeddMemoria) ListarPedidos() ([]*modelos.Pedido, error) {
	return r.pedidos, nil
}
