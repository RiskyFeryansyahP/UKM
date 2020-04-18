package repository

import (
	"context"
	"github.com/confus1on/UKM/internal/model"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"

	"github.com/confus1on/UKM/ent"

	_ "github.com/mattn/go-sqlite3" // SQLite dialect
)

func TestUKMRepository(t *testing.T) {
	ctx := context.Background()

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// initialize value ukm
	initValue(client, t)

	// initialize repository ukm
	ukm := NewRepositoryUKM(client)

	t.Run("register ukm to profile should success", func(t *testing.T) {
		input := model.InputRegisterUKM{
			Name: "SceN",
		}

		profileID := 1

		profile, err := ukm.RegisterUKM(ctx, profileID, input)

		require.NoError(t, err)
		require.Equal(t, "Risky", profile.FirstName)
		require.Equal(t, "SceN", profile.Edges.Ukm[0].Name)
	})

	t.Run("failed not found profileID", func(t *testing.T) {
		input := model.InputRegisterUKM{
			Name: "SceN",
		}

		profileID := 2

		profile, err := ukm.RegisterUKM(ctx, profileID, input)

		require.Error(t, err)
		require.Nil(t, profile)
	})

	t.Run("success get all ukm", func(t *testing.T) {
		result, err := ukm.FetchAll(ctx)

		require.Nil(t, err)
		require.Equal(t, "SceN", result[0].Name)
	})

	t.Run("success update name ukm", func(t *testing.T) {
		input := model.InputUpdateUKM{
			Name: "SFC",
		}

		ukmID := 1

		result, err := ukm.UpdateOne(ctx, ukmID, input)

		require.NoError(t, err)
		require.Equal(t, "SFC", result.Name)
	})

	t.Run("failed update name not found ukm id", func(t *testing.T) {
		input := model.InputUpdateUKM{
			Name: "Kompeni",
		}

		ukmID := 4

		result, err := ukm.UpdateOne(ctx, ukmID, input)

		require.Error(t, err)
		require.Nil(t, result)
	})
}

func initValue(client *ent.Client, t *testing.T) {
	ctx := context.Background();

	_, err := client.Ukm.Create().
		SetName("SceN").
		Save(ctx)

	require.NoError(t, err)

	user := client.User.Create().
		SetEmail("171111010@mhs.stiki.ac.id").
		SetPassword("risky").
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SaveX(ctx)

	_, err = client.Profile.Create().
		SetFirstName("Risky").
		SetLastName("Feryansyah").
		SetJk("Male").
		SetYearGeneration("2017").
		SetPhone("083834121715").
		AddOwner(user).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	require.NoError(t, err)
}