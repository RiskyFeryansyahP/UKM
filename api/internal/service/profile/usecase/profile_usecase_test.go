package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/profile/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProfileUsecase_GetProfile(t *testing.T) {
	controller := gomock.NewController(t)

	t.Run("get detail profile should success", func(t *testing.T) {
		email := "171111040@mhs.stiki.ac.id"

		expect := &ent.Profile{
			ID:             1,
			FirstName:      "Galih",
			LastName:       "Satriawan",
			Jk:             "Male",
			YearGeneration: "2016",
			Phone:          "089691952594",
			Status:         true,
			ImgProfile:     "",
		}

		repo := mock.NewMockRepositoryProfile(controller)
		repo.EXPECT().FindByEmail(context.Background(), email).Return(expect, nil).Times(1)

		p := NewProfileUsecase(repo)
		response, err := p.GetDetailProfile(context.Background(), email)

		require.Nil(t, err)
		require.Equal(t, 200, response.StatusCode)
		require.Equal(t, "Galih", response.Result.FirstName)
		require.Equal(t, "Satriawan", response.Result.LastName)
	})

	t.Run("failed get detail profile empty email", func(t *testing.T) {
		email := ""

		err := errors.New("email can't be empty")

		repo := mock.NewMockRepositoryProfile(controller)
		repo.EXPECT().FindByEmail(context.Background(), email).Return(nil, err).Times(1)

		p := NewProfileUsecase(repo)
		response, respErr := p.GetDetailProfile(context.Background(), email)

		require.Equal(t, 400, respErr.StatusCode)
		require.Nil(t, response)
	})

	t.Run("failed get detail profile not found email", func(t *testing.T) {
		email := "171111041@mhs.stiki.ac.id"

		err := errors.New("profile not found")

		repo := mock.NewMockRepositoryProfile(controller)
		repo.EXPECT().FindByEmail(context.Background(), email).Return(nil, err).Times(1)

		p := NewProfileUsecase(repo)
		response, respErr := p.GetDetailProfile(context.Background(), email)

		require.Equal(t, 200, respErr.StatusCode)
		require.Nil(t, response)
	})
}

func TestProfileUsecase_UpdateProfile(t *testing.T) {
	controller := gomock.NewController(t)

	t.Run("update profile should success", func(t *testing.T) {
		input := model.InputUpdateProfile{
			FirstName:      "Galih",
			LastName:       "Satriawan",
			Jk:             "Male",
			YearGeneration: "2016",
			Phone:          "089691952594",
			Status:         true,
			ImgProfile:     "",
		}

		email := "171111040@mhs.stiki.ac.id"

		expect := &ent.Profile{
			ID:             1,
			FirstName:      "Galih",
			LastName:       "Satriawan",
			Jk:             "Male",
			YearGeneration: "2016",
			Phone:          "089691952594",
			Status:         true,
			ImgProfile:     "",
		}

		repo := mock.NewMockRepositoryProfile(controller)
		repo.EXPECT().UpdateOne(context.Background(), email, input).Return(expect, nil).Times(1)

		p := NewProfileUsecase(repo)
		response, err := p.UpdateProfile(context.Background(), email, input)

		require.Nil(t, err)
		require.Equal(t, 200, response.StatusCode)
		require.Equal(t, true, response.Status)
		require.Equal(t, "Galih", response.Result.FirstName)
		require.Equal(t, "Satriawan", response.Result.LastName)
	})

	t.Run("update profile should success", func(t *testing.T) {
		input := model.InputUpdateProfile{
			FirstName:      "Galih",
			LastName:       "Satriawan",
			Jk:             "Male",
			YearGeneration: "2016",
			Phone:          "089691952594",
			Status:         true,
			ImgProfile:     "",
		}

		email := "171111040@mhs.stiki.ac.id"

		expect := &ent.Profile{
			ID:             1,
			FirstName:      "Galih",
			LastName:       "Satriawan",
			Jk:             "Male",
			YearGeneration: "2016",
			Phone:          "089691952594",
			Status:         true,
			ImgProfile:     "",
		}

		repo := mock.NewMockRepositoryProfile(controller)
		repo.EXPECT().UpdateOne(context.Background(), email, input).Return(expect, nil).Times(1)

		p := NewProfileUsecase(repo)
		response, err := p.UpdateProfile(context.Background(), email, input)

		require.Nil(t, err)
		require.Equal(t, 200, response.StatusCode)
		require.Equal(t, true, response.Status)
		require.Equal(t, "Galih", response.Result.FirstName)
		require.Equal(t, "Satriawan", response.Result.LastName)
	})

	t.Run("update profile failed empty email", func(t *testing.T) {
		input := model.InputUpdateProfile{
			FirstName:      "Galih",
			LastName:       "Satriawan",
			Jk:             "Male",
			YearGeneration: "2016",
			Phone:          "089691952594",
			Status:         true,
			ImgProfile:     "",
		}

		email := ""

		err := errors.New("email can't be empty")

		repo := mock.NewMockRepositoryProfile(controller)
		repo.EXPECT().UpdateOne(context.Background(), email, input).Return(nil, err).Times(1)

		p := NewProfileUsecase(repo)
		response, respErr := p.UpdateProfile(context.Background(), email, input)

		require.Equal(t, 200, respErr.StatusCode)
		require.Nil(t, response)
	})

	t.Run("update profile failed empty first name", func(t *testing.T) {
		input := model.InputUpdateProfile{
			FirstName:      "",
			LastName:       "Satriawan",
			YearGeneration: "2016",
			Phone:          "089691952594",
			Status:         true,
			ImgProfile:     "",
		}

		email := "171111040@mhs.stiki.ac.id"

		err := errors.New("first name can't be empty")

		repo := mock.NewMockRepositoryProfile(controller)
		repo.EXPECT().UpdateOne(context.Background(), email, input).Return(nil, err).Times(1)

		p := NewProfileUsecase(repo)
		response, respErr := p.UpdateProfile(context.Background(), email, input)

		require.Equal(t, 400, respErr.StatusCode)
		require.Nil(t, response)
	})

	t.Run("update profile failed empty last name", func(t *testing.T) {
		input := model.InputUpdateProfile{
			FirstName:      "Galih",
			LastName:       "",
			YearGeneration: "2016",
			Phone:          "089691952594",
			Status:         true,
			ImgProfile:     "",
		}

		email := "171111040@mhs.stiki.ac.id"

		err := errors.New("last name can't be empty")

		repo := mock.NewMockRepositoryProfile(controller)
		repo.EXPECT().UpdateOne(context.Background(), email, input).Return(nil, err).Times(1)

		p := NewProfileUsecase(repo)
		response, respErr := p.UpdateProfile(context.Background(), email, input)

		require.Equal(t, 400, respErr.StatusCode)
		require.Nil(t, response)
	})

	t.Run("update profile failed empty phone", func(t *testing.T) {
		input := model.InputUpdateProfile{
			FirstName:      "Galih",
			LastName:       "Satriawan",
			Jk:             "Male",
			YearGeneration: "2016",
			Phone:          "08",
			Status:         true,
			ImgProfile:     "",
		}

		email := "171111040@mhs.stiki.ac.id"

		err := errors.New("wrong format phone number")

		repo := mock.NewMockRepositoryProfile(controller)
		repo.EXPECT().UpdateOne(context.Background(), email, input).Return(nil, err).Times(1)

		p := NewProfileUsecase(repo)
		response, respErr := p.UpdateProfile(context.Background(), email, input)

		require.Equal(t, 400, respErr.StatusCode)
		require.Nil(t, response)
	})

	t.Run("update profile failed empty JK", func(t *testing.T) {
		input := model.InputUpdateProfile{
			FirstName:      "Galih",
			LastName:       "Satriawan",
			Jk:             "",
			YearGeneration: "2016",
			Phone:          "089691952594",
			Status:         true,
			ImgProfile:     "",
		}

		email := "171111040@mhs.stiki.ac.id"

		err := errors.New("jenis kelamin can't be empty")

		repo := mock.NewMockRepositoryProfile(controller)
		repo.EXPECT().UpdateOne(context.Background(), email, input).Return(nil, err).Times(1)

		p := NewProfileUsecase(repo)
		response, respErr := p.UpdateProfile(context.Background(), email, input)

		require.Equal(t, 400, respErr.StatusCode)
		require.Nil(t, response)
	})
}
