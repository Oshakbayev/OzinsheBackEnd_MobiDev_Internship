package service

import (
	"errors"
	"ozinshe/pkg/entity"
)

type VideoService interface {
	AddSeason(int, []string) error
	AddSeries(int, int, []string) error
}

func (s *Service) AddSeason(movieId int, videoLinks []string) error {
	seasonId, err := s.repo.GetMaxSeason(movieId)
	seasonId = seasonId + 1
	var videos []entity.Video
	for i, link := range videoLinks {
		var video entity.Video
		video.Link = link
		video.SeasonId = seasonId
		video.SeriesNumber = i + 1
		videos = append(videos, video)
	}
	if err != nil {
		return err
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
