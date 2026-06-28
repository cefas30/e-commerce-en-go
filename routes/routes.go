package routes

import (
	"ecommerce/handlers"

	"github.com/gin-gonic/gin"
)

func ConfigurarRutas(router *gin.Engine) {

	api := router.Group("/api")

	// Productos
	api.GET("/productos", handlers.ObtenerProductos)
	api.GET("/productos/:id", handlers.ObtenerProducto)
	api.POST("/productos", handlers.CrearProducto)
	api.PUT("/productos/:id", handlers.ActualizarProducto)
	api.DELETE("/productos/:id", handlers.EliminarProducto)

	// Usuarios
	api.GET("/usuarios", handlers.ObtenerUsuarios)
	api.GET("/usuarios/:id", handlers.ObtenerUsuario)
	api.POST("/usuarios", handlers.CrearUsuario)
	api.PUT("/usuarios/:id", handlers.ActualizarUsuario)
	api.DELETE("/usuarios/:id", handlers.EliminarUsuario)

	// Carrito
	api.GET("/carritos", handlers.ObtenerCarritos)
	api.POST("/carritos", handlers.CrearCarrito)

	// Pedidos
	api.GET("/pedidos", handlers.ObtenerPedidos)
	api.POST("/pedidos", handlers.CrearPedido)

}
