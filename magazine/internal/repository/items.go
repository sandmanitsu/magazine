package repository

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type IItems interface {
	Items(where []string, values []interface{}, offset, limit int) ([]Item, error)
}

type Item struct {
	ID          int     `json:"id"`
	BrandId     int     `json:"brand_id"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Size        *string `json:"size"`
	Price       uint    `json:"price"`
	Discount    *int    `json:"discount"`
	Description string  `json:"description"`
	Gender      *string `json:"gender"`
	Status      uint    `json:"status"`
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

// Return list of items
func (r *ItemRepository) Items(where []string, values []interface{}, offset, limit int) ([]Item, error) {
	var item []Item

	result := r.db.Where(strings.Join(where, " AND "), values...).Limit(limit).Offset(offset).Find(&item)
	if result.Error != nil {
		return nil, fmt.Errorf("error: db select error")
	}

	return item, nil
}
