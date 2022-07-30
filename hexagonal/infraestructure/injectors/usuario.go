package injectors

import (
	"hexagonal/aplication/useCases"
	entities "hexagonal/domain/Entities"
	"hexagonal/infraestructure/adapters"
)

// Servicio de Usuario
type UsuarioInjector struct{}

// Construye el caso de uso para optener un usuario e inyecto las dependencias necesarias para la operación
func (usuarioService *UsuarioInjector) GetUsuario(id string) entities.Usuario {
	usuarioAdapter := adapters.UsuarioAdapter{}
	getUsuario := useCases.UsuarioFactory(usuarioAdapter)
	return getUsuario.Get(id)
}

// Entrega una instancia de Usuario Service
// Lista para invocar los metodos del servicio que inyectan las dependencias de cada caso de uso
func NewUsuarioInjector() *UsuarioInjector {
	return &UsuarioInjector{}
}
