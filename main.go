package main

import (
	"ecommerce/modelos"
	"ecommerce/repositorios"
	"fmt"
)

func main() {
	// Crear un nuevo repositorio de productos en memoria
	repo := repositorios.NewProMemoria()
	// Crear un nuevo producto
	producto := modelos.NuevoProducto(1, "Laptop", 999.99, 10)
	// Guardar el producto en el repositorio
	err := repo.SaveProducto(producto)
	if err != nil {
		fmt.Println("Error al guardar el producto:", err)
		return
	}
	// Listar todos los productos
	productos, err := repo.ListProductos()
	if err != nil {
		fmt.Println("Error al listar los productos:", err)
		return
	}
	fmt.Println("Productos en el repositorio:")
	for _, p := range productos {
		fmt.Printf("ID: %d, Nombre: %s, Precio: %.2f, Stock: %d\n", p.GetID(), p.GetNombre(), p.GetPrecio(), p.GetStock())
	}
}
