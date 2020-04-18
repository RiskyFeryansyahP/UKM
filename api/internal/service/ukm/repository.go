package ukm

import (
	"context"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"
)

// RepositoryUKM abstract to be implemented in repository
type RepositoryUKM interface {
	FetchAll(ctx context.Context) ([]*ent.Ukm, error)
	RegisterUKM(ctx context.Context, profileID int, input model.InputRegisterUKM) (*ent.Profile, error)
	UpdateOne(ctx context.Context, id int, input model.InputUpdateUKM) (*ent.Ukm, error)
}
