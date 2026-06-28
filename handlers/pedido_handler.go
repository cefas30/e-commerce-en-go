package handlers

import (
	"net/http"

	modelos "ecommerce/models"
	"ecommerce/services"

	"github.com/gin-gonic/gin"
)

var pedidoService = services.NuevoPedidoService()

func ObtenerPedidos(c *gin.Context) {

	c.JSON(http.StatusOK, pedidoService.ObtenerTodos())

}

func CrearPedido(c *gin.Context) {

	var pedido modelos.Pedido

	if err := c.ShouldBindJSON(&pedido); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"error": "json inválido",
		})

		return

	}

	nuevo := pedidoService.Crear(pedido)

	c.JSON(http.StatusCreated, nuevo)

}
