package service

import (
	"magazine/internal/repository"
)

type Services struct {
	Items IItem
}

func NewService(repos *repository.Repositories) *Services {
	return &Services{
		Items: NewItemService(repos.Items),
	}
}
