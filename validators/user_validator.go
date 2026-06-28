package validators

import (
	"errors"
	"strings"

	modelos "ecommerce/models"
)

func ValidateUser(user modelos.Usuario) error {

	if strings.TrimSpace(user.Nombre) == "" {

		return errors.New("el nombre es obligatorio")

	}

	if strings.TrimSpace(user.Correo) == "" {

		return errors.New("el correo es obligatorio")

	}

	if len(user.Password) < 6 {

		return errors.New("la contraseña debe tener mínimo 6 caracteres")

	}

	return nil

}
