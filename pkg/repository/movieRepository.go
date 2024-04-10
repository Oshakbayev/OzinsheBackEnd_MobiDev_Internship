package repository

import (
	"context"
	"github.com/lib/pq"
	"ozinshe/pkg/entity"
)

type MovieRepoInterface interface {
}

func (r *RepoStruct) CreateMovie(movie *entity.Movie) error {
	query := `INSERT INTO movie (created_date, description, director, keywords, last_modified_date, movie_type, name, producer, season_count, series_count, timing, trend, watch_count, year) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	// Execute the query
	_, err := r.db.Exec(context.Background(), query, movie.CreatedDate, movie.Description, movie.Director, movie.Keywords, movie.LastModifiedDate, movie.MovieType, movie.Name, movie.Producer, movie.SeasonCount, movie.SeriesCount, movie.Timing, movie.Trend, movie.WatchCount, movie.Year)
	if err != nil {
		r.log.Printf("error in CreateMovie(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) CreateMovieCategories(movieID int, categoryIDs []int) error {
	categoryIDsArray := pq.Array(categoryIDs)
	query := `INSERT INTO movie_category (movie_id, category_id)  VALUES ( $1, unnest($2::int[])) `
	_, err := r.db.Exec(context.Background(), query, movieID, categoryIDsArray)
	if err != nil {
		r.log.Printf("error in CreateMovieCategory(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) CreateMovieCategoryAges(movieID int, CategoryAgesIDs []int) error {
	CategoryAgesIDsArray := pq.Array(CategoryAgesIDs)
	query := `INSERT INTO movie_categoryage (movie_id, categoryage_id)  VALUES ( $1, unnest($2::int[])) `
	_, err := r.db.Exec(context.Background(), query, movieID, CategoryAgesIDsArray)
	if err != nil {
		r.log.Printf("error in CreateMovieCategory(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) CreateMovieGenres(movieID int, genreIDs []int) error {
	genreIDsArray := pq.Array(genreIDs)
	query := `INSERT INTO movie_genre (movie_id, genre_id)  VALUES ( $1, unnest($2::int[])) `
	_, err := r.db.Exec(context.Background(), query, movieID, genreIDsArray)
	if err != nil {
		r.log.Printf("error in CreateMovieCategory(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) CreateMoviePoster(movieID int, posterLink string) error {
	query := `INSERT INTO poster (movie_id, link)  VALUES ( $1, $2) `
	_, err := r.db.Exec(context.Background(), query, movieID, posterLink)
	if err != nil {
		r.log.Printf("error in CreateMoviePoster(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) CreateMovieScreenshots(movieID int, screenshots []string) error {
	screenshotsArray := pq.Array(screenshots)
	query := `INSERT INTO screenshot (movie_id, link)  VALUES ( $1, unnest($2::int[])) `
	_, err := r.db.Exec(context.Background(), query, movieID, screenshotsArray)
	if err != nil {
		r.log.Printf("error in CreateMovieCategory(repository):%s", err.Error())
	}
	return err
}
