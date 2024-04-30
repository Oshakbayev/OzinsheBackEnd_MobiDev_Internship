package repository

import (
	"context"
	"ozinshe/pkg/entity"
)

type FavoriteMovieRepo interface {
	CreatFavoriteMovie(entity.Favorite) error
	GetUserFavoriteMovieMains(int) ([]entity.MovieMain, error)
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

func (r *RepoStruct) GetUserFavoriteMovieMains(userId int) ([]entity.MovieMain, error) {
	query := `SELECT 
	t1.id,
	t1.movie_id,
	t1.movie_name,
	t1.poster_link,
	t1.year,
	t1.genre_names,
	t1.is_favorite
FROM(	
	SELECT 
		mm.*,
		ARRAY_AGG(g.name) AS genre_names,
		ARRAY_AGG(g.id) AS genre_ids,
		CASE WHEN f.movie_id IS NOT NULL THEN true ELSE false END AS is_favorite
	FROM 
		movie_genre AS mg
	LEFT JOIN 
		movie_main AS mm ON mm.movie_id = mg.movie_id
	LEFT JOIN 
		genre AS g ON g.id = mg.genre_id	
	LEFT JOIN (
    	SELECT movie_id
    	FROM favorites
    	WHERE user_id = $1
	) f ON mm.movie_id = f.movie_id
	GROUP BY 
		mm.id, f.movie_id
) AS t1
WHERE 
	t1.is_favorite = true;
`
	return r.GetMovieMainsByQuery(query, userId)
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
