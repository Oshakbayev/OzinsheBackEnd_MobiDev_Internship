package repository

import (
	"context"
	"github.com/lib/pq"
	"ozinshe/pkg/entity"
)

type GenreRepo interface {
	CreateMovieGenres(int, []int) error
	GetAllGenres() ([]entity.Genre, error)
	GetMovieMainsByGenre(int, int) ([]entity.MovieMain, error)
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

func (r *RepoStruct) GetAllGenres() ([]entity.Genre, error) {
	query := `SELECT id,name FROM genre`
	var allGenres []entity.Genre
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		r.log.Printf("error in GetAllCategories(repository):%s", err.Error())
		return nil, err
	}
	for rows.Next() {
		var genre entity.Genre
		err := rows.Scan(&genre.Id, &genre.Name)
		if err != nil {
			r.log.Printf("error in GetAllCategories(repository):%s", err.Error())
			return nil, err
		}
		allGenres = append(allGenres, genre)
	}
	return allGenres, err
}

func (r *RepoStruct) GetMovieMainsByGenre(userId, genreId int) ([]entity.MovieMain, error) {
	query := `SELECT 
	t1.id,
	t1.movie_id,
	t1.movie_name,
	t1.poster_link,
	t1.year,
	t1.genre_names,
	CASE WHEN f.movie_id IS NOT NULL THEN true ELSE false END AS is_favorite
FROM(	
SELECT 
		mm.*,
		ARRAY_AGG(g.name) AS genre_names,
		ARRAY_AGG(g.id) AS genre_ids
 	FROM 
		movie_genre as mg
	LEFT JOIN 
		movie_main as mm on mm.movie_id =mg.movie_id
	LEFT JOIN genre as g on g.id = mg.genre_id	
	GROUP BY 
		mm.id
) AS t1
LEFT JOIN (
    SELECT movie_id
    FROM favorites
    WHERE user_id = $1
) f ON t1.id = f.movie_id
WHERE $2 = ANY(t1.genre_ids)`
	return r.GetMovieMainsByQuery(query, userId, genreId)
}
