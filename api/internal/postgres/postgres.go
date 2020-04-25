package postgres

import (
	"context"
	sqlDB "database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/ent/ukm"

	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/lib/pq" // Dialect Postgres
	"github.com/pkg/errors"
)

// Check DB
func checkDB(db *sqlDB.DB) error {
	databaseNAME := os.Getenv("DATABASE_NAME")

	if databaseNAME == "" {
		return nil
	}

	statement := fmt.Sprintf("SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database WHERE datname = '%s')", databaseNAME)
	row := db.QueryRow(statement)

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		statement = fmt.Sprintf("CREATE DATABASE %s", databaseNAME)
		_, err = db.Exec(statement)
		if err != nil {
			return err
		}
	}

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

	err = checkDB(db)
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

	err = initValueForRoleUKM(client)
	if err != nil {
		err = errors.Wrap(err , "failed init value role ukm :")
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
		roleValue := []string{"DEVELOPER", "KEMAHASISWAAN", "BEM", "MAHASISWA"}

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

func initValueForRoleUKM(client *ent.Client) error {
	ctx := context.Background()

	roleUKMValue := []string{"PENGURUS", "ANGGOTA"}

	role := client.RoleUKM.Query().
		AllX(ctx)

	if len(role) <= 0 {
		for _, v := range roleUKMValue {
			_, err := client.RoleUKM.Create().
				SetStatusRole(v).
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
