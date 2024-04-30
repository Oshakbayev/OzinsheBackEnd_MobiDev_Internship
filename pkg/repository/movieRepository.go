package repository

import (
	"context"
	"fmt"
	"github.com/lib/pq"
	"ozinshe/pkg/entity"
)

type MovieRepo interface {
	CreateMovie(*entity.Movie) (int, error)
	CreateMoviePoster(int, string) error
	CreateMovieScreenshots(int, []string) error
	CreateMovieMain(*entity.MovieMain) error
	GetMoviesByPage(limit, offset int) ([]entity.Movie, error)
	GetMovieById(int) (*entity.Movie, error)
	UpdateMovieById(*entity.Movie) error
	DeleteMovieById(int) error
	DeleteMovieGenresByMovieID(int) error
	GetMovieSeason(movieID, seasonId int) ([]string, error)
	GetMovieSeries(movieID, seriesId, seasonId int) (string, error)
	GetMovieMainsByTitle(string) ([]entity.MovieMain, error)
}

func (r *RepoStruct) CreateMovie(movie *entity.Movie) (int, error) {
	query := `INSERT INTO movie (created_date, description, director, keywords, last_modified_date, movie_type, name, producer, season_count, series_count, timing, trend, watch_count, year,poster_link,video_directory_link) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14,$15,$16) RETURNING id`
	var insertedID int
	// Execute the query
	err := r.db.QueryRow(context.Background(), query, movie.CreatedDate, movie.Description, movie.Director, movie.Keywords, movie.LastModifiedDate, movie.MovieType, movie.Name, movie.Producer, movie.SeasonCount, movie.SeriesCount, movie.Timing, movie.Trend, movie.WatchCount, movie.Year, movie.PosterLink, movie.VideoDirectoryLink).Scan(&insertedID)
	if err != nil {
		r.log.Printf("error in CreateMovie(repository):%s", err.Error())
	}
	return insertedID, err
}

func (r *RepoStruct) CreateMoviePoster(movieID int, posterLink string) error {
	query := `INSERT INTO poster (movie_id, link)  VALUES ($1, $2) `
	_, err := r.db.Exec(context.Background(), query, movieID, posterLink)
	if err != nil {
		r.log.Printf("error in CreateMoviePoster(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) CreateMovieScreenshots(movieID int, screenshots []string) error {
	screenshotsArray := pq.Array(screenshots)
	query := `INSERT INTO screenshot (movie_id, link)  VALUES ( $1, unnest($2::text[])) `
	_, err := r.db.Exec(context.Background(), query, movieID, screenshotsArray)
	if err != nil {
		r.log.Printf("error in CreateMovieScreenshots(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) CreateMovieMain(movieMain *entity.MovieMain) error {
	query := `INSERT INTO movie_main (movie_id, movie_name, poster_link)  VALUES ($1, $2,$3) `
	_, err := r.db.Exec(context.Background(), query, movieMain.MovieId, movieMain.MovieName, movieMain.PosterLink)
	if err != nil {
		r.log.Printf("error in CreateMovieMain(repository): %s", err.Error())
	}
	return err
}

func (r *RepoStruct) GetMoviesByPage(limit, offset int) ([]entity.Movie, error) {
	query := `SELECT 
    t1.*,
    COALESCE(categories, '{NULL}') AS categories,
	COALESCE(category_ids, '{NULL}') AS category_ids,
    COALESCE(genres, '{NULL}') AS genres,
	COALESCE(genre_ids, '{NULL}') AS genre_ids,
	COALESCE(screenshot_link, '{NULL}') AS screenshot_link,
	COALESCE(video_link, '{NULL}') AS video_link,
	COALESCE(season_num,'{NULL}') AS season_num,
	COALESCE(series_num,'{NULL}') AS series_num
FROM (
    SELECT 
        movie.*,
		COALESCE(ARRAY_AGG(categoryage_id::int),'{NULL}') AS category_age_id,
        COALESCE(ARRAY_AGG(c.name),'{NULL}') AS category_age
    FROM 
        movie
    LEFT JOIN 
        movie_categoryage AS ca ON ca.movie_id = movie.id
	LEFT JOIN category_age AS c ON c.id = ca.categoryage_id
    GROUP BY 
        movie.id
) t1 
LEFT JOIN (
    SELECT 
        movie_id, 
        ARRAY_AGG(ct.name) AS categories,
		ARRAY_AGG(category_id::int) AS category_ids
    FROM 
        movie_category AS mc
	LEFT JOIN category AS ct ON ct.id = mc.category_id
    GROUP BY 
        movie_id
) t2 ON t1.id = t2.movie_id
LEFT JOIN (
    SELECT 
        movie_id, 
        ARRAY_AGG(g.name) AS genres,
		ARRAY_AGG(genre_id::int) AS genre_ids
    FROM 
        movie_genre AS mg
	LEFT JOIN genre AS g ON g.id = mg.genre_id
    GROUP BY 
        movie_id
) t3 ON t1.id = t3.movie_id
LEFT JOIN (
    SELECT 
        movie_id, 
        ARRAY_AGG(link) AS screenshot_link
    FROM 
        screenshot
    GROUP BY 
        movie_id
) t4 ON t1.id = t4.movie_id
LEFT JOIN (
    SELECT 
        movie_id, 
        ARRAY_AGG(link) AS video_link,
		 ARRAY_AGG(season_num) AS season_num ,
		ARRAY_AGG(series_num) AS series_num
    FROM 
        video
    GROUP BY 
        movie_id
) t5 ON t1.id = t5.movie_id
ORDER BY 
    t1.id
LIMIT $1 OFFSET $2`
	// Execute the query
	var movies []entity.Movie
	rows, err := r.db.Query(context.Background(), query, limit, offset)
	if err != nil {
		r.log.Printf("error in GetAllMovies(repository):%s", err.Error())
		return nil, err
	}

	for rows.Next() {
		fmt.Println("-----------------")
		fmt.Println(rows.Values())
		fmt.Println("-----------------")
		movie := entity.Movie{}
		var screenshotLinks, videoLinks, categories, categoryAges, genres []string
		var categoryAgeIds, categoryIds, genreIds, videoSeason, videoSeries []int
		err := rows.Scan(&movie.Id, &movie.CreatedDate, &movie.Description, &movie.Director, &movie.Keywords, &movie.LastModifiedDate, &movie.MovieType, &movie.Name, &movie.Producer, &movie.SeasonCount, &movie.SeriesCount, &movie.Timing, &movie.Trend, &movie.WatchCount, &movie.Year, &movie.PosterLink, &categoryAgeIds, &categoryAges, &categories, &categoryIds, &genres, &genreIds, &screenshotLinks, &videoLinks, &videoSeason, &videoSeries)
		if err != nil {
			r.log.Printf("error in GetAllMovies(repository):%s", err.Error())
			return nil, err
		}
		movie.CategoryIDs = categoryIds
		for i, val := range categories {
			category := entity.Category{}
			category.Name = val
			category.Id = categoryIds[i]
			movie.Categories = append(movie.Categories, category)
		}

		for i, val := range categoryAges {
			categoryAge := entity.CategoryAge{}
			categoryAge.Name = val
			categoryAge.Id = categoryAgeIds[i]
			movie.CategoryAges = append(movie.CategoryAges, categoryAge)
		}
		for i, val := range genres {
			genre := entity.Genre{}
			genre.Name = val
			genre.Id = genreIds[i]
			movie.Genres = append(movie.Genres, genre)
		}
		movie.CategoryAgeIDs = categoryAgeIds
		movie.GenreIDs = genreIds
		movie.ScreenshotLinks = screenshotLinks
		for i, val := range videoLinks {
			video := entity.Video{}
			video.Link = val
			video.SeasonId = videoSeason[i]
			video.SeriesNumber = videoSeries[i]
			movie.Videos = append(movie.Videos, video)
		}
		movies = append(movies, movie)
	}
	return movies, err
}

func (r *RepoStruct) GetMovieById(id int) (*entity.Movie, error) {
	query := `SELECT 
    t1.*,
    COALESCE(categories, '{NULL}') AS categories,
	COALESCE(category_ids, '{NULL}') AS category_ids,
    COALESCE(genres, '{NULL}') AS genres,
	COALESCE(genre_ids, '{NULL}') AS genre_ids,
	COALESCE(screenshot_link, '{NULL}') AS screenshot_link,
	COALESCE(video_link, '{NULL}') AS video_link,
	COALESCE(season_num,'{NULL}') AS season_num,
	COALESCE(series_num,'{NULL}') AS series_num
FROM (
    SELECT 
        movie.*,
		COALESCE(ARRAY_AGG(categoryage_id::int),'{NULL}') AS category_age_id,
        COALESCE(ARRAY_AGG(c.name),'{NULL}') AS category_age
    FROM 
        movie
    LEFT JOIN 
        movie_categoryage AS ca ON ca.movie_id = movie.id
	LEFT JOIN category_age AS c ON c.id = ca.categoryage_id
    GROUP BY 
        movie.id
) t1 
LEFT JOIN (
    SELECT 
        movie_id, 
        ARRAY_AGG(ct.name) AS categories,
		ARRAY_AGG(category_id::int) AS category_ids
    FROM 
        movie_category AS mc
	LEFT JOIN category AS ct ON ct.id = mc.category_id
    GROUP BY 
        movie_id
) t2 ON t1.id = t2.movie_id
LEFT JOIN (
    SELECT 
        movie_id, 
        ARRAY_AGG(g.name) AS genres,
		ARRAY_AGG(genre_id::int) AS genre_ids
    FROM 
        movie_genre AS mg
	LEFT JOIN genre AS g ON g.id = mg.genre_id
    GROUP BY 
        movie_id
) t3 ON t1.id = t3.movie_id
LEFT JOIN (
    SELECT 
        movie_id, 
        ARRAY_AGG(link) AS screenshot_link
    FROM 
        screenshot
    GROUP BY 
        movie_id
) t4 ON t1.id = t4.movie_id
LEFT JOIN (
    SELECT 
        movie_id, 
        ARRAY_AGG(link) AS video_link,
		 ARRAY_AGG(season_num) AS season_num ,
		ARRAY_AGG(series_num) AS series_num
    FROM 
        video
    GROUP BY 
        movie_id
) t5 ON t1.id = t5.movie_id
Where t1.id = $1
ORDER BY 
    t1.id
`
	movie := entity.Movie{}
	var screenshotLinks, videoLinks, categories, categoryAges, genres []string
	var categoryAgeIds, categoryIds, genreIds, videoSeason, videoSeries []int
	err := r.db.QueryRow(context.Background(), query, id).Scan(&movie.Id, &movie.CreatedDate, &movie.Description, &movie.Director, &movie.Keywords, &movie.LastModifiedDate, &movie.MovieType, &movie.Name, &movie.Producer, &movie.SeasonCount, &movie.SeriesCount, &movie.Timing, &movie.Trend, &movie.WatchCount, &movie.Year, &movie.PosterLink, &categoryAgeIds, &categoryAges, &categories, &categoryIds, &genres, &genreIds, &screenshotLinks, &videoLinks, &videoSeason, &videoSeries)
	if err != nil {
		r.log.Printf("error in GetAllMovies(repository):%s", err.Error())
		return nil, err
	}
	movie.CategoryIDs = categoryIds
	for i, val := range categories {
		category := entity.Category{}
		category.Name = val
		category.Id = categoryIds[i]
		movie.Categories = append(movie.Categories, category)
	}

	for i, val := range categoryAges {
		categoryAge := entity.CategoryAge{}
		categoryAge.Name = val
		categoryAge.Id = categoryAgeIds[i]
		movie.CategoryAges = append(movie.CategoryAges, categoryAge)
	}
	for i, val := range genres {
		genre := entity.Genre{}
		genre.Name = val
		genre.Id = genreIds[i]
		movie.Genres = append(movie.Genres, genre)
	}
	movie.CategoryAgeIDs = categoryAgeIds
	movie.GenreIDs = genreIds
	movie.ScreenshotLinks = screenshotLinks
	for i, val := range videoLinks {
		video := entity.Video{}
		video.Link = val
		video.SeasonId = videoSeason[i]
		video.SeriesNumber = videoSeries[i]
		movie.Videos = append(movie.Videos, video)
	}
	return &movie, err
}

func (r *RepoStruct) UpdateMovieById(movie *entity.Movie) error {
	query := `UPDATE movie 
SET 
    created_date = $1,
    description = $2,
    director = $3,
    keywords = $4,
    last_modified_date = $5,
    movie_type = $6,
    name = $7,
    producer = $8,
    season_count = $9,
    series_count = $10,
    timing = $11,
    trend = $12,
    watch_count = $13,
    year = $14,
    poster_link = $15
WHERE 
    id = $16;`
	var insertedID int
	// Execute the query
	err := r.db.QueryRow(context.Background(), query, movie.CreatedDate, movie.Description, movie.Director, movie.Keywords, movie.LastModifiedDate, movie.MovieType, movie.Name, movie.Producer, movie.SeasonCount, movie.SeriesCount, movie.Timing, movie.Trend, movie.WatchCount, movie.Year, movie.PosterLink, movie.Id).Scan(&insertedID)
	if err != nil {
		r.log.Printf("error in UpdateMovieInfoById(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) DeleteMovieById(movieID int) error {
	query := `DELETE FROM movie WHERE id = $1`
	_, err := r.db.Exec(context.Background(), query, movieID)
	if err != nil {
		r.log.Printf("error in DeleteMovieById(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) DeleteMovieGenresByMovieID(movieID int) error {

	query := `DELETE * FROM movie_genre WHERE movie_id = $1`
	_, err := r.db.Exec(context.Background(), query, movieID)
	if err != nil {
		r.log.Printf("error in DeleteMovieGenresByID(repository):%s", err.Error())
	}
	return err
}

func (r *RepoStruct) GetMovieSeason(movieID, seasonId int) ([]string, error) {
	var links []string
	query := `SELECT link FROM video WHERE movie_id = $1 AND season_num=$2`
	rows, err := r.db.Query(context.Background(), query, movieID, seasonId)
	if err != nil {
		r.log.Printf("error in GetMovieSeason(repository):%s", err.Error())
		return nil, err
	}
	for rows.Next() {
		var link string
		err = rows.Scan(&link)
		if err != nil {
			r.log.Printf("error in GetMovieSeason(repository):%s", err.Error())
			return nil, err
		}
		links = append(links, link)
	}
	return links, err

}

func (r *RepoStruct) GetMovieSeries(movieID, seriesId, seasonId int) (string, error) {
	var link string
	query := `SELECT link FROM video WHERE movie_id = $1 AND series_num = $2 AND season_num=$3`
	err := r.db.QueryRow(context.Background(), query, movieID, seriesId, seasonId).Scan(&link)
	if err != nil {
		r.log.Printf("error in GetMovieSeries(repository):%s", err.Error())
	}
	return link, err

}

func (r *RepoStruct) GetMovieMainsByTitle(title string) ([]entity.MovieMain, error) {

	title = "%" + title + "%"
	query := `SELECT 
	t1.id,
	t1.movie_id,
	t1.movie_name,
	t1.poster_link,
	t1.MovieYear,
	t1.genre_names
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
WHERE  t1.movie_name ILIKE  $1`
	return r.GetMovieMainsByQuery(query, title)
}

func (r *RepoStruct) GetMovieMainsByQuery(query string, params ...any) ([]entity.MovieMain, error) {
	var movieMains []entity.MovieMain
	rows, err := r.db.Query(context.Background(), query, params...)
	if err != nil {
		r.log.Printf("error in GetMoviesByTitle(repository):%s", err.Error())
		return nil, err
	}
	for rows.Next() {
		var movieMain entity.MovieMain
		err := rows.Scan(&movieMain.Id, &movieMain.MovieId, &movieMain.MovieName, &movieMain.PosterLink, &movieMain.MovieYear, &movieMain.MovieGenres)
		if err != nil {
			r.log.Printf("error in GetMoviesByTitle(repository):%s", err.Error())
			return nil, err
		}
		movieMains = append(movieMains, movieMain)
	}
	return movieMains, err
}
