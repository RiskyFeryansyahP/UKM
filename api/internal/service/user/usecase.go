package user

import (
	"github.com/confus1on/UKM/internal/model"

	"github.com/valyala/fasthttp"
)

// UsecaseUser is inteface to implemented in usecase user
type UsecaseUser interface {
	CreateUser(ctx *fasthttp.RequestCtx, input model.InputCreateUser) (*model.ResponseRegister, error)
}
