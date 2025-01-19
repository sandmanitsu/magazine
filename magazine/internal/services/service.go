package service

import (
	"magazine/internal/repository"
	"magazine/pkg/hash"
	"magazine/pkg/jwt"
)

type Services struct {
	Brand IBrand
	Items IItem
}

type Deps struct {
	Hasher         hash.IHasher
	JWTManager     jwt.JWTManager
	AccessTokenTTL int
}

func NewService(repos *repository.Repositories, deps Deps) *Services {
	return &Services{
		Items: NewItemService(repos.Items),
		Brand: NewBrandService(repos.Brands, deps.Hasher, deps.JWTManager, deps.AccessTokenTTL),
	}
}
