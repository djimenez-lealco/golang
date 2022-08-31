package saludar

import (
	saludaradapter "saludarAdapter"

	"identity.leal.co/domain/ports"
)

func SaludarFactory() ports.ISaludar {
	return &saludaradapter.Saludador{}
}
