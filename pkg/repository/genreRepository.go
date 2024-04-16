package repository

import (
	"context"
	"github.com/lib/pq"
	"ozinshe/pkg/entity"
)

type GenreRepo interface {
	CreateMovieGenres(movieID int, genreIDs []int) error
	GetAllGenres() ([]entity.Category, error)
}

func (r *RepoStruct) CreateMovieGenres(movieID int, genreIDs []int) error {
	genreIDsArray := pq.Array(genreIDs)
	query := `INSERT INTO movie_genre (movie_id, genre_id)  VALUES ( $1, unnest($2::int[])) `
	_, err := r.db.Exec(context.Background(), query, movieID, genreIDsArray)
	if err != nil {
		r.log.Printf("error in CreateMovieGenres(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) GetAllGenres() ([]entity.Category, error) {
	query := `SELECT id,name FROM genre`
	var allGenres []entity.Category
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		r.log.Printf("error in GetAllCategories(repository):%s", err.Error())
		return nil, err
	}
	for rows.Next() {
		var genre entity.Category
		err := rows.Scan(&genre.Id, &genre.Name)
		if err != nil {
			r.log.Printf("error in GetAllCategories(repository):%s", err.Error())
			return nil, err
		}
		allGenres = append(allGenres, genre)
	}
	return allGenres, err
}
