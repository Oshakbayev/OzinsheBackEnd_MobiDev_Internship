package repository

import "context"

type VideoRepo interface {
	GetMaxSeason(int) (int, error)
	GetMaxSeriesInSeason(int, int) (int, error)
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
	query := `SELECT max(season_num) FROM video WHERE movie_id = $1 AND season_num = $2`
	err := r.db.QueryRow(context.Background(), query, movieId, seasonId).Scan(&seriesId)
	if err != nil {
		r.log.Printf("error in GetMaxSeriesInSeason(Repository): %v", err)
		return seriesId, err
	}
	return seriesId, nil
}
