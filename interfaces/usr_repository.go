package interfaces

import "ecommerce/modelos"

type UsuarioRepository interface {
	Guardar(usuario *modelos.Usuario) error
	ObtenerPorID(id int) (*modelos.Usuario, error)
	ObtenerPorEmail(email string) (*modelos.Usuario, error)
	Actualizar(usuario *modelos.Usuario) error
	Eliminar(id int) error
}
