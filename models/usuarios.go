package modelos

type Usuario struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Correo   string `json:"correo"`
	Password string `json:"password"`
	Rol      string `json:"rol"`
}
