package usecase

import (
	"github.com/RiskyFeryansyahP/UKM/internal/model"
	"github.com/RiskyFeryansyahP/UKM/internal/service/user"

	"github.com/valyala/fasthttp"
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
func (u *UserUsecase) CreateUser(ctx *fasthttp.RequestCtx, input model.InputCreateUser) (*model.ResponseRegister, error) {
	user, err := u.UserRepo.Register(ctx, input)
	if err != nil {
		return nil, err
	}

	response := &model.ResponseRegister{
		StatusCode: 200,
		Status:     true,
		Result: model.ResultRegister{
			User: user,
		},
	}

	return response, nil
}

// SigninUser validation input and return to model ResponseLogin
func (u *UserUsecase) SigninUser(ctx *fasthttp.RequestCtx, input model.InputLoginUser) (*model.ResponseLogin, error) {
	p, err := u.UserRepo.Login(ctx, input)
	if err != nil {
		return nil, err
	}

	profile := &model.ResponseLogin{
		StatusCode: 200,
		Status:     true,
		Profile:    p,
	}

	return profile, nil
}
