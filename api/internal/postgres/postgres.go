package postgres

import (
	"time"

	"github.com/RiskyFeryansyahP/UKM/ent"

	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/lib/pq" // Dialect Postgres
)

// NewPostgreSQL create a client connection to database
func NewPostgreSQL() (*ent.Client, error) {
	drv, err := sql.Open("postgres", "postgres://feaxpsyq:K5J6nBn1tv-GrMn7-yGh-EqB6dj9tbIY@rosie.db.elephantsql.com:5432/feaxpsyq")
	if err != nil {
		return nil, err
	}

	db := drv.DB()
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

	return ent.NewClient(ent.Driver(drv)), nil
}
