package repository

import (
	"context"
	"github.com/lib/pq"
	"ozinshe/pkg/entity"
)

type VideoRepo interface {
	CreateMovieVideos(int, []entity.Video) error
	GetMaxSeason(int) (int, error)
	GetMaxSeriesInSeason(int, int) (int, error)
	DeleteMovieSeason(int, int) error
	DeleteMovieSeries(int, int, int) error
	GetVideoDirectoryLinkByMovieId(int) (string, error)
}

func (r *RepoStruct) CreateMovieVideos(movieID int, videos []entity.Video) error {
	var values []interface{}
	query := `INSERT INTO video (movie_id, link, season_num, series_num) 
              SELECT $1, unnest($2::text[]), unnest($3::int[]), unnest($4::int[])`
	links := make([]string, len(videos))
	seasonIDs := make([]int, len(videos))
	seriesNums := make([]int, len(videos))
	for i, v := range videos {
		links[i] = v.Link
		seasonIDs[i] = v.SeasonId
		seriesNums[i] = v.SeriesNumber
	}
	values = append(values, movieID, pq.Array(links), pq.Array(seasonIDs), pq.Array(seriesNums))
	_, err := r.db.Exec(context.Background(), query, values...)
	if err != nil {
		r.log.Printf("error in CreateMovieVideos(repository): %s", err.Error())
	}
	return err

}

func (r *RepoStruct) GetMaxSeason(movieId int) (int, error) {
	var seasonId int
	query := `SELECT max(season_num) FROM video WHERE movie_id = $1`
	err := r.db.QueryRow(context.Background(), query, movieId).Scan(&seasonId)
	if err != nil {
		r.log.Printf("error in GetMaxSeason(Repository): %v", err)
		return seasonId, err
	}
	return seasonId, nil
}

func (r *RepoStruct) GetMaxSeriesInSeason(movieId int, seasonId int) (int, error) {
	var seriesId int
	query := `SELECT max(series_num) FROM video WHERE movie_id = $1 AND season_num = $2`
	err := r.db.QueryRow(context.Background(), query, movieId, seasonId).Scan(&seriesId)
	if err != nil {
		r.log.Printf("error in GetMaxSeriesInSeason(Repository): %v", err)
		return seriesId, err
	}
	return seriesId, nil
}

func (r *RepoStruct) DeleteMovieSeason(movieId int, seasonId int) error {
	query := `DELETE FROM video WHERE movie_id = $1 AND season_num = $2`
	_, err := r.db.Exec(context.Background(), query, movieId, seasonId)
	if err != nil {
		r.log.Printf("error in DeleteMovieSeason(repository): %s", err.Error())
	}
	return err
}

func (r *RepoStruct) DeleteMovieSeries(movieId int, seasonId int, seriesId int) error {
	query := `DELETE FROM video WHERE movie_id = $1 AND season_num = $2 AND series_num = $3`
	_, err := r.db.Exec(context.Background(), query, movieId, seasonId, seriesId)
	if err != nil {
		r.log.Printf("error in DeleteMovieSeries(repository): %s", err.Error())
	}
	return err
}

func (r *RepoStruct) GetVideoDirectoryLinkByMovieId(movieId int) (string, error) {
	query := `SELECT video_directory_link FROM movie WHERE id = $1`
	var videoDirectoryLink string
	err := r.db.QueryRow(context.Background(), query, movieId).Scan(&videoDirectoryLink)
	if err != nil {
		r.log.Printf("error in GetVideoDirectoryLinkByMovieId(repository): %s", err.Error())
	}
	return videoDirectoryLink, err
}
