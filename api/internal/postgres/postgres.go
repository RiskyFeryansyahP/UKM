package postgres

import (
	"os"
	"time"

	"github.com/confus1on/UKM/ent"

	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/lib/pq" // Dialect Postgres
)

// NewPostgreSQL create a client connection to database
func NewPostgreSQL() (*ent.Client, error) {
	databaseURL := os.Getenv("DATABASE_URL") // load url database from environment variable

	drv, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	db := drv.DB()
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	return ent.NewClient(ent.Driver(drv)), nil
}
