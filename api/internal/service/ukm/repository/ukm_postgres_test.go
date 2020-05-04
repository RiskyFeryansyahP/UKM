package repository

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"

	_ "github.com/mattn/go-sqlite3" // SQLite dialect
	"github.com/stretchr/testify/require"
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
			Reason: "i like scen",
			Name: "SceN",
		}

		profileID := 1

		profile, err := ukm.RegisterUKM(ctx, profileID, input)

		require.NoError(t, err)
		require.Equal(t, "Risky", profile.FirstName)
		require.Equal(t, "i like scen", profile.Edges.Ukms[0].Reason)
	})

	t.Run("failed registration ukm is closed", func(t *testing.T) {
		input := model.InputRegisterUKM{
			Reason: "Love sf",
			Name: "SFC",
		}

		profileID := 1

		profile, err := ukm.RegisterUKM(ctx, profileID, input)

		require.Error(t, err)
		require.Nil(t, profile)
	})

	t.Run("failed registration ukm is closed", func(t *testing.T) {
		input := model.InputRegisterUKM{
			Reason: "Love scen",
			Name: "SceN",
		}

		profileID := 1

		profile, err := ukm.RegisterUKM(ctx, profileID, input)

		t.Log("err", err)

		require.Error(t, err)
		require.Nil(t, profile)
	})

	t.Run("failed registration reason is empty", func(t *testing.T) {
		input := model.InputRegisterUKM{
			Reason: "",
			Name: "Pencak Silat",
		}

		profileID := 1

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
			Name:   "SFC",
			Status: "open",
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
	ctx := context.Background()

	_, err := client.Ukm.Create().
		SetName("SceN").
		SetStatus("open").
		Save(ctx)

	require.NoError(t, err)

	_, err = client.Ukm.Create().
		SetName("SFC").
		SetStatus("close").
		Save(ctx)

	require.NoError(t, err)

	_, err = client.Ukm.Create().
		SetName("Pencak Silat").
		SetStatus("open").
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

	_, err = client.RoleUKM.Create().
		SetStatusRole("PENGURUS").
		Save(ctx)

	require.NoError(t, err)

	_, err = client.RoleUKM.Create().
		SetStatusRole("ANGGOTA").
		Save(ctx)

	require.NoError(t, err)
}
