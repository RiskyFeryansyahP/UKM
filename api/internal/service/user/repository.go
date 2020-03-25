package user

import (
	"github.com/RiskyFeryansyahP/UKM/ent"
	"github.com/RiskyFeryansyahP/UKM/internal/model"

	"github.com/valyala/fasthttp"
)

// RepositoryUser repository interface to implemented in user repository
type RepositoryUser interface {
	Register(ctx *fasthttp.RequestCtx, input model.InputCreateUser) (*ent.User, error)
	Login(ctx *fasthttp.RequestCtx, input model.InputLoginUser) (*ent.Profile, error)
}
