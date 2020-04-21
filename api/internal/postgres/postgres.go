package postgres

import (
	"context"
	"github.com/confus1on/UKM/ent/ukm"
	"log"
	"os"
	"time"
	sqlDB "database/sql"

	"github.com/pkg/errors"

	"github.com/confus1on/UKM/ent"

	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/lib/pq" // Dialect Postgres
)

// Check DB
func checkDB(databaseURL string, db *sqlDB.DB) error {
	databaseNAME := os.Getenv("DATABASE_NAME")

	if databaseNAME == "" {
		return nil
	}

	statement := `SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = '` + databaseNAME + `');`
	row := db.QueryRow(statement)
	var exists bool
	err := row.Scan(&exists)

	if !exists {
		statement = `CREATE DATABASE ` + databaseNAME + `;`
		_, err = db.Exec(statement)
	}

	if err != nil {
		return err
	}

	_ = db.Close()

	return nil
}

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

	err = checkDB(databaseURL, db)
	if err != nil {
		log.Fatalf("error when checking db %v", err)
	}

	return ent.NewClient(ent.Driver(drv)), nil
}

// InitDefaultValue create defualt value to db for role and ukm
func InitDefaultValue(client *ent.Client) error {
	err := initValueRoleDB(client)
	if err != nil {
		err = errors.Wrap(err, "Failed init value role :")
		return err
	}

	err = initValueForUkmDB(client)
	if err != nil {
		err = errors.Wrap(err, "Failed init value ukm :")
		return err
	}

	return nil
}

func initValueRoleDB(client *ent.Client) error {
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

func initValueForUkmDB(client *ent.Client) error {
	ctx := context.Background()

	now := time.Now()

	var status ukm.Status = "close" // initial enum value status ukm

	ukms, err := client.Ukm.Query().
		All(ctx)

	if err != nil {
		return err
	}

	if len(ukms) <= 0 {
		ukmInitValue := []string{"HIC", "Byte", "SceN", "SFC", "IMAKRISKA", "Ikatan Studi Islam", "Kompeni", "Kepak Elang"}

		for _, v := range ukmInitValue {
			_, err := client.Ukm.Create().
				SetName(v).
				SetStatus(status).
				SetUpdatedAt(now).
				SetCreatedAt(now).
				Save(ctx)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
