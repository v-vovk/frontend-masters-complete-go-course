package store

import (
	"database/sql"
	"testing"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost port=5432 user=postgres password=postgres dbname=postgres port=5533 sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	err = Migrate(db, "../../migrations/")
	if err != nil {
		t.Fatalf("Failed to migrate test database: %v", err)
	}

	_, err = db.Exec(`TRUNCATE TABLE workouts, workout_entries CASCADE`)
	if err != nil {
		t.Fatalf("Failed to truncate test database: %v", err)
	}

	return db
}
