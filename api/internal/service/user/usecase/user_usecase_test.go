package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/user/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserUsecase_CreateUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("convert return to model.register and should successfully", func(t *testing.T) {
		now := time.Now()

		profile := &ent.Profile{
			FirstName:      "Risky F",
			LastName:       "Pribadi",
			Jk:             "Male",
			YearGeneration: "2017",
			Phone:          "083834121712",
			Status:         true,
			ImgProfile:     "",
			CreatedAt:      now,
			UpdatedAt:      now,
		}

		input := model.InputCreateUser{
			Email:       "171111040@mhs.stiki.ac.id",
			Password:    "risky",
			UserProfile: profile,
		}

		u := mock.NewMockRepositoryUser(controller)
		u.EXPECT().Register(context.Background(), input).Return(&ent.User{}, nil).Times(1)

		user := NewUserUsecase(u)
		res, err := user.CreateUser(context.Background(), input)

		require.Nil(t, err)
		require.Equal(t, 200, res.StatusCode)
		require.Equal(t, true, res.Status)
	})

	t.Run("failed conver return because nil data", func(t *testing.T) {
		input := model.InputCreateUser{
			Email:       "",
			Password:    "",
			UserProfile: nil,
		}

		err := errors.New("error invalid data")

		u := mock.NewMockRepositoryUser(controller)
		u.EXPECT().Register(context.Background(), input).Return(nil, err).Times(1)

		user := NewUserUsecase(u)
		_, err = user.CreateUser(context.Background(), input)

		require.Error(t, err)
	})
}

func TestUserUsecase_SigninUser(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("should success convert model data", func(t *testing.T) {
		input := model.InputLoginUser{
			Email:    "171111040@mhs.stiki.ac.id",
			Password: "risky",
		}

		u := mock.NewMockRepositoryUser(controller)
		u.EXPECT().Login(context.Background(), input).Return(&ent.User{}, nil).AnyTimes()

		user := NewUserUsecase(u)
		p, err := user.SigninUser(context.Background(), input)

		require.NoError(t, err)
		require.Equal(t, 200, p.StatusCode)
		require.Equal(t, true, p.Status)
	})

	t.Run("nil data profile and failed", func(t *testing.T) {
		input := model.InputLoginUser{
			Email:    "",
			Password: "",
		}

		err := errors.New("error data nil")

		u := mock.NewMockRepositoryUser(controller)
		u.EXPECT().Login(context.Background(), input).Return(nil, err).Times(1)

		user := NewUserUsecase(u)
		_, err = user.SigninUser(context.Background(), input)
		require.Error(t, err)
	})
}
