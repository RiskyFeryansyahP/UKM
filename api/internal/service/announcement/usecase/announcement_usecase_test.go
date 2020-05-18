package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/announcement/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestAnnouncementUsecase_ValidationPostAnnouncement(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()

	announcementRepoMock := mock.NewMockRepositoryAnnouncement(controller)

	t.Run("test validation posting should success", func(t *testing.T) {
		// dump data
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 1",
			Description: "Dimohon para anggota hadir dalam rapat ini",
			Time:        "20/05/2020 18:00:00",
		}

		ukmID := 1 // dump data

		announcementRepoMock.EXPECT().Posting(ctx, ukmID, input).Return(&ent.Announcement{}, nil).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)
		result, err := announcement.ValidationPostAnnouncement(ctx, ukmID, input)
		require.Nil(t, err)
		require.Equal(t, 200, result.StatusCode)
		require.Equal(t, true, result.Status)
	})

	t.Run("test validation posting failed foreign key not matched in ukm", func(t *testing.T) {
		// dump data
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 1",
			Description: "Dimohon para anggota hadir dalam rapat ini",
			Time:        "20/05/2020 18:00:00",
		}

		ukmID := 100 // error not matched with ukm value in table

		resultErr := errors.New("foreign key error not matched with each ukm id")

		announcementRepoMock.EXPECT().Posting(ctx, ukmID, input).Return(nil, resultErr).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)
		result, err := announcement.ValidationPostAnnouncement(ctx, ukmID, input)
		require.NotNil(t, err)
		require.Equal(t, 400, err.StatusCode)
		require.Nil(t, result)
	})

	t.Run("test validation posting failed zero ukm id", func(t *testing.T) {
		// dump data
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 1",
			Description: "Dimohon para anggota hadir dalam rapat ini",
			Time:        "20/05/2020 18:00:00",
		}

		ukmID := 0 // error cause zero id

		var resultErr error

		announcementRepoMock.EXPECT().Posting(ctx, ukmID, input).Return(nil, resultErr).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)
		result, err := announcement.ValidationPostAnnouncement(ctx, ukmID, input)
		require.NotNil(t, err)
		require.Nil(t, result)
	})

	t.Run("test validation posting failed empty title", func(t *testing.T) {
		// dump data
		input := model.InputAnnouncement{
			Title:       "",
			Description: "Dimohon para anggota hadir dalam rapat ini",
			Time:        "20/05/2020 18:00:00",
		}

		ukmID := 1 // dump data

		var resultErr error

		announcementRepoMock.EXPECT().Posting(ctx, ukmID, input).Return(nil, resultErr).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)
		result, err := announcement.ValidationPostAnnouncement(ctx, ukmID, input)
		require.NotNil(t, err)
		require.Nil(t, result)
	})

	t.Run("test validation posting failed empty description", func(t *testing.T) {
		// dump data
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 1",
			Description: "",
			Time:        "20/05/2020 18:00:00",
		}

		ukmID := 1 // dump data

		var resultErr error

		announcementRepoMock.EXPECT().Posting(ctx, ukmID, input).Return(nil, resultErr).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)
		result, err := announcement.ValidationPostAnnouncement(ctx, ukmID, input)
		require.NotNil(t, err)
		require.Nil(t, result)
	})

	t.Run("test validation posting failed empty time", func(t *testing.T) {
		// dump data
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 1",
			Description: "Dimohon para anggota hadir dalam rapat ini",
			Time:        "",
		}

		ukmID := 1 // dump data

		var resultErr error

		announcementRepoMock.EXPECT().Posting(ctx, ukmID, input).Return(nil, resultErr).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)
		result, err := announcement.ValidationPostAnnouncement(ctx, ukmID, input)
		require.NotNil(t, err)
		require.Nil(t, result)
	})
}

func TestAnnouncementUsecase_GetSpecificAnnouncement(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()

	announcementRepoMock := mock.NewMockRepositoryAnnouncement(controller)

	t.Run("test get specific announcement should success", func(t *testing.T) {
		ukmID := 1

		announcementRepoMock.EXPECT().GetUkmAnnouncement(ctx, ukmID).Return([]*ent.Announcement{}, nil).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)

		result, err := announcement.GetSpecificAnnouncement(ctx, ukmID)
		require.Nil(t, err)
		require.Equal(t, 200, result.StatusCode)
		require.Equal(t, true, result.Status)
	})

	t.Run("test get specific announcement failed zero ukm id", func(t *testing.T) {
		ukmID := 0 // error cause zero ukm id invalid validation

		resultErr := errors.New("ukm id can't be null or zero")

		announcementRepoMock.EXPECT().GetUkmAnnouncement(ctx, ukmID).Return(nil, resultErr).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)

		result, err := announcement.GetSpecificAnnouncement(ctx, ukmID)
		require.Nil(t, result)
		require.NotNil(t, err)
		require.Equal(t, 400, err.StatusCode)
	})

	t.Run("test get specific announcement failed request time out", func(t *testing.T) {
		ukmID := 1

		resultErr := errors.New("request time out")

		announcementRepoMock.EXPECT().GetUkmAnnouncement(ctx, ukmID).Return(nil, resultErr).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)

		result, err := announcement.GetSpecificAnnouncement(ctx, ukmID)
		require.Nil(t, result)
		require.NotNil(t, err)
		require.Equal(t, 400, err.StatusCode)
	})
}

func TestAnnouncementUsecase_UpdateOneAnnouncement(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()

	announcementRepoMock := mock.NewMockRepositoryAnnouncement(controller)

	t.Run("test update one announcement should success", func(t *testing.T) {
		// dump data
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 1",
			Description: "tidak wajib hadir dalam rapat ini",
			Time:        "20/05/2020 19:00:00",
		}

		// dump data
		announcementID := 1

		announcementRepoMock.EXPECT().Update(ctx, announcementID, input).Return(&ent.Announcement{}, nil).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)

		result, err := announcement.UpdateOneAnnouncement(ctx, announcementID, input)
		require.Nil(t, err)
		require.Equal(t, 200, result.StatusCode)
		require.Equal(t, true, result.Status)
	})

	t.Run("test update one announcement failed not found announcement id database", func(t *testing.T) {
		// dump data
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 1",
			Description: "tidak wajib hadir dalam rapat ini",
			Time:        "20/05/2020 19:00:00",
		}

		// dump data
		announcementID := 100 // error cause zero id is invalid validation

		resultErr := errors.New("not found announcement id matched")

		announcementRepoMock.EXPECT().Update(ctx, announcementID, input).Return(nil, resultErr).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)

		result, err := announcement.UpdateOneAnnouncement(ctx, announcementID, input)
		require.Nil(t, result)
		require.NotNil(t, err)
		require.Equal(t, 400, err.StatusCode)
	})

	t.Run("test update one announcement failed zero announcement id", func(t *testing.T) {
		// dump data
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 1",
			Description: "tidak wajib hadir dalam rapat ini",
			Time:        "20/05/2020 19:00:00",
		}

		// dump data
		announcementID := 0 // error cause zero id is invalid validation

		resultErr := errors.New("announcement id can't be null or zero")

		announcementRepoMock.EXPECT().Update(ctx, announcementID, input).Return(nil, resultErr).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)

		result, err := announcement.UpdateOneAnnouncement(ctx, announcementID, input)
		require.Nil(t, result)
		require.NotNil(t, err)
		require.Equal(t, 400, err.StatusCode)
	})

	t.Run("test update one announcement failed empty field title", func(t *testing.T) {
		// dump data
		input := model.InputAnnouncement{
			Title:       "",
			Description: "tidak wajib hadir dalam rapat ini",
			Time:        "20/05/2020 19:00:00",
		}

		// dump data
		announcementID := 1

		resultErr := errors.New("title can't be null or empty")

		announcementRepoMock.EXPECT().Update(ctx, announcementID, input).Return(nil, resultErr).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)

		result, err := announcement.UpdateOneAnnouncement(ctx, announcementID, input)
		require.Nil(t, result)
		require.NotNil(t, err)
		require.Equal(t, 400, err.StatusCode)
	})

	t.Run("test update one announcement failed empty field description", func(t *testing.T) {
		// dump data
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 1",
			Description: "",
			Time:        "20/05/2020 19:00:00",
		}

		// dump data
		announcementID := 1

		resultErr := errors.New("description can't be null or empty")

		announcementRepoMock.EXPECT().Update(ctx, announcementID, input).Return(nil, resultErr).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)

		result, err := announcement.UpdateOneAnnouncement(ctx, announcementID, input)
		require.Nil(t, result)
		require.NotNil(t, err)
		require.Equal(t, 400, err.StatusCode)
	})

	t.Run("test update one announcement failed empty field time", func(t *testing.T) {
		// dump data
		input := model.InputAnnouncement{
			Title:       "Rapat Diklat 1",
			Description: "tidak wajib hadir dalam rapat ini",
			Time:        "",
		}

		// dump data
		announcementID := 1

		resultErr := errors.New("time can't be null or emptry")

		announcementRepoMock.EXPECT().Update(ctx, announcementID, input).Return(nil, resultErr).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)

		result, err := announcement.UpdateOneAnnouncement(ctx, announcementID, input)
		require.Nil(t, result)
		require.NotNil(t, err)
		require.Equal(t, 400, err.StatusCode)
	})
}

func TestAnnouncementUsecase_DeleteOneAnnouncement(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()

	announcementRepoMock := mock.NewMockRepositoryAnnouncement(controller)

	t.Run("test delete one announcement should success", func(t *testing.T) {
		announcementID := 1

		announcementRepoMock.EXPECT().Delete(ctx, announcementID).Return(nil).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)

		result, err := announcement.DeleteOneAnnouncement(ctx, announcementID)
		require.Nil(t, err)
		require.Equal(t, 200, result.StatusCode)
		require.Equal(t, true, result.Status)
	})

	t.Run("test delete one announcement failed not found announcement id", func(t *testing.T) {
		announcementID := 100

		resultErr := errors.New("not found announcement id matched")

		announcementRepoMock.EXPECT().Delete(ctx, announcementID).Return(resultErr).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)

		result, err := announcement.DeleteOneAnnouncement(ctx, announcementID)
		require.Nil(t, result)
		require.NotNil(t, err)
		require.Equal(t, 400, err.StatusCode)
	})

	t.Run("test delete one announcement failed zero announcement id", func(t *testing.T) {
		announcementID := 0

		resultErr := errors.New("announcement id can't be null or zero")

		announcementRepoMock.EXPECT().Delete(ctx, announcementID).Return(resultErr).Times(1)

		announcement := NewAnnouncementUsecase(announcementRepoMock)

		result, err := announcement.DeleteOneAnnouncement(ctx, announcementID)
		require.Nil(t, result)
		require.NotNil(t, err)
		require.Equal(t, 400, err.StatusCode)
	})
}
