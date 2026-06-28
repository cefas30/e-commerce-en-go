package services

import (
	modelos "ecommerce/models"
	repositorios "ecommerce/repositories"
)

type PedidoService struct {
	repository *repositorios.PedidoRepository
}

func NuevoPedidoService() *PedidoService {

	return &PedidoService{

		repository: repositorios.NuevoPedidoRepository(),
	}

}

func (s *PedidoService) ObtenerTodos() []modelos.Pedido {

	return s.repository.ObtenerTodos()

}

func (s *PedidoService) Crear(pedido modelos.Pedido) modelos.Pedido {

	pedido.Estado = modelos.Pendiente

	return s.repository.Crear(pedido)

}
