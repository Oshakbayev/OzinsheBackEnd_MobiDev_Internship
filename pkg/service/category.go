package service

import "ozinshe/pkg/entity"

type CategoryService interface {
	GetCategoryIdByName(categoryName string) (int, error)
	GetAllCategories() ([]entity.Category, error)
}

func (s *Service) GetCategoryIdByName(categoryName string) (int, error) {
	return s.repo.GetCategoryIdByName(categoryName)
}

func (s *Service) GetAllCategories() ([]entity.Category, error) {
	return s.repo.GetAllCategories()
}
