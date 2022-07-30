package injectors

import (
	"hexagonal/aplication/useCases"
	"hexagonal/infraestructure/adapters"
)

type SaludarInjector struct{}

func (saludarInjector *SaludarInjector) Hello() string {
	return useCases.SaludarFactory(new(adapters.MensajeAdapter)).Saludar().Mensaje
}

func NewSaludarInjector() *SaludarInjector {
	return &SaludarInjector{}
}
