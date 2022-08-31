package models

import "gorm.io/gorm"

type Orderitem struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Quantity  uint
}
