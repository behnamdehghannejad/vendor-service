package model

import "time"

type ProductModel struct {
	ID          int       `gorm:"primaryKey"`
	Name        string    `gorm:"size:255"`
	Description string    `gorm:"size:255"`
	Active      bool      `gorm:"default:true"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (ProductModel) TableName() string {
	return "products"
}
