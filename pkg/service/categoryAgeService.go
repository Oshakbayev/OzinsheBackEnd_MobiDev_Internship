package service

import "ozinshe/pkg/entity"

type CategoryAgeService interface {
	GetAllCategoryAges() ([]entity.CategoryAge, error)
}

func (s *Service) GetAllCategoryAges() ([]entity.CategoryAge, error) {
	return s.repo.GetAllCategoryAges()
}
