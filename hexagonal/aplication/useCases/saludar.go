package useCases

import (
	entities "hexagonal/domain/Entities"
	ports "hexagonal/domain/Ports"
)

type SaludarUseCase struct {
	mesageService ports.IMensajePort
}

func (saludar SaludarUseCase) Saludar() *entities.Mensaje {
	return saludar.mesageService.GetMensaje()
}

func SaludarFactory(getmessage ports.IMensajePort) *SaludarUseCase {
	s := new(SaludarUseCase)
	s.mesageService = getmessage
	return s
}
