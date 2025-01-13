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

// Return list of brands with filter
func (r *BrandRepository) Brands(where []string, values []interface{}, offset, limit int) ([]Brand, error) {
	var brand []Brand

	result := r.db.Table(table).Where(strings.Join(where, " AND "), values...).Limit(limit).Offset(offset).Find(&brand)
	if result.Error != nil {
		return nil, fmt.Errorf("error: db select error")
	}

	return brand, nil
}
