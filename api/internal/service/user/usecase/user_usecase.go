package usecase

import (
	"context"

	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/user"
)

// UserUsecase struct to inject depedency from repository user
type UserUsecase struct {
	UserRepo user.RepositoryUser
}

// NewUserUsecase is a construct to initialize UserUsecase
func NewUserUsecase(userRepo user.RepositoryUser) user.UsecaseUser {
	return &UserUsecase{UserRepo: userRepo}
}

// CreateUser create user with return a model response register and error
func (u *UserUsecase) CreateUser(ctx context.Context, input model.InputCreateUser) (*model.ResponseRegister, error) {
	user, err := u.UserRepo.Register(ctx, input)
	if err != nil {
		return nil, err
	}

	response := &model.ResponseRegister{
		StatusCode: 200,
		Status:     true,
		Result:     user,
	}

	return response, nil
}

// SigninUser validation input and return to model ResponseLogin
func (u *UserUsecase) SigninUser(ctx context.Context, input model.InputLoginUser) (*model.ResponseLogin, error) {
	user, err := u.UserRepo.Login(ctx, input)
	if err != nil {
		return nil, err
	}

	response := &model.ResponseLogin{
		StatusCode: 200,
		Status:     true,
		Result:    user,
	}

	return response, nil
}
