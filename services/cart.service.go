package services

import (
	modelos "ecommerce/models"
)

type CarritoService struct {
	carritos []modelos.Carrito
}

func NuevoCarritoService() *CarritoService {

	return &CarritoService{}
}

func (s *CarritoService) ObtenerTodos() []modelos.Carrito {

	return s.carritos

}

func (s *CarritoService) Crear(usuarioID int) modelos.Carrito {

	carrito := modelos.Carrito{

		ID:        len(s.carritos) + 1,
		UsuarioID: usuarioID,
		Items:     []modelos.ItemCarrito{},
		Total:     0,
	}

	s.carritos = append(s.carritos, carrito)

	return carrito

}
