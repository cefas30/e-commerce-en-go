package services

import (
	modelos "ecommerce/models"
	repositorios "ecommerce/repositories"
)

type ProductoService struct {
	repository *repositorios.ProductoRepository
}

func NuevoProductoService() *ProductoService {

	return &ProductoService{

		repository: repositorios.NuevoProductoRepository(),
	}
}

func (s *ProductoService) ObtenerTodos() []modelos.Producto {

	return s.repository.ObtenerTodos()

}

func (s *ProductoService) ObtenerPorID(id int) (*modelos.Producto, error) {

	return s.repository.ObtenerPorID(id)

}

func (s *ProductoService) Crear(producto modelos.Producto) modelos.Producto {

	return s.repository.Crear(producto)

}

func (s *ProductoService) Actualizar(id int, producto modelos.Producto) error {

	return s.repository.Actualizar(id, producto)

}

func (s *ProductoService) Eliminar(id int) error {

	return s.repository.Eliminar(id)

}
