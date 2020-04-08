package postgres

import (
	"context"
	"log"
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

// InitValueRoleDB initial value table role into database
func InitValueRoleDB(client *ent.Client) error {
	ctx := context.Background()

	roles, err := client.Role.Query().
		All(ctx)

	if err != nil {
		log.Printf("Error query roles: %v", err)
	}

	if len(roles) <= 0 {
		roleValue := []string{"DEVELOPER", "KEMAHASISWAAN", "BEM", "PENGURUS", "ANGGOTA"}

		for _, value := range roleValue {
			_, err := client.Role.Create().
				SetValue(value).
				Save(ctx)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
