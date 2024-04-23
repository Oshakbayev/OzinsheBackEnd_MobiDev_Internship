package repository

import (
	"context"
	"ozinshe/pkg/entity"
)

type FavoriteMovieRepo interface {
	CreatFavoriteMovie(entity.Favorite) error
	GetUserFavoriteMovieIDs(int) ([]int, error)
	DeleteFavoriteMovie(entity.Favorite) error
	GetFavoriteMovie(entity.Favorite) (entity.Favorite, error)
}

func (r *RepoStruct) CreatFavoriteMovie(favorite entity.Favorite) error {
	query := `INSERT INTO favorites (user_id,movie_id) VALUES ($1,$2)`
	_, err := r.db.Exec(context.Background(), query, favorite.UserId, favorite.MovieId)
	if err != nil {
		r.log.Printf("error in CreatFavoriteMovie(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) GetUserFavoriteMovieIDs(userId int) ([]int, error) {
	query := `SELECT movie_id FROM favorites WHERE user_id = $1`
	rows, err := r.db.Query(context.Background(), query, userId)
	if err != nil {
		r.log.Printf("error in GetUserFavoriteMovieIDs(repository):%s", err.Error())
		return nil, err
	}
	var movieIds []int
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			r.log.Printf("error in GetUserFavoriteMovieIDs(repository):%s", err.Error())
			return nil, err
		}
		movieIds = append(movieIds, id)
	}
	return movieIds, err
}

func (r *RepoStruct) DeleteFavoriteMovie(favorite entity.Favorite) error {
	query := `DELETE FROM favorites WHERE user_id = $1 AND movie_id = $2`
	_, err := r.db.Exec(context.Background(), query, favorite.UserId, favorite.MovieId)
	if err != nil {
		r.log.Printf("error in DeleteFavoriteMovie(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) GetFavoriteMovie(favorite entity.Favorite) (entity.Favorite, error) {
	query := `SELECT *  FROM favorites WHERE user_id = $1 AND movie_id = $2`
	err := r.db.QueryRow(context.Background(), query, favorite.UserId, favorite.UserId).Scan(&favorite.Id, &favorite.UserId, &favorite.UserId)
	if err != nil {
		r.log.Printf("error in DeleteFavoriteMovie(repository):%s", err.Error())
	}
	return favorite, err
}
