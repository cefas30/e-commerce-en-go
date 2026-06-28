package handlers

import (
	"net/http"
	"strconv"

	modelos "ecommerce/models"
	"ecommerce/services"
	"ecommerce/utils"
	"ecommerce/validators"

	"github.com/gin-gonic/gin"
)

var usuarioService = services.NuevoUsuarioService()

func ObtenerUsuarios(c *gin.Context) {

	c.JSON(http.StatusOK, usuarioService.ObtenerTodos())

}

func ObtenerUsuario(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	usuario, err := usuarioService.ObtenerPorID(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, usuario)

}

func CrearUsuario(c *gin.Context) {

	var usuario modelos.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "datos inválidos"})

		return

	}
	if err := validators.ValidateUser(usuario); err != nil {

		utils.Error(c, http.StatusBadRequest, err.Error())

		return

	}

	utils.Created(c, "usuario creado", usuarioService.Crear(usuario))

}

func ActualizarUsuario(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var usuario modelos.Usuario

	c.ShouldBindJSON(&usuario)

	err := usuarioService.Actualizar(id, usuario)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "usuario actualizado"})

}

func EliminarUsuario(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	err := usuarioService.Eliminar(id)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"mensaje": "usuario eliminado"})

}
