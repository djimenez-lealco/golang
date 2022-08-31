package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	FirstName string
	LastName  string
	Tel       string
	Orders    []Order
}
