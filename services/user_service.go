package services

import (
	modelos "ecommerce/models"
	repositorios "ecommerce/repositories"
)

type UsuarioService struct {
	repository *repositorios.UsuarioRepository
}

func NuevoUsuarioService() *UsuarioService {

	return &UsuarioService{

		repository: repositorios.NuevoUsuarioRepository(),
	}

}

func (s *UsuarioService) ObtenerTodos() []modelos.Usuario {

	return s.repository.ObtenerTodos()

}

func (s *UsuarioService) ObtenerPorID(id int) (*modelos.Usuario, error) {

	return s.repository.ObtenerPorID(id)

}

func (s *UsuarioService) Crear(usuario modelos.Usuario) modelos.Usuario {

	return s.repository.Crear(usuario)

}

func (s *UsuarioService) Actualizar(id int, usuario modelos.Usuario) error {

	return s.repository.Actualizar(id, usuario)

}

func (s *UsuarioService) Eliminar(id int) error {

	return s.repository.Eliminar(id)

}
