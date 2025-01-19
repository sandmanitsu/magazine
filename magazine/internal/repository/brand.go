package repository

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

const (
	table = "brand"
)

type IBrand interface {
	Brands(where []string, values []interface{}, offset, limit int) ([]Brand, error)
	Brand(login string) Brand
	Create(data Brand) (int, error)
}

type Brand struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Login    string `json:"login"`
}

type BrandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) *BrandRepository {
	return &BrandRepository{
		db: db,
	}
}

// Creating new brand
func (r *BrandRepository) Create(data Brand) (int, error) {
	result := r.db.Table(table).Create(&data)
	if result.Error != nil {
		return 0, result.Error
	}

	return data.ID, nil
}

// Return list of brands with filter
func (r *BrandRepository) Brands(where []string, values []interface{}, offset, limit int) ([]Brand, error) {
	var brand []Brand

	result := r.db.Table(table).Where(strings.Join(where, " AND "), values...).Limit(limit).Offset(offset).Find(&brand)
	if result.Error != nil {
		return nil, fmt.Errorf("error: db select error")
	}

	return brand, nil
}

// Return brand by login and password
func (r *BrandRepository) Brand(login string) Brand {
	var brand Brand

	r.db.Raw("SELECT * FROM brand WHERE login = ?", login).Scan(&brand)

	return brand
}
