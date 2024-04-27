package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type RepoInterface interface {
	UserRepo
	VerificationEmailRepo
	MovieRepo
	CategoryRepo
	CategoryAgeRepo
	GenreRepo
	UserProfileRepo
	FavoriteMovieRepo
	VideoRepo
}

type RepoStruct struct {
	log *log.Logger
	db  *pgxpool.Pool
}

func CreateRepository(db *pgxpool.Pool, log *log.Logger) RepoInterface {
	return &RepoStruct{db: db, log: log}
}
