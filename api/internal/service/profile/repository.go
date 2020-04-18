package profile

import (
	"context"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"
)

// RepositoryProfile is a interface to be implemented in profile repository
type RepositoryProfile interface {
	FindByEmail(ctx context.Context, email string) (*ent.Profile, error)
	UpdateOne(ctx context.Context, email string, input model.InputUpdateProfile) (*ent.Profile, error)
}
