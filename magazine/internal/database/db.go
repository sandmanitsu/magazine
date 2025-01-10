package db

import (
	"fmt"
	"magazine/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type PostgreDB struct {
// 	db gorm.Dialector
// }

func NewPostgreInstance(config *config.DB) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBname,
		config.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
