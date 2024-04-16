package service

import "ozinshe/pkg/entity"

type GenreSerivce interface {
	GetAllGenres() ([]entity.Category, error)
}

func (s *Service) GetAllGenres() ([]entity.Category, error) {
	return s.repo.GetAllGenres()
}
