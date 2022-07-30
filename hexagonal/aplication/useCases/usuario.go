package useCases

import (
	entities "hexagonal/domain/Entities"
	"hexagonal/infraestructure/adapters"
)

type UsuarioUseCase struct {
	usuarioPort adapters.UsuarioAdapter
}

func (gu *UsuarioUseCase) Get(id string) entities.Usuario {
	return gu.usuarioPort.GetUsuario(id)
}

func UsuarioFactory(UsuarioAdapter adapters.UsuarioAdapter) *UsuarioUseCase {
	return &UsuarioUseCase{
		usuarioPort: UsuarioAdapter,
	}
}
