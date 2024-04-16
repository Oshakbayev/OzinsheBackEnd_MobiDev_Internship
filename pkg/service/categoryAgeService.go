package service

import "ozinshe/pkg/entity"

type CategoryAgeService interface {
	GetAllGenres() ([]entity.Category, error)
}

func (s *Service) GetAllCategoryAges() ([]entity.Category, error) {
	return s.repo.GetAllCategoryAges()
}
