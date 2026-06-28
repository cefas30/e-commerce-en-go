package repositories

import (
	modelos "ecommerce/models"
	"errors"
)

type ProductoRepository struct {
	productos []modelos.Producto
}

func NuevoProductoRepository() *ProductoRepository {

	return &ProductoRepository{

		productos: []modelos.Producto{

			{
				ID:     1,
				Nombre: "Laptop Lenovo",
				Precio: 850,
				Stock:  10,
			},

			{
				ID:     2,
				Nombre: "Mouse Logitech",
				Precio: 25,
				Stock:  50,
			},
		},
	}
}

func (r *ProductoRepository) ObtenerTodos() []modelos.Producto {
	return r.productos
}

func (r *ProductoRepository) ObtenerPorID(id int) (*modelos.Producto, error) {

	for i := range r.productos {

		if r.productos[i].ID == id {
			return &r.productos[i], nil
		}

	}

	return nil, errors.New("producto no encontrado")
}

func (r *ProductoRepository) Crear(producto modelos.Producto) modelos.Producto {

	producto.ID = len(r.productos) + 1

	r.productos = append(r.productos, producto)

	return producto
}

func (r *ProductoRepository) Actualizar(id int, producto modelos.Producto) error {

	for i := range r.productos {

		if r.productos[i].ID == id {

			producto.ID = id

			r.productos[i] = producto

			return nil

		}

	}

	return errors.New("producto no encontrado")

}

func (r *ProductoRepository) Eliminar(id int) error {

	for i := range r.productos {

		if r.productos[i].ID == id {

			r.productos = append(r.productos[:i], r.productos[i+1:]...)

			return nil

		}

	}

	return errors.New("producto no encontrado")

}
