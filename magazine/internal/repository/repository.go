package repository

import "gorm.io/gorm"

type Repositories struct {
	Items  IItems
	Brands IBrand
}

func NewRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		Items:  NewItemRepository(db),
		Brands: NewBrandRepository(db),
	}
}
