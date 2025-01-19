package app

import (
	"fmt"
	"log"
	"magazine/internal/config"
	db "magazine/internal/database"
	"magazine/internal/repository"
	service "magazine/internal/services"
	"magazine/internal/transport/myrouter"
	"magazine/pkg/hash"
	"magazine/pkg/jwt"
	"strconv"

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

	jwtManager, err := jwt.NewManager(config.Auth.SecretKey)
	if err != nil {
		log.Fatalf("error: jwt manager: %s", err)
	}

	jwt_ttl, err := strconv.Atoi(config.Auth.JWT_TTL)
	if err != nil {
		log.Fatalf("error: conver jwt ttl to INT")
	}

	services := service.NewService(repository, service.Deps{
		Hasher:         hash.NewBcryptHasher(),
		JWTManager:     jwtManager,
		AccessTokenTTL: jwt_ttl,
	})

	handlers := myrouter.NewHandler(services)
	router := handlers.Init()

	return &App{
		Router: router,
		DB:     db,
	}
}
