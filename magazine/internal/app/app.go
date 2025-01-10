package app

import (
	"fmt"
	"magazine/internal/config"
	db "magazine/internal/database"
	"magazine/internal/repository"
	service "magazine/internal/services"
	"magazine/internal/transport/myrouter"

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

	// ??? можно сделать стреутуру Services с полем repos и полями содержащими общие данные для сервисов
	services := service.NewService(repository)

	handlers := myrouter.NewHandler(services)
	router := handlers.Init()

	return &App{
		Router: router,
		DB:     db,
	}
}
