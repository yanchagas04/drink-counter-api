package drinks

import (
	"gorm.io/gorm"
)

type Drink struct {
	gorm.Model
	ID uint
	Quantity int
}