package modelos

type EstadoPedido string

const (
	Pendiente EstadoPedido = "PENDIENTE"

	Pagado EstadoPedido = "PAGADO"

	Enviado EstadoPedido = "ENVIADO"

	Entregado EstadoPedido = "ENTREGADO"

	Cancelado EstadoPedido = "CANCELADO"
)

type Pedido struct {
	ID int `json:"id"`

	Usuario Usuario `json:"usuario"`

	Items []ItemCarrito `json:"items"`

	Total float64 `json:"total"`

	Estado EstadoPedido `json:"estado"`
}
