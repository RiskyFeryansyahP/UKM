package user

import (
	"context"
	"github.com/confus1on/UKM/internal/model"
)

// UsecaseUser is inteface to implemented in usecase user
type UsecaseUser interface {
	CreateUser(ctx context.Context, input model.InputCreateUser) (*model.ResponseRegister, error)
	SigninUser(ctx context.Context, input model.InputLoginUser) (*model.ResponseLogin, error)
}
