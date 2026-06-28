package config

import (
	"ecommerce/repositories"
	"ecommerce/services"
)

var ProductoRepository = repositories.NuevoProductoRepository()
var UsuarioRepository = repositories.NuevoUsuarioRepository()

var ProductoService = services.NuevoProductoService()
var UsuarioService = services.NuevoUsuarioService()
