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
