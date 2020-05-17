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

func TestNewAnnouncementRepository(t *testing.T) {
	ctx := context.Background()

	// initialize connection to db sqlite for testing only
	client := initConnection(ctx)
	defer client.Close()

	announcement := NewAnnouncementRepository(client)

	t.Run("test posting announcement should success", func(t *testing.T) {
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 1",
			Description: "Dimohon para anggota hadir dalam rapat ini",
			Time:        "20/05/2020 18:00:00",
		}

		ukmIDSceN := 1 // number id get in the first create in initUkm

		result, err := announcement.Posting(ctx, ukmIDSceN, input)
		require.NoError(t, err)
		require.Equal(t, "Rapat Diklat 1", result.Title)
	})

	t.Run("test posting announcement failed not found ukm", func(t *testing.T) {
		input := model.InputAnnouncement{
			Title:       "Rapat Hasil Pemilihan Ketum",
			Description: "Rapat tidak wajib tapi diharap dateng",
			Time:        "25/06/2020 15:00:00",
		}

		ukmID := 3 // id of 3 is not found because in initUkm() only declare 2 ukm values

		result, err := announcement.Posting(ctx, ukmID, input)
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("test get all announcement should success", func(t *testing.T) {
		ukmID := 1

		results, err := announcement.GetUkmAnnouncement(ctx, ukmID)
		require.NoError(t, err)
		require.NotNil(t, results)
	})

	t.Run("test get all announcement should failed timeout", func(t *testing.T) {
		// trying to make timeout operation in database and will return error
		ctxTimeout, cancel := context.WithTimeout(ctx, 0*time.Millisecond)
		defer cancel()

		ukmID := 1

		results, err := announcement.GetUkmAnnouncement(ctxTimeout, ukmID)
		require.Error(t, err)
		require.Nil(t, results)
	})

	t.Run("test update announcement should success", func(t *testing.T) {
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 1",
			Description: "tidak wajib hadir dalam rapat ini",
			Time:        "20/05/2020 19:00:00",
		}

		announcementID := 1 // announcement id 1 get from first posting announcement in above

		result, err := announcement.Update(ctx, announcementID, input)
		require.NoError(t, err)
		require.Equal(t, "tidak wajib hadir dalam rapat ini", result.Description)
		require.Equal(t, "20/05/2020 19:00:00", result.Time)
	})

	t.Run("test update announcement failed not found announcement matched", func(t *testing.T) {
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 1",
			Description: "tidak wajib hadir dalam rapat ini",
			Time:        "20/05/2020 19:00:00",
		}

		announcementID := 2

		result, err := announcement.Update(ctx, announcementID, input)
		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("test delete announcement should success", func(t *testing.T) {
		announcementID := 1 // announcement id 1 get from first posting announcement in above

		err := announcement.Delete(ctx, announcementID)
		require.NoError(t, err)
	})

	t.Run("test delete announcement failed not found announcement id matched", func(t *testing.T) {
		announcementID := 2

		err := announcement.Delete(ctx, announcementID)
		require.Error(t, err)
	})
}

func initConnection(ctx context.Context) *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// initialize ukm value
	initUkm(client, ctx)

	return client
}

func initUkm(client *ent.Client, ctx context.Context) {
	now := time.Now()

	_, err := client.Ukm.Create().
		SetName("SceN").
		SetStatus("open").
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		log.Fatal("error creating ukm in announcement testing", err.Error())
	}

	_, err = client.Ukm.Create().
		SetName("SFC").
		SetStatus("close").
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
	if err != nil {
		log.Fatal("error creating ukm in announcement testing", err.Error())
	}
}
