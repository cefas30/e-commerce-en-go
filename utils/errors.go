package utils

import "errors"

var (
	ErrorProductoNoEncontrado = errors.New("producto no encontrado")
)
var (
	ErrorProductoSinStock = errors.New("producto sin stock")
)
var (
	ErrCantidadInvalida = errors.New("cantidad inválida")
)
var (
	ErrPedidoVacío = errors.New("pedido vacío")
)

var (
	ErrPedidoConMultiplesItems = errors.New("el pedido solo admite un producto por ahora")
)

var (
	ErrPedidoNoExiste = errors.New("pedido no existe")
)
