package repositorios

import (
	"ecommerce/modelos"
	"errors"
)

type UsuarioMemoria struct {
	usuarios []*modelos.Usuario
}

// instancia del repositorio de usuarios en memoria
func NewUsrMemoria() *UsuarioMemoria {
	return &UsuarioMemoria{
		usuarios: []*modelos.Usuario{},
	}
}

// guardar un nuevo usuario en memoria
func (repo *UsuarioMemoria) Guardar(usuario *modelos.Usuario) error {
	for _, u := range repo.usuarios {
		if u.GetEmail() == usuario.GetEmail() {
			return errors.New("el email ya está registrado")
		}
	}
	repo.usuarios = append(repo.usuarios, usuario)
	return nil
}

// buscar un usuario por su ID
func (repo *UsuarioMemoria) ObtenerPorID(id int) (*modelos.Usuario, error) {
	for _, u := range repo.usuarios {
		if u.GetID() == id {
			return u, nil
		}
	}
	return nil, errors.New("usuario no encontrado")
}

// listar todos los usuarios
func (repo *UsuarioMemoria) ListarUsuarios() ([]*modelos.Usuario, error) {
	return repo.usuarios, nil
}
