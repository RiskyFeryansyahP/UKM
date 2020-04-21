package repository

import (
	"context"
	"errors"

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

// FetchAll fetch all ukm data in database
func (u *UKMRepository) FetchAll(ctx context.Context) ([]*ent.Ukm, error) {
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
	rowUpdated := u.DB.Ukm.Update().
		Where(
			ukmField.And(
				ukmField.NameEQ(input.Name),
				ukmField.StatusEQ("open"),
			),
		).
		AddProfileIDs(profileID).
		SaveX(ctx)

	// failed register ukm, probably status close
	if rowUpdated <= 0 {
		err := errors.New("registration ukm is closed")
		return nil, err
	}

	profile := u.DB.Profile.Query().
		Where(profile.IDEQ(profileID)).
		WithUkm().
		OnlyX(ctx)

	return profile, nil
}

// UpdateOne set a new name for ukm
func (u *UKMRepository) UpdateOne(ctx context.Context, id int, input model.InputUpdateUKM) (*ent.Ukm, error) {
	ukm, err := u.DB.Ukm.UpdateOneID(id).
		SetName(input.Name).
		SetStatus(input.Status).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return ukm, nil
}
