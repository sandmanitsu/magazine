package service

import (
	"fmt"
	"magazine/internal/repository"
	"net/url"
	"strconv"
)

type IItem interface {
	Items(params url.Values) ([]repository.Item, error)
}

type ItemService struct {
	repos repository.IItems
}

func NewItemService(repos repository.IItems) *ItemService {
	return &ItemService{
		repos: repos,
	}
}

func (s *ItemService) Items(params url.Values) ([]repository.Item, error) {
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

	item, err := s.repos.Items(whereStatements, values, offset, limit)
	if err != nil {
		return nil, err
	}

	return item, nil
}
