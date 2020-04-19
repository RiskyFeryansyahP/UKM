package repository

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"

	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3" // SQLite dialect
)

func TestProfileRepository(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	t.Run("create role", func(t *testing.T) {
		_, err := client.Role.Create().
			SetValue("DEVELOPER").
			Save(context.Background())

		require.NoError(t, err)
	})

	t.Run("create new profile", func(t *testing.T) {
		now := time.Now()

		u, err := client.User.Create().
			SetEmail("171111040@mhs.stiki.ac.id").
			SetPassword("risky").
			SetCreatedAt(now).
			SetUpdatedAt(now).
			SetRoleID(1).
			Save(context.Background())

		require.NoError(t, err)

		profile, err := client.Profile.Create().
			SetFirstName("Risky").
			SetLastName("Pribadi").
			SetYearGeneration("2017").
			SetJk("Male").
			SetStatus(true).
			SetPhone("083834121712").
			SetImgProfile("").
			AddOwner(u).
			SetCreatedAt(now).
			SetUpdatedAt(now).
			Save(context.Background())

		require.NoError(t, err)
		require.Equal(t, "Risky", profile.FirstName)
		require.Equal(t, "Pribadi", profile.LastName)
	})

	t.Run("update profile should successfull", func(t *testing.T) {
		p := NewProfileRepository(client)

		input := model.InputUpdateProfile{
			FirstName:      "Galih",
			LastName:       "Satriawan",
			Jk:             "Male",
			Address:         "Jln Tidar",
			BirthDate:   "12-07-1997",
			YearGeneration: "2016",
			Phone:          "089691952594",
			Status:         true,
			ImgProfile:     "",
		}

		email := "171111040@mhs.stiki.ac.id"

		profile, err := p.UpdateOne(context.Background(), email, input)

		require.NoError(t, err)
		require.Equal(t, "Galih", profile.FirstName)
	})

	t.Run("update profile should be failed empty email", func(t *testing.T) {
		p := NewProfileRepository(client)

		input := model.InputUpdateProfile{
			FirstName:      "Galih",
			LastName:       "Satriawan",
			YearGeneration: "2017",
			Phone:          "089691952594",
			Status:         true,
			ImgProfile:     "",
		}

		email := ""

		_, err := p.UpdateOne(context.Background(), email, input)

		require.Error(t, err)
	})

	t.Run("update profile should be failed empty phone", func(t *testing.T) {
		p := NewProfileRepository(client)

		input := model.InputUpdateProfile{
			FirstName:      "Galih",
			LastName:       "Satriawan",
			YearGeneration: "2017",
			Phone:          "",
			Status:         true,
			ImgProfile:     "",
		}

		email := "171111040@mhs.stiki.ac.id"

		profile, err := p.UpdateOne(context.Background(), email, input)

		log.Println(profile)

		require.Error(t, err)
	})

	t.Run("get detail profile should success", func(t *testing.T) {
		p := NewProfileRepository(client)

		email := "171111040@mhs.stiki.ac.id"

		profile, err := p.FindByEmail(context.Background(), email)

		require.NoError(t, err)
		require.Equal(t, "Galih", profile.FirstName)
		require.Equal(t, "Satriawan", profile.LastName)
		require.Equal(t, "2016", profile.YearGeneration)
	})

	t.Run("get detail profile should be failed", func(t *testing.T) {
		p := NewProfileRepository(client)

		email := "171111041@mhs.stiki.ac.id"

		profile, err := p.FindByEmail(context.Background(), email)

		require.Error(t, err)
		require.Nil(t, profile)
	})
}
