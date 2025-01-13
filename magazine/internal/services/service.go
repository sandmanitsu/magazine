package service

import (
	"magazine/internal/repository"
)

type Services struct {
	Items Item
	Brand IBrand
}

func NewService(repos *repository.Repositories) *Services {
	return &Services{
		Items: NewItemService(repos.Items),
		Brand: NewBrandService(repos.Brands),
	}
}
