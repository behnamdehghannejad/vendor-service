package model

type CategoryModel struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `gorm:"size:255"`
	Active   bool   `gorm:"default:true"`
	Path     string `gorm:"size:1024"`
	ParentID *int
}

func (CategoryModel) TableName() string {
	return "categories"
}
