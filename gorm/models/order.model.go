package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ClientID   uuid.UUID
	Ref        string
	Orderitems []Orderitem
}
