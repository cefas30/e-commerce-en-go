package handlers

import (
	"ecommerce/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var carritoService = services.NuevoCarritoService()

func ObtenerCarritos(c *gin.Context) {

	c.JSON(http.StatusOK, carritoService.ObtenerTodos())

}

func CrearCarrito(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("usuario"))

	carrito := carritoService.Crear(id)

	c.JSON(http.StatusCreated, carrito)

}
