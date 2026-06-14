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
