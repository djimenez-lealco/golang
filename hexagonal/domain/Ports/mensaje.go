package ports

import entities "hexagonal/domain/Entities"

type IMensajePort interface {
	GetMensaje() *entities.Mensaje
}
