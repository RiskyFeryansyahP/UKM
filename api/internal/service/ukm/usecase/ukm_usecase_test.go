package usecase

import (
	"context"
	"errors"
	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/ukm/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUKMUsecase_GetAll(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()


	mockRepo := mock.NewMockRepositoryUKM(controller)

	t.Run("success get all ukm", func(t *testing.T) {
		mockRepo.EXPECT().FetchAll(ctx).Return([]*ent.Ukm{}, nil).Times(1)

		ukmUC := NewUKMUsecase(mockRepo)
		response, err := ukmUC.GetAll(ctx)

		require.Nil(t, err)
		require.Equal(t, 200, response.StatusCode)
		require.Equal(t, true, response.Status)
	})

	t.Run("failed get all ukm connection refused", func(t *testing.T) {
		respErr := errors.New("connection refused")

		mockRepo.EXPECT().FetchAll(ctx).Return(nil, respErr).Times(1)

		ukmUC := NewUKMUsecase(mockRepo)
		response, err := ukmUC.GetAll(ctx)

		require.Equal(t, 200, err.StatusCode)
		require.Equal(t, "failed get all ukm: : connection refused", err.Message)
		require.Nil(t, response)
	})
}

func TestUKMUsecase_Register(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()

	mockRepo := mock.NewMockRepositoryUKM(controller)

	t.Run("success register ukm to profile", func(t *testing.T) {
		input := model.InputRegisterUKM{
			Name: "SceN",
		}

		profileID := 1

		mockRepo.EXPECT().RegisterUKM(ctx, profileID, input).Return(&ent.Profile{}, nil).Times(1)

		ukmUC := NewUKMUsecase(mockRepo)
		response, err := ukmUC.RegisterUKM(ctx, profileID, input)

		require.Nil(t, err)
		require.Equal(t, 200, response.StatusCode)
		require.Equal(t, true, response.Status)
	})

	t.Run("failed register ukm to profile not found profile", func(t *testing.T) {
		input := model.InputRegisterUKM{
			Name: "SceN",
		}

		profileID := 5

		respErr := errors.New("not found profile match with profile id")

		mockRepo.EXPECT().RegisterUKM(ctx, profileID, input).Return(nil, respErr).Times(1)

		ukmUC := NewUKMUsecase(mockRepo)
		response, err := ukmUC.RegisterUKM(ctx, profileID, input)

		require.Equal(t, 200, err.StatusCode)
		require.Equal(t, "failed register ukm: : not found profile match with profile id", err.Message)
		require.Nil(t, response)
	})

	t.Run("failed register ukm to profile empty profileid", func(t *testing.T) {
		input := model.InputRegisterUKM{
			Name: "SceN",
		}

		profileID := 0

		respErr := errors.New("phone can't be empty")

		mockRepo.EXPECT().RegisterUKM(ctx, profileID, input).Return(nil, respErr).Times(1)

		ukmUC := NewUKMUsecase(mockRepo)
		response, err := ukmUC.RegisterUKM(ctx, profileID, input)

		require.Equal(t, 400, err.StatusCode)
		require.Equal(t, "profile can't be empty", err.Message)
		require.Nil(t, response)
	})

	t.Run("failed register ukm to profile empty name", func(t *testing.T) {
		input := model.InputRegisterUKM{
			Name: "",
		}

		profileID := 1

		respErr := errors.New("name can't be empty")

		mockRepo.EXPECT().RegisterUKM(ctx, profileID, input).Return(nil, respErr).Times(1)

		ukmUC := NewUKMUsecase(mockRepo)
		response, err := ukmUC.RegisterUKM(ctx, profileID, input)

		require.Equal(t, 400, err.StatusCode)
		require.Equal(t, "name can't be empty", err.Message)
		require.Nil(t, response)
	})
}

func TestUKMUsecase_Update(t *testing.T) {
	controller := gomock.NewController(t)

	ctx := context.Background()

	mockRepo := mock.NewMockRepositoryUKM(controller)

	t.Run("success update ukm", func(t *testing.T) {
		input := model.InputUpdateUKM{
			Name: "Kompeni",
		}

		ukmID := 1

		mockRepo.EXPECT().UpdateOne(ctx, ukmID, input).Return(&ent.Ukm{}, nil).Times(1)

		ukmUC := NewUKMUsecase(mockRepo)
		response, err := ukmUC.UpdateUKM(ctx, ukmID, input)

		require.Nil(t, err)
		require.Equal(t, 200, response.StatusCode)
		require.Equal(t, true, response.Status)
	})

	t.Run("failed update ukm wrong ukm id", func(t *testing.T) {
		input := model.InputUpdateUKM{
			Name: "Kompeni",
		}

		ukmID := 100

		respErr := errors.New("not found id match in database")

		mockRepo.EXPECT().UpdateOne(ctx, ukmID, input).Return(nil, respErr).Times(1)

		ukmUC := NewUKMUsecase(mockRepo)
		response, err := ukmUC.UpdateUKM(ctx, ukmID, input)

		require.Equal(t, 200, err.StatusCode)
		require.Equal(t, "failed update ukm: : not found id match in database", err.Message)
		require.Nil(t, response)
	})

	t.Run("failed update ukm empty name", func(t *testing.T) {
		input := model.InputUpdateUKM{
			Name: "",
		}

		ukmID := 1

		respErr := errors.New("name can't be empty")

		mockRepo.EXPECT().UpdateOne(ctx, ukmID, input).Return(nil, respErr).Times(1)

		ukmUC := NewUKMUsecase(mockRepo)
		response, err := ukmUC.UpdateUKM(ctx, ukmID, input)

		require.Equal(t, 400, err.StatusCode)
		require.Equal(t, "name can't be empty", err.Message)
		require.Nil(t, response)
	})

	t.Run("failed update ukm empty ukmID", func(t *testing.T) {
		input := model.InputUpdateUKM{
			Name: "SceN",
		}

		ukmID := 0

		respErr := errors.New("id can't be empty")

		mockRepo.EXPECT().UpdateOne(ctx, nil, input).Return(nil, respErr).Times(1)

		ukmUC := NewUKMUsecase(mockRepo)
		response, err := ukmUC.UpdateUKM(ctx, ukmID, input)

		require.Equal(t, 400, err.StatusCode)
		require.Equal(t, "id can't be empty", err.Message)
		require.Nil(t, response)
	})
}