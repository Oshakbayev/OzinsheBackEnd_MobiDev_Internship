package service

import "ozinshe/pkg/entity"

type CategoryService interface {
	GetCategoryIdByName(categoryName string) (int, error)
	GetAllCategories() ([]entity.Category, error)
	GetAllMovieMainsByCategory(categoryId int) ([]entity.MovieMain, error)
	GetMovieMainsByCategory(categoryId, limit, offset int) ([]entity.MovieMain, error)
}

func (s *Service) GetCategoryIdByName(categoryName string) (int, error) {
	return s.repo.GetCategoryIdByName(categoryName)
}

func (s *Service) GetAllCategories() ([]entity.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *Service) GetAllMovieMainsByCategory(categoryId int) ([]entity.MovieMain, error) {
	return s.repo.GetAllMovieMainsByCategory(categoryId)
}

func (s *Service) GetMovieMainsByCategory(categoryId, limit, offset int) ([]entity.MovieMain, error) {
	return s.repo.GetMovieMainsByCategory(categoryId, limit, offset)
}
