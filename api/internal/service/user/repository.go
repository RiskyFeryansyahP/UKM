package user

import (
	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"

	"github.com/valyala/fasthttp"
)

// RepositoryUser repository interface to implemented in user repository
type RepositoryUser interface {
	Register(ctx *fasthttp.RequestCtx, input model.InputCreateUser) (*ent.User, error)
}
