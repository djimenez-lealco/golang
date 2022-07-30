package ports

import entities "hexagonal/domain/Entities"

type IUsuarioPort interface {
	GetUsuario(id string) entities.Usuario
	//	GetUsuarios() []entities.Usuario
}
