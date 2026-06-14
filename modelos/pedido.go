package modelos

import "fmt"

//ESTADOS DEL PEDIDO
const (
	EstadoPendiente = "Pendiente"
	EstadoEnviado   = "Enviado"
	EstadoEntregado = "Entregado"
)

// Pedido representa un pedido realizado por un usuario
type Pedido struct {
	ID       int
	Usuario  *Usuario
	Producto *Producto
	Cantidad int
	Estado   string
}

//constructor para crear un nuevo pedido
func NuevoPedido(id int, usuario *Usuario, producto *Producto, cantidad int) *Pedido {
	return &Pedido{
		ID:       id,
		Usuario:  usuario,
		Producto: producto,
		Cantidad: cantidad,
		Estado:   EstadoPendiente,
	}
}

// getters para acceder a los campos del pedido
func (p *Pedido) GetID() int {
	return p.ID
}
func (p *Pedido) GetUsuario() *Usuario {
	return p.Usuario
}
func (p *Pedido) GetProducto() *Producto {
	return p.Producto
}
func (p *Pedido) GetCantidad() int {
	return p.Cantidad
}
func (p *Pedido) GetEstado() string {
	return p.Estado
}

// setter para actualizar el estado del pedido
func (p *Pedido) SetEstado(newEstado string) error {
	if newEstado != EstadoPendiente && newEstado != EstadoEnviado && newEstado != EstadoEntregado {
		return fmt.Errorf("estado inválido: %s", newEstado)
	}
	p.Estado = newEstado
	return nil
}

// cambiar el tipo de estado del pedido
func (p *Pedido) CambiarEstado(newEstado string) error {
	return p.SetEstado(newEstado)
}

// mostrar estado del pedido
func (p *Pedido) MostrarEstado() string {
	switch p.Estado {
	case EstadoPendiente:
		return "El pedido está pendiente de envío."

	case EstadoEnviado:
		return "El pedido ha sido enviado."
	case EstadoEntregado:
		return "El pedido ha sido entregado."
	default:
		return "Estado desconocido."
	}
}
