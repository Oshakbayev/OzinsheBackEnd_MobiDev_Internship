package service

import (
	"errors"
	"ozinshe/pkg/entity"
	"ozinshe/pkg/helpers"
)

type VideoService interface {
	AddSeason(int, []string) error
	AddSeries(int, int, []string) error
	DeleteMovieSeason(int, int) error
	DeleteMovieSeries(int, int, int) error
	GetVideoDirectoryLinkByMovieId(int) (string, error)
	UpdateSeries(int, int, int, string) error
}

func (s *Service) AddSeason(movieId int, videoLinks []string) error {
	seasonId, err := s.repo.GetMaxSeason(movieId)
	if err != nil {
		return err
	}
	seasonId = seasonId + 1
	var videos []entity.Video
	for i, link := range videoLinks {
		var video entity.Video
		video.Link = link
		video.SeasonId = seasonId
		video.SeriesNumber = i + 1
		videos = append(videos, video)
	}
	if err := s.repo.CreateMovieVideos(movieId, videos); err != nil {
		if err2 := s.DeleteMovieById(movieId); err2 != nil {
			return errors.New("error in delete: " + err2.Error())
		}
		return err
	}
	return nil
}

func (s *Service) AddSeries(movieId, seasonId int, videoLinks []string) error {
	seriesId, err := s.repo.GetMaxSeriesInSeason(movieId, seasonId)
	if err != nil {
		return err
	}
	seriesId = seriesId + 1

	var videos []entity.Video
	for _, link := range videoLinks {
		var video entity.Video
		video.Link = link
		video.SeasonId = seasonId
		video.SeriesNumber = seriesId
		seriesId++
		videos = append(videos, video)
	}
	if err := s.repo.CreateMovieVideos(movieId, videos); err != nil {
		if err2 := s.DeleteMovieById(movieId); err2 != nil {
			return errors.New("error in delete: " + err2.Error())
		}
		return err
	}
	return nil
}

func (s *Service) DeleteMovieSeason(movieId, seasonId int) error {
	linksArr, err := s.repo.GetMovieSeason(movieId, seasonId)
	if err != nil {
		s.log.Println("error in DeleteMovieSeason(Service)")
		return err
	}
	for _, link := range linksArr {
		err = helpers.DeleteFile(link)
		if err != nil {
			s.log.Println("error in DeleteMovieSeason(Service)")
			return err
		}
	}
	return s.repo.DeleteMovieSeason(movieId, seasonId)
}

func (s *Service) DeleteMovieSeries(movieId, seasonId int, seriesId int) error {
	link, err := s.repo.GetMovieSeries(movieId, seriesId, seasonId)
	if err != nil {
		s.log.Println("error in DeleteMovieSeries(Service)")
		return err
	}
	err = helpers.DeleteFile(link)
	if err != nil {
		s.log.Println("error in DeleteMovieSeries(Service)")
		return err
	}
	return s.repo.DeleteMovieSeries(movieId, seasonId, seriesId)
}

func (s *Service) GetVideoDirectoryLinkByMovieId(movieId int) (string, error) {
	return s.repo.GetVideoDirectoryLinkByMovieId(movieId)
}

func (s *Service) UpdateSeries(movieId int, seasonId int, seriesId int, link string) error {
	return s.repo.UpdateSeries(movieId, seasonId, seriesId, link)
}
