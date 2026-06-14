package servicios

import "strings"

// verificacion de email
func validarEmail(email string) bool {
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return false
	}
	return true
}

// verificacion de contraseña
func validarPassword(password string) bool {
	if len(password) < 6 {
		return false
	}
	return true
}
