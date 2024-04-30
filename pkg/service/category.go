package service

import "ozinshe/pkg/entity"

type CategoryService interface {
	GetCategoryIdByName(categoryName string) (int, error)
	GetAllCategories() ([]entity.Category, error)
	GetAllMovieMainsByCategory(int, int) ([]entity.MovieMain, error)
	GetMovieMainsByCategory(int, int, int, int) ([]entity.MovieMain, error)
}

func (s *Service) GetCategoryIdByName(categoryName string) (int, error) {
	return s.repo.GetCategoryIdByName(categoryName)
}

func (s *Service) GetAllCategories() ([]entity.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *Service) GetAllMovieMainsByCategory(userId int, categoryId int) ([]entity.MovieMain, error) {
	return s.repo.GetAllMovieMainsByCategory(userId, categoryId)
}

func (s *Service) GetMovieMainsByCategory(userId, categoryId, limit, offset int) ([]entity.MovieMain, error) {
	return s.repo.GetMovieMainsByCategory(userId, categoryId, limit, offset)
}
