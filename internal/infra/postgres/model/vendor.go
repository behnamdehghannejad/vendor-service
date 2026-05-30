package model

import "time"

type VendorModel struct {
	ID        int       `gorm:"primaryKey"`
	Code      string    `gorm:"size:50"`
	Name      string    `gorm:"size:200"`
	Email     string    `gorm:"size:100"`
	Phone     string    `gorm:"size:20"`
	Address   string    `gorm:"size:500"`
	Active    bool      `gorm:"default:true"`
	CreatedAt time.Time `gorm:column"created_at"`
	UpdatedAt time.Time `gorm:column"updated_at"`
}

func (VendorModel) TableName() string {
	return "vendors"
}
