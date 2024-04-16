package repository

import (
	"context"
	"github.com/lib/pq"
	"ozinshe/pkg/entity"
)

type CategoryAgeRepo interface {
	CreateMovieCategoryAges(movieID int, CategoryAgesIDs []int) error
	GetAllCategoryAges() ([]entity.Category, error)
}

func (r *RepoStruct) CreateMovieCategoryAges(movieID int, CategoryAgesIDs []int) error {
	CategoryAgesIDsArray := pq.Array(CategoryAgesIDs)
	query := `INSERT INTO movie_categoryage (movie_id, categoryage_id)  VALUES ( $1, unnest($2::int[])) `
	_, err := r.db.Exec(context.Background(), query, movieID, CategoryAgesIDsArray)
	if err != nil {
		r.log.Printf("error in CreateMovieCategoryAges(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) GetAllCategoryAges() ([]entity.Category, error) {
	query := `SELECT id,name FROM category_age`
	var allCategoryAges []entity.Category
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		r.log.Printf("error in GetAllCategories(repository):%s", err.Error())
		return nil, err
	}
	for rows.Next() {
		var categoryAge entity.Category
		err := rows.Scan(&categoryAge.Id, &categoryAge.Name)
		if err != nil {
			r.log.Printf("error in GetAllCategories(repository):%s", err.Error())
			return nil, err
		}
		allCategoryAges = append(allCategoryAges, categoryAge)
	}
	return allCategoryAges, err
}
