package service

import (
	"log"
	"ozinshe/pkg/bucket"
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
	bc   bucket.BucketInterface
	repo repository.RepoInterface
}

func CreateService(repo repository.RepoInterface, logger *log.Logger, bc bucket.BucketInterface) SvcInterface {
	return &Service{repo: repo, bc: bc, log: logger}
}
