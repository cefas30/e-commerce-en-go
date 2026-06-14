package modelos

import "errors"

// Usuario representa un usuario en el sistema
type Usuario struct {
	ID       int
	Nombre   string
	Email    string
	Password string
}

// constructor para crear un nuevo usuario
func NuevoUsuario(id int, nombre, email, password string) *Usuario {
	return &Usuario{
		ID:       id,
		Nombre:   nombre,
		Email:    email,
		Password: password,
	}
}

// getters para acceder a los campos del usuario
func (u *Usuario) GetID() int {
	return u.ID
}
func (u *Usuario) GetNombre() string {
	return u.Nombre
}
func (u *Usuario) GetEmail() string {
	return u.Email
}

// setter para actualizar la contraseña del usuario
func (u *Usuario) SetNombre(newName string) error {
	if newName == "" {
		return errors.New("el nombre no puede estar vacío")
	}
	u.Nombre = newName
	return nil
}
func (u *Usuario) SetEmail(newEmail string) error {
	if newEmail == "" {
		return errors.New("el email no puede estar vacío")
	}
	u.Email = newEmail
	return nil
}
