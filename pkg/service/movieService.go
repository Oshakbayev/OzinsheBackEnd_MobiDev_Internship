package service

import (
	"mime/multipart"
	"ozinshe/pkg/entity"
	"ozinshe/pkg/helpers"
	"time"
)

type MovieService interface {
	CreateMovie(*entity.Movie, *multipart.Form) error
	GetAllMovies(int, int) ([]entity.Movie, error)
	GetMovieById(int) (*entity.Movie, error)
	UpdateMovieById(*entity.Movie) error
	DeleteMovieById(int) error
	DeleteMovieGenreByMovieId(int) error
	GetMovieSeason(int, int) ([]string, error)
	GetMovieSeries(int, int, int) (string, error)
	GetMovieMainsByTitle(string) ([]entity.MovieMain, error)
}

func (s *Service) CreateMovie(movie *entity.Movie, formData *multipart.Form) error {
	screenshotsDirectoryName := movie.Name + "/Screenshots"
	videoDirectoryName := movie.Name + "/Video"
	PostersDirectoryName := movie.Name + "/Poster"
	ScreenshotHeaders := formData.File["screenshots"]
	VideoHeaders := formData.File["video"]
	PosterHeaders := formData.File["poster"]
	movie.CreatedDate = time.Now()
	movie.LastModifiedDate = movie.CreatedDate
	movieId, err := s.repo.CreateMovie(movie)
	if err != nil {
		return err
	}
	var movieMain entity.MovieMain
	movieMain.MovieId = movieId
	movieMain.MovieName = movie.Name
	if err = s.repo.CreateMovieMain(&movieMain); err != nil {
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
	if movie.ScreenshotLinks, err = s.UploadFile(ScreenshotHeaders, screenshotsDirectoryName); err != nil {
		return err
	}
	if movie.VideoLinks, err = s.UploadFile(VideoHeaders, videoDirectoryName); err != nil {
		return err
	}
	posters, err := s.UploadFile(PosterHeaders, PostersDirectoryName)
	movie.PosterLink = posters[0]
	movieMain.PosterLink = posters[0]
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

func (s *Service) GetMovieMainsByTitle(title string) ([]entity.MovieMain, error) {
	//title = strings.ToLower(title)
	return s.repo.GetMovieMainsByTitle(title)
}

func (s *Service) UploadFile(fileHeaders []*multipart.FileHeader, directoryName string) ([]string, error) {
	var fileNames []string
	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			s.log.Printf("\nError UploadFile(service) in opening files : %s\n", err.Error())
			return nil, err
		}
		fileName := helpers.GenerateRandomKey(entity.PicturesLinkNameLength)
		fileNames = append(fileNames, fileName)
		if err := s.bc.UploadFile(entity.BucketName, directoryName+"/"+fileName, file); err != nil {
			s.log.Printf("error during upload file in UploadFile(Service) %s", err.Error())
			return nil, err
		}
		err = file.Close()
		if err != nil {
			s.log.Printf("\nError UploadFile(service) in closing files : %s\n", err.Error())
			return nil, err
		}
		//movie.ScreenshotLinks = append(movie.ScreenshotLinks, helpers.GenerateRandomKey(entity.PicturesLinkNameLength))
		//videoLink := helpers.GenerateRandomKey(entity.PicturesLinkNameLength)
		//video := entity.Video{Link: videoLink, SeasonId: 1, SeriesNumber: i}
		//movie.Videos = append(movie.Videos, video)
	}
	return fileNames, nil
}
