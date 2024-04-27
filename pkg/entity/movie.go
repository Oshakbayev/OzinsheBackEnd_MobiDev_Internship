package entity

import "time"

type MovieMain struct {
	Id          int
	MovieId     int
	MovieName   string
	MovieGenres []string
	MovieYear   string
	PosterLink  string
}

type Movie struct {
	Id               int
	CategoryIDs      []int         `json:"categoryIDs"`
	Categories       []Category    `json:"categories"`
	CategoryAgeIDs   []int         `json:"categoryAgeIDs"`
	CategoryAges     []CategoryAge `json:"categoryAges"`
	CreatedDate      time.Time
	Description      string  `json:"description"`
	Director         string  `json:"director"`
	Favorite         bool    `json:"favorite"`
	GenreIDs         []int   `json:"genreIDs"`
	Genres           []Genre `json:"genres"`
	Keywords         string  `json:"keywords"`
	LastModifiedDate time.Time
	MovieType        int    `json:"movieType"`
	Name             string `json:"name"`
	PosterLink       string
	Producer         string `json:"producer"`
	ScreenshotLinks  []string
	Screenshots      []Screenshot
	SeasonCount      int  `json:"seasonCount"`
	SeriesCount      int  `json:"seriesCount"`
	Timing           int  `json:"timing"`
	Trend            bool `json:"trend"`
	Videos           []Video
	VideoLinks       []string
	WatchCount       int
	Year             int `json:"year"`
}

type Favorite struct {
	Id      int
	UserId  int
	MovieId int `json:"movieId"`
}
