package repository

import (
	"context"
	"github.com/lib/pq"
	"ozinshe/pkg/entity"
)

type CategoryRepo interface {
	CreateMovieCategories(movieID int, categoryIDs []int) error
	GetCategoryIdByName(categoryName string) (int, error)
	GetAllCategories() ([]entity.Category, error)
	GetAllMovieMainsByCategory(int, int) ([]entity.MovieMain, error)
	GetMovieMainsByCategory(int, int, int, int) ([]entity.MovieMain, error)
}

func (r *RepoStruct) CreateMovieCategories(movieID int, categoryIDs []int) error {
	categoryIDsArray := pq.Array(categoryIDs)
	query := `INSERT INTO movie_category (movie_id, category_id)  VALUES ( $1, unnest($2::int[])) `
	_, err := r.db.Exec(context.Background(), query, movieID, categoryIDsArray)
	if err != nil {
		r.log.Printf("error in CreateMovieCategories(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) GetCategoryIdByName(categoryName string) (int, error) {
	query := `SELECT id FROM category WHERE  name = $1`
	var categoryId int
	err := r.db.QueryRow(context.Background(), query, categoryName).Scan(categoryId)
	if err != nil {
		r.log.Printf("error in GetCategoryIdByName(repository):%s", err.Error())
	}
	return categoryId, err
}

func (r *RepoStruct) GetAllCategories() ([]entity.Category, error) {
	query := `SELECT id,name FROM category`
	var allCategories []entity.Category
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		r.log.Printf("error in GetAllCategories(repository):%s", err.Error())
		return nil, err
	}
	for rows.Next() {
		var category entity.Category
		err := rows.Scan(&category.Id, &category.Name)
		if err != nil {
			r.log.Printf("error in GetAllCategories(repository):%s", err.Error())
			return nil, err
		}
		allCategories = append(allCategories, category)
	}
	return allCategories, err
}

func (r *RepoStruct) GetAllMovieMainsByCategory(userId int, categoryId int) ([]entity.MovieMain, error) {
	query := `SELECT 
	t1.*,
	COALESCE(genre_names, '{NULL}') AS genre_names,
	CASE WHEN f.movie_id IS NOT NULL THEN true ELSE false END AS is_favorite
FROM (
	SELECT 
		mm.*
		
 	FROM 
		movie_category as mc
	LEFT JOIN 
		movie_main as mm on mm.movie_id =mc.movie_id
	WHERE 
		mc.category_id = $1
	) as t1 
LEFT JOIN (
	SELECT mg.movie_id,
	ARRAY_AGG(g.name) AS genre_names
	FROM
		movie_genre as mg
	LEFT JOIN 
		genre AS g ON g.id = mg.genre_id
	GROUP BY
		 mg.movie_id
) as t2 on t1.movie_id = t2.movie_id
LEFT JOIN (
    SELECT movie_id
    FROM favorites
    WHERE user_id = $2
) f ON t1.id = f.movie_id`
	//rows, err := r.db.Query(context.Background(), query, categoryId)
	//if err != nil {
	//	r.log.Printf("error in GetMovieIdByGenre(repository):%s", err.Error())
	//	return nil, err
	//}
	//var movieIds []entity.MovieMain
	//var movieMains []entity.MovieMain
	//for rows.Next() {
	//	var movieMain entity.MovieMain
	//	err := rows.Scan(&movieMain.Id, &movieMain.MovieId, &movieMain.MovieName, &movieMain.PosterLink, &movieMain.MovieGenres)
	//	if err != nil {
	//		r.log.Printf("error in GetMovieMainByMovieIds(repository):%s", err.Error())
	//		return nil, err
	//	}
	//	movieMains = append(movieMains, movieMain)
	//}
	//return movieIds, err
	return r.GetMovieMainsByQuery(query, categoryId, userId)
}

func (r *RepoStruct) GetMovieMainsByCategory(userId, categoryId, limit, offset int) ([]entity.MovieMain, error) {
	query := `SELECT 
    t1.*,
    COALESCE(genre_names, '{NULL}') AS genre_names,
	CASE WHEN f.movie_id IS NOT NULL THEN true ELSE false END AS is_favorite
FROM (
	SELECT 
		mm.*
 	FROM 
		movie_category as mc
	LEFT JOIN 
		movie_main as mm on mm.movie_id =mc.movie_id
	WHERE 
		mc.category_id = $1
	) as t1 
LEFT JOIN (
	SELECT mg.movie_id,
	ARRAY_AGG(g.name) AS genre_names
	FROM
		movie_genre as mg
	LEFT JOIN 
		genre AS g ON g.id = mg.genre_id
	GROUP BY
		 mg.movie_id
) as t2 on t1.movie_id = t2.movie_id
LEFT JOIN (
    SELECT movie_id
    FROM favorites
    WHERE user_id = $2
) f ON t1.id = f.movie_id
LIMIT $3 OFFSET $4`
	//rows, err := r.db.Query(context.Background(), query, categoryId, limit, offset)
	//if err != nil {
	//	r.log.Printf("error in GetMovieIdByGenre(repository):%s", err.Error())
	//	return nil, err
	//}
	//var movieMains []entity.MovieMain
	//for rows.Next() {
	//	var movieMain entity.MovieMain
	//	err := rows.Scan(&movieMain.Id, &movieMain.MovieId, &movieMain.MovieName, &movieMain.PosterLink, &movieMain.MovieGenres)
	//	if err != nil {
	//		r.log.Printf("error in GetMovieMainByMovieIds(repository):%s", err.Error())
	//		return nil, err
	//	}
	//	movieMains = append(movieMains, movieMain)
	//}
	//return movieMains, err
	return r.GetMovieMainsByQuery(query, categoryId, userId, limit, offset)
}
