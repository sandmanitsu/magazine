package service

import (
	"fmt"
	"magazine/internal/repository"
	"magazine/pkg/hash"
	"net/url"
	"strconv"
)

type IBrand interface {
	Brands(params url.Values) ([]repository.Brand, error)
	SignUp(data BrandSignUpData) error
}

type BrandService struct {
	repos  repository.IBrand
	hasher hash.IHasher
}

func NewBrandService(repos repository.IBrand, hasher hash.IHasher) *BrandService {
	return &BrandService{
		repos:  repos,
		hasher: hasher,
	}
}

type BrandSignUpData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Login    string `json:"login"`
}

// todo. Добавить валидацию
func (s *BrandService) SignUp(data BrandSignUpData) error {
	hashedPassword, err := s.hasher.Hash(data.Password)
	if err != nil {
		return fmt.Errorf("error: creating pass")
	}

	brand := repository.Brand{
		Name:     data.Name,
		Email:    data.Email,
		Password: hashedPassword,
		Login:    data.Login,
	}

	id, err := s.repos.Create(brand)
	fmt.Println(id)
	if err != nil {
		return fmt.Errorf("error: creating brand")
	}

	return nil
}

func (s *BrandService) Brands(params url.Values) ([]repository.Brand, error) {
	offset := 0
	limit := 100
	var whereStatements []string
	var values []interface{}
	for param, value := range params {
		if param == "offset" && len(value) == 1 {
			var err error
			offset, err = strconv.Atoi(value[0])
			if err != nil {
				return nil, fmt.Errorf("error: incorrect OFFSET param")
			}
		} else if param == "limit" && len(value) == 1 {
			var err error
			limit, err = strconv.Atoi(value[0])
			if err != nil {
				return nil, fmt.Errorf("error: incorrect LIMIT param")
			}
		} else {
			// ?? мб нужно дописать логику построение условий where, сейчас все условия делаются через IN и values хранит []interface{}{[]string}
			whereStatements = append(whereStatements, fmt.Sprintf("%s IN ?", param))
			values = append(values, []string{value[0]})
		}
	}

	item, err := s.repos.Brands(whereStatements, values, offset, limit)
	if err != nil {
		return nil, err
	}

	return item, nil
}
