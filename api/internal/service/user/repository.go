package user

import (
	"context"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"
)

// RepositoryUser repository interface to implemented in user repository
type RepositoryUser interface {
	Register(ctx context.Context, input model.InputCreateUser) (*ent.User, error)
	Login(ctx context.Context, input model.InputLoginUser) (*ent.User, error)
}
