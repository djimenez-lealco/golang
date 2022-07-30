package adapters

import (
	entities "hexagonal/domain/Entities"
)

type MensajeAdapter struct{}

func (ms MensajeAdapter) GetMensaje() *entities.Mensaje {
	mensaje := entities.Mensaje{Mensaje: "Hello, World en un package saludar! con inyeccion 2 👍🏻"}
	return &mensaje
}
