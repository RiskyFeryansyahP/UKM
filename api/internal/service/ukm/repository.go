package ukm

import (
	"context"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"
)

// RepositoryUKM abstract to be implemented in repository
type RepositoryUKM interface {
	RegisterUKM(ctx context.Context, input model.InputRegisterUKM) (*ent.Profile, error)
}
