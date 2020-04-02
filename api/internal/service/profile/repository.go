package profile

import (
	"context"
	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"
)

// RepositoryProfile is a interface to be implemented in profile repository
type RepositoryProfile interface {
	Update(ctx context.Context, id int, input model.InputUpdateProfile) (*ent.Profile, error)
}