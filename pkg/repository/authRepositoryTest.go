package repository

//
//import (
//	"context"
//	"database/sql"
//	"github.com/jackc/pgx/v5/pgxpool"
//	"log"
//	"ozinshe/pkg/entity"
//	"ozinshe/pkg/repository/database"
//	"testing"
//)
//
//type DB interface {
//	QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row
//}
//type MockPool struct {
//	// Define any fields needed for mocking.
//	Result *pgxpool.Pool
//	Err    error
//}
//
//// QueryRow mocks the QueryRow method of the DB interface.
//func (m *MockPool) QueryRow(ctx context.Context, query string, args ...interface{}) *pgxpool.Pool {
//	return m.Result
//}
//func TestCreateUser(t *testing.T) {
//	// Mock data
//	mockUser := &entity.User{
//		Email:    "test@example.com",
//		Password: "password",
//		Id:       1, // Assuming Id is initialized to 1
//	}
//	db, err := database.ConnectToDB(config.DSN, logger)
//	if err != nil {
//		logger.Println(err)
//		log.Fatal(err)
//	}
//	defer db.Close()
//	// Mock pool
//	mockPool := &MockPool{ // Create mock result if needed,
//		Err: nil, // No error
//	}
//
//	// Create repository with mock pool
//	repo := CreateRepository()
//
//	// Test CreateUser function
//	err := repo.CreateUser(mockUser)
//	if err != nil {
//		t.Errorf("Unexpected error: %v", err)
//	}
//
//	// Check if the user ID is set after insertion
//	if mockUser.Id != 1 {
//		t.Errorf("Expected user ID to be 1, got %d", mockUser.Id)
//	}
//
//	// Add more test cases as needed
//}
