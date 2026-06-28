package interfaces

import modelos "ecommerce/models"

type UsuarioRepository interface {
	ObtenerTodos() []modelos.Usuario
	ObtenerPorID(id int) (*modelos.Usuario, error)
	Crear(usuario modelos.Usuario) modelos.Usuario
	Actualizar(id int, usuario modelos.Usuario) error
	Eliminar(id int) error
}
