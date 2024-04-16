package service

import (
	"ozinshe/pkg/entity"
	"ozinshe/pkg/helpers"
	"time"
)

type MovieService interface {
	CreateMovie(movie *entity.Movie) error
	GetAllMovies(limit, offset int) ([]entity.Movie, error)
	GetMovieById(id int) (*entity.Movie, error)
	UpdateMovieById(movie *entity.Movie) error
	DeleteMovieById(movieId int) error
	DeleteMovieGenreByMovieId(movieId int) error
	GetMovieSeason(movieId, seasonId int) ([]string, error)
	GetMovieSeries(movieId, seriesId, seasonId int) (string, error)
	GetMovieMainsByCategory(category, limit, offset int) ([]entity.MovieMain, error)
}

func (s *Service) CreateMovie(movie *entity.Movie) error {
	//if err:= s.repo.CreateMovie(movie); err != nil {
	//
	//}
	posterLink := helpers.GenerateRandomKey(entity.PicturesLinkNameLength)
	movie.PosterLink = posterLink
	movie.CreatedDate = time.Now()
	movie.LastModifiedDate = movie.CreatedDate
	movieId, err := s.repo.CreateMovie(movie)
	if err != nil {
		return err
	}
	if err = s.repo.CreateMovieCategories(movieId, movie.CategoryIDs); err != nil {
		return err
	}
	if err = s.repo.CreateMovieCategoryAges(movieId, movie.CategoryAgeIDs); err != nil {
		return err
	}
	if err = s.repo.CreateMovieGenres(movieId, movie.GenreIDs); err != nil {
		return err
	}
	for i := 0; i < 3; i++ {
		movie.ScreenshotLinks = append(movie.ScreenshotLinks, helpers.GenerateRandomKey(entity.PicturesLinkNameLength))
		videoLink := helpers.GenerateRandomKey(entity.PicturesLinkNameLength)
		video := entity.Video{Link: videoLink, SeasonId: 1, SeriesNumber: i}
		movie.Videos = append(movie.Videos, video)
	}
	if err = s.repo.CreateMovieScreenshots(movieId, movie.ScreenshotLinks); err != nil {
		return err
	}
	if err = s.repo.CreateMovieVideos(movieId, movie.Videos); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAllMovies(limit, offset int) ([]entity.Movie, error) {
	return s.repo.GetMoviesByPage(limit, offset)
}

func (s *Service) GetMovieById(id int) (*entity.Movie, error) {
	return s.repo.GetMovieById(id)
}

func (s *Service) UpdateMovieById(movie *entity.Movie) error {
	return s.repo.UpdateMovieById(movie)
}

func (s *Service) DeleteMovieById(movieId int) error {
	return s.repo.DeleteMovieById(movieId)
}

func (s *Service) DeleteMovieGenreByMovieId(movieId int) error {
	return s.repo.DeleteMovieGenresByMovieID(movieId)
}

func (s *Service) GetMovieSeason(movieId, seasonId int) ([]string, error) {
	return s.repo.GetMovieSeason(movieId, seasonId)
}

func (s *Service) GetMovieSeries(movieId, seriesId, seasonId int) (string, error) {
	return s.repo.GetMovieSeries(movieId, seriesId, seasonId)
}

func (s *Service) GetMovieMainsByCategory(category, limit, offset int) ([]entity.MovieMain, error) {
	movieIds, err := s.repo.GetMovieIdByCategory(category, limit, offset)
	if err != nil {
		s.log.Print("error during GetMovieMainsByCategory(service)\n")
		return nil, err
	}
	movieMains, err := s.repo.GetMovieMainByMovieIds(movieIds)
	if err != nil {
		s.log.Print("error during GetMovieMainsByCategory(service)\n")
		return nil, err
	}
	return movieMains, err
}
