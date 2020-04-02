package profile

import (
	"context"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"
)

// RepositoryProfile is a interface to be implemented in profile repository
type RepositoryProfile interface {
	GetByEmail(ctx context.Context, email string) (*ent.Profile, error)
	Update(ctx context.Context, email string, input model.InputUpdateProfile) (*ent.Profile, error)
}
