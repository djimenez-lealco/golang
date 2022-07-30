package adapters

import entities "hexagonal/domain/Entities"

type UsuarioAdapter struct {
}

func (ua *UsuarioAdapter) GetUsuario(id string) entities.Usuario {
	return entities.Usuario{
		Id:        id,
		Uuid:      "363a152b-d4a9-40aa-ba17-67e97c7337bb",
		Login:     "adgranados",
		Email:     "adgranados@gmail.com",
		Pass:      "esxdrfctgvbhyjnu124556w4yhrtsdxfgmk",
		BirthDate: "1985/10/31",
	}
}
