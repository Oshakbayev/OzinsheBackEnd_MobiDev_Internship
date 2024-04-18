package service

import (
	"log"
	"ozinshe/pkg/repository"
)

type SvcInterface interface {
	UserService
	MovieService
	CategoryService
	CategoryAgeService
	GenreSerivce
	UserProfileService
	FavoriteMovieService
}

type Service struct {
	log  *log.Logger
	repo repository.RepoInterface
}

func CreateService(repo repository.RepoInterface, logger *log.Logger) SvcInterface {
	return &Service{repo: repo, log: logger}
}
