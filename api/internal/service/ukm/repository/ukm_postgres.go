package repository

import (
	"context"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/ent/profile"
	ukmField "github.com/confus1on/UKM/ent/ukm"
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

// GetAll fetch all ukm data in database
func (u *UKMRepository) GetAll(ctx context.Context) ([]*ent.Ukm, error) {
	ukms, err := u.DB.Ukm.Query().
		Order(ent.Asc(ukmField.FieldID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return ukms, nil
}

// RegisterUKM registering ukm for users
func (u *UKMRepository) RegisterUKM(ctx context.Context, profileID int, input model.InputRegisterUKM) (*ent.Profile, error) {
	_, err := u.DB.Ukm.Update().
		AddProfileIDs(profileID).
		Where(ukmField.NameEQ(input.Name)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	profile, err := u.DB.Profile.Query().
		Where(profile.IDEQ(profileID)).
		WithUkm().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return profile, err
}

// Update set a new name for ukm
func (u *UKMRepository) Update(ctx context.Context, id int, input model.InputUpdateUKM) (*ent.Ukm, error) {
	ukm, err := u.DB.Ukm.UpdateOneID(id).
		SetName(input.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return ukm, nil
}