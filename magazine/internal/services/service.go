package service

import (
	"magazine/internal/repository"
	"magazine/pkg/hash"
)

type Services struct {
	Brand IBrand
	Items IItem
}

type Deps struct {
	Hasher hash.IHasher
}

func NewService(repos *repository.Repositories, deps Deps) *Services {
	return &Services{
		Items: NewItemService(repos.Items),
		Brand: NewBrandService(repos.Brands, deps.Hasher),
	}
}
