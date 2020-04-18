package repository

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"

	_ "github.com/mattn/go-sqlite3" // SQLite Dialect
	"github.com/stretchr/testify/require"
)

func TestUserRepository(t *testing.T) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	t.Run("create role", func(t *testing.T) {
		_, err := client.Role.Create().
			SetValue("DEVELOPER").
			Save(context.Background())

		require.NoError(t, err)
	})

	t.Run("create user and profile", func(t *testing.T) {
		now := time.Now()

		profile := &ent.Profile{
			ID:             1,
			FirstName:      "Risky F",
			LastName:       "Pribadi",
			Jk:             "Male",
			YearGeneration: "2017",
			Phone:          "083834121715",
			Status:         true,
			ImgProfile:     "",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		user := model.InputCreateUser{
			Email:       "171111040@mhs.stiki.ac.id",
			Password:    "risky",
			Role:        1,
			UserProfile: profile,
		}

		u := NewUserRepository(client)

		_, err = u.Register(context.Background(), user)
		require.NoError(t, err)
	})

	t.Run("fail create user if not enter email / password", func(t *testing.T) {
		now := time.Now()

		profile := &ent.Profile{
			ID:             2,
			FirstName:      "Teuku R",
			LastName:       "Syahputra",
			YearGeneration: "2017",
			Phone:          "083834121713",
			Status:         true,
			ImgProfile:     "",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		input := model.InputCreateUser{
			Email:       "",
			Password:    "teuku",
			UserProfile: profile,
		}

		user := NewUserRepository(client)

		u, err := user.Register(context.Background(), input)
		require.Error(t, err)
		require.Nil(t, u)
	})

	t.Run("fail create profile if data not valid", func(t *testing.T) {
		now := time.Now()

		profile := &ent.Profile{
			ID:             2,
			FirstName:      "Bryan",
			LastName:       "Pradana",
			YearGeneration: "2016",
			Phone:          "",
			Status:         true,
			ImgProfile:     "",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		input := model.InputCreateUser{
			Email:       "171111041@mhs.stiki.ac.id",
			Password:    "teuku",
			Role:        1,
			UserProfile: profile,
		}

		user := NewUserRepository(client)

		u, err := user.Register(context.Background(), input)
		require.Error(t, err)
		require.Nil(t, u)
	})

	t.Run("success login with valid email and password", func(t *testing.T) {
		input := model.InputLoginUser{
			Email:    "171111040@mhs.stiki.ac.id",
			Password: "risky",
		}

		userRepo := NewUserRepository(client)

		user, err := userRepo.Login(context.Background(), input)
		require.NoError(t, err)
		require.Equal(t, "Risky F", user.Edges.Profile.FirstName)
		require.Equal(t, "Pribadi", user.Edges.Profile.LastName)
		require.Equal(t, true, user.Edges.Profile.Status)
		require.Equal(t, 1, user.Edges.Role.ID)
	})

	t.Run("failed login wrong email / password", func(t *testing.T) {
		input := model.InputLoginUser{
			Email:    "17111101@mhs.stiki.ac.id",
			Password: "risky",
		}

		userRepo := NewUserRepository(client)

		user, err := userRepo.Login(context.Background(), input)
		require.Error(t, err)
		require.Nil(t, user)
	})
}
