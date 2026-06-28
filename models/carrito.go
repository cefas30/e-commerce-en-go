package modelos

type ItemCarrito struct {
	Producto Producto `json:"producto"`
	Cantidad int      `json:"cantidad"`
}

type Carrito struct {
	ID        int           `json:"id"`
	UsuarioID int           `json:"usuario_id"`
	Items     []ItemCarrito `json:"items"`
	Total     float64       `json:"total"`
}
