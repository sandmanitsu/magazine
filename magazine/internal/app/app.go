package app

import (
	"fmt"
	"magazine/internal/config"
	db "magazine/internal/database"
	"magazine/internal/repository"
	service "magazine/internal/services"
	"magazine/internal/transport/myrouter"
	"magazine/pkg/hash"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func NewApp(config *config.Config) *App {
	fmt.Println("New App initializing...")

	db := db.NewPostgreInstance(&config.DB)
	repository := repository.NewRepository(db)

	services := service.NewService(repository, service.Deps{
		Hasher: hash.NewBcryptHasher(),
	})

	handlers := myrouter.NewHandler(services)
	router := handlers.Init()

	return &App{
		Router: router,
		DB:     db,
	}
}
