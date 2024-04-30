package service

import (
	"errors"
	"ozinshe/pkg/entity"
)

type FavoriteMovieService interface {
	CreatFavoriteMovie(entity.Favorite) error
	DeleteFavoriteMovie(entity.Favorite) error
	GetUserFavoriteMovieIDs(int) ([]entity.MovieMain, error)
}

func (s *Service) CreatFavoriteMovie(favorite entity.Favorite) error {
	_, err := s.repo.GetFavoriteMovie(favorite)
	if err.Error() == entity.ErrNoRows {
		return s.repo.CreatFavoriteMovie(favorite)
	} else if err == nil {
		return errors.New("409")
	}
	return err
}

func (s *Service) GetUserFavoriteMovieIDs(userId int) ([]entity.MovieMain, error) {
	return s.repo.GetUserFavoriteMovieMains(userId)
}

func (s *Service) DeleteFavoriteMovie(favorite entity.Favorite) error {
	return s.repo.DeleteFavoriteMovie(favorite)
}
