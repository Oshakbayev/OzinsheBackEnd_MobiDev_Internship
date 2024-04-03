package database

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func Config(DATABASE_URL string, errLog *log.Logger) *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	// Your own Database URL
	//const DATABASE_URL string = "postgres://postgres:12345678@localhost:5432/postgres?"

	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		errLog.Printf("Failed to create a config, error: %s", err.Error())
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		errLog.Println("Before acquiring the connection pool to the database!!")
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		errLog.Println("After releasing the connection pool to the database!!")
		log.Println("After releasing the connection pool to the database!!")
		return true
	}

	dbConfig.BeforeClose = func(c *pgx.Conn) {
		errLog.Println("Closed the connection pool to the database!!")
		log.Println("Closed the connection pool to the database!!")
	}

	return dbConfig
}

func ConnectToDB(DATABASE_URL string, errlog *log.Logger) (*pgxpool.Pool, error) {
	connPool, err := pgxpool.NewWithConfig(context.Background(), Config(DATABASE_URL, errlog))
	if err != nil {
		return nil, errors.New("Error while creating connection to the database!!")
	}

	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		return nil, errors.New("Error while creating connection to the database!!")
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		return nil, errors.New("Could not ping database")
	}

	errlog.Println("Connected to the database!!")
	return connPool, nil
}
