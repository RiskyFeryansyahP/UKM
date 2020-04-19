package postgres

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"

	"github.com/confus1on/UKM/ent"

	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/lib/pq" // Dialect Postgres
)

// Check DB
func checkDB() {
	databaseURL := os.Getenv("DATABASE_URL") // load url database from environment variable
	databaseNAME := os.Getenv("DATABASE_NAME")
	databasePARAMS := os.Getenv("DATABASE_PARAMS")
	drv, err := sql.Open("postgres", databaseURL+"?"+databasePARAMS)
	if err != nil {
		log.Fatalf("Error Connect DB %v", err)
	}

	db := drv.DB()

	statement := `SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = '` + databaseNAME + `');`
	row := db.QueryRow(statement)
	var exists bool
	err = row.Scan(&exists)

	if exists == false {
		statement = `CREATE DATABASE ` + databaseNAME + `;`
		_, err = db.Exec(statement)
	}
	if err != nil {
		log.Fatalf("Error Create DB %v", err)
	}
	db.Close()
}

// NewPostgreSQL create a client connection to database
func NewPostgreSQL() (*ent.Client, error) {
	checkDB()
	databaseURL := os.Getenv("DATABASE_URL") // load url database from environment variable
	databaseNAME := os.Getenv("DATABASE_NAME")
	databasePARAMS := os.Getenv("DATABASE_PARAMS")
	drv, err := sql.Open("postgres", databaseURL+databaseNAME+"?"+databasePARAMS)
	if err != nil {
		return nil, err
	}

	db := drv.DB()
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)

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
