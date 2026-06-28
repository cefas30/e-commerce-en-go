package handlers

import (
	"net/http"
	"strconv"

	modelos "ecommerce/models"
	servicios "ecommerce/services"
	"ecommerce/validators"

	"ecommerce/utils"

	"github.com/gin-gonic/gin"
)

var productoService = servicios.NuevoProductoService()

// GET /productos
func ObtenerProductos(c *gin.Context) {
	productos := productoService.ObtenerTodos()
	c.JSON(http.StatusOK, productos)
}

// GET /productos/:id
func ObtenerProducto(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	producto, err := productoService.ObtenerPorID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, producto)
}

// POST /productos
func CrearProducto(c *gin.Context) {

	var producto modelos.Producto

	if err := c.ShouldBindJSON(&producto); err != nil {

		utils.Error(c, http.StatusBadRequest, "Datos inválidos")

		return

	}

	if err := validators.ValidateProduct(producto); err != nil {

		utils.Error(c, http.StatusBadRequest, err.Error())

		return

	}

	nuevo := productoService.Crear(producto)

	utils.Created(c, "Producto creado correctamente", nuevo)

}

// PUT /productos/:id
func ActualizarProducto(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})

		return

	}

	var producto modelos.Producto

	if err := c.ShouldBindJSON(&producto); err != nil {
		if err := validators.ValidateProduct(producto); err != nil {

			utils.Error(c, http.StatusBadRequest, err.Error())

			return

		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Datos inválidos",
		})

		return

	}

	err = productoService.Actualizar(id, producto)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return

	}

	utils.Success(c, "Producto actualizado correctamente", producto)

}

// DELETE /productos/:id
func EliminarProducto(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})

		return

	}

	err = productoService.Eliminar(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})

		return

	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Producto eliminado",
	})

}
