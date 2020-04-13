package repository

import (
	"context"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/ent/profile"
	ukmPredict "github.com/confus1on/UKM/ent/ukm"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/ukm"
)

// UKMRepository hold value of client DB
type UKMRepository struct {
	DB *ent.Client
}

// NewRepositoryUKM init repository ukm
func NewRepositoryUKM(DB *ent.Client) ukm.RepositoryUKM {
	return &UKMRepository{DB: DB}
}

// RegisterUKM registering ukm for users
func (u *UKMRepository) RegisterUKM(ctx context.Context, input model.InputRegisterUKM) (*ent.Profile, error) {
	_, err := u.DB.Ukm.Update().
		AddProfileIDs(input.ProfileID).
		Where(ukmPredict.NameEQ(input.Name)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	profile, err := u.DB.Profile.Query().
		Where(profile.IDEQ(input.ProfileID)).
		WithUkm().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return profile, err
}
