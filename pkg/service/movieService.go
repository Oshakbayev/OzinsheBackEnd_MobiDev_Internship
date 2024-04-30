package service

import (
	"errors"
	"mime/multipart"
	"ozinshe/pkg/entity"
	"time"
)

type MovieService interface {
	CreateMovie(*entity.Movie, *multipart.Form) error
	GetAllMovies(int, int, int) ([]entity.Movie, error)
	GetMovieById(int, int) (*entity.Movie, error)
	UpdateMovieById(*entity.Movie) error
	DeleteMovieById(int) error
	DeleteMovieGenreByMovieId(int) error
	GetMovieSeason(int, int) ([]string, error)
	GetMovieSeries(int, int, int) (string, error)
	GetMovieMainsByTitle(int, string) ([]entity.MovieMain, error)
}

func (s *Service) CreateMovie(movie *entity.Movie, formData *multipart.Form) error {
	movie.CreatedDate = time.Now()
	movie.LastModifiedDate = movie.CreatedDate
	movieId, err := s.repo.CreateMovie(movie)
	if err != nil {
		return err
	}
	var movieMain entity.MovieMain
	movieMain.MovieId = movieId
	movieMain.MovieName = movie.Name
	movieMain.PosterLink = movie.PosterLink
	movieMain.MovieYear = movie.Year
	if err = s.repo.CreateMovieMain(&movieMain); err != nil {
		return err
	}
	if err = s.repo.CreateMovieCategories(movieId, movie.CategoryIDs); err != nil {
		if err2 := s.DeleteMovieById(movieId); err2 != nil {
			return errors.New("error in movie delete: " + err2.Error())
		}
		return err
	}
	if err = s.repo.CreateMovieCategoryAges(movieId, movie.CategoryAgeIDs); err != nil {
		if err2 := s.DeleteMovieById(movieId); err2 != nil {
			return errors.New("error in movie delete: " + err2.Error())
		}
		return err
	}
	if err = s.repo.CreateMovieGenres(movieId, movie.GenreIDs); err != nil {
		return err
	}
	//if movie.ScreenshotLinks, err = s.UploadFile(ScreenshotHeaders); err != nil {
	//	return err
	//}
	//if err := s.AddSeason(movieId, VideoHeaders); err != nil {
	//	if err2 := s.DeleteMovieById(movieId); err2 != nil {
	//		return errors.New("error in movie delete: " + err2.Error())
	//	}
	//	return err
	//}
	for i, link := range movie.VideoLinks {
		var video entity.Video
		video.Link = link
		video.SeriesNumber = i + 1
		video.SeasonId = 1
		movie.Videos = append(movie.Videos, video)
	}
	if err := s.repo.CreateMovieVideos(movieId, movie.Videos); err != nil {
		if err2 := s.DeleteMovieById(movieId); err2 != nil {
			return errors.New("error in delete: " + err2.Error())
		}
		return err
	}

	if err = s.repo.CreateMovieScreenshots(movieId, movie.ScreenshotLinks); err != nil {
		if err2 := s.DeleteMovieById(movieId); err2 != nil {
			return errors.New("error in movie delete: " + err2.Error())
		}
		return err
	}
	return nil
}

func (s *Service) GetAllMovies(userid, limit, offset int) ([]entity.Movie, error) {
	return s.repo.GetMoviesByPage(userid, limit, offset)
}

func (s *Service) GetMovieById(userId, id int) (*entity.Movie, error) {
	return s.repo.GetMovieById(id, userId)
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

func (s *Service) GetMovieMainsByTitle(userId int, title string) ([]entity.MovieMain, error) {
	//title = strings.ToLower(title)
	return s.repo.GetMovieMainsByTitle(userId, title)
}
