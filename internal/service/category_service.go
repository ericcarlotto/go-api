package service

import (
	"github.com/ericcarlotto/imersao17/goapi/internal/database"
	"github.com/ericcarlotto/imersao17/goapi/internal/entity"
)

type CategoryService struct {
	CategoryDb database.CategoryDB
}

func NewCategoryService(categoryDb database.CategoryDB) *CategoryService {
	return &CategoryService{CategoryDb: categoryDb}
}

func (cs *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := cs.CategoryDb.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (cs *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.Category{Name: name}
	_, err := cs.CategoryDb.CreateCategory(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (cs *CategoryService) GetCategory(id string) (*entity.Category, error) {
	category, err := cs.CategoryDb.GetCategoryByID(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
