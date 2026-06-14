package servicios

// funciones relacionadas con la lógica de negocio de los productos
func calcularIVA(precio float64) float64 {
	const tasaIVA = 0.15
	return precio * tasaIVA
}

// calcularPrecioConIVA calcula el precio total incluyendo el IVA
func calcularPrecioConIVA(precio float64) float64 {
	return precio + calcularIVA(precio)
}
