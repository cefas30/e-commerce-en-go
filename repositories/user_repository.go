package repositories

import (
	modelos "ecommerce/models"
	"errors"
)

type UsuarioRepository struct {
	usuarios []modelos.Usuario
}

func NuevoUsuarioRepository() *UsuarioRepository {

	return &UsuarioRepository{

		usuarios: []modelos.Usuario{

			{
				ID:       1,
				Nombre:   "Administrador",
				Correo:   "admin@ecommerce.com",
				Password: "123456",
				Rol:      "ADMIN",
			},

			{
				ID:       2,
				Nombre:   "Cliente",
				Correo:   "cliente@gmail.com",
				Password: "123456",
				Rol:      "CLIENTE",
			},
		},
	}

}

func (r *UsuarioRepository) ObtenerTodos() []modelos.Usuario {

	return r.usuarios

}

func (r *UsuarioRepository) ObtenerPorID(id int) (*modelos.Usuario, error) {

	for i := range r.usuarios {

		if r.usuarios[i].ID == id {

			return &r.usuarios[i], nil

		}

	}

	return nil, errors.New("usuario no encontrado")

}

func (r *UsuarioRepository) Crear(usuario modelos.Usuario) modelos.Usuario {

	usuario.ID = len(r.usuarios) + 1

	r.usuarios = append(r.usuarios, usuario)

	return usuario

}

func (r *UsuarioRepository) Actualizar(id int, usuario modelos.Usuario) error {

	for i := range r.usuarios {

		if r.usuarios[i].ID == id {

			usuario.ID = id

			r.usuarios[i] = usuario

			return nil

		}

	}

	return errors.New("usuario no encontrado")

}

func (r *UsuarioRepository) Eliminar(id int) error {

	for i := range r.usuarios {

		if r.usuarios[i].ID == id {

			r.usuarios = append(r.usuarios[:i], r.usuarios[i+1:]...)

			return nil

		}

	}

	return errors.New("usuario no encontrado")

}
