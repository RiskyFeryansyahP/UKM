package repository

import (
	"context"
	"errors"
	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/ent/profile"
	"github.com/confus1on/UKM/ent/profileukm"
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
	ukmID, err := u.DB.Ukm.Query().
		Where(
			ukmField.And(
				ukmField.Name(input.Name),
				ukmField.StatusEQ("open"),
			),
		).
		OnlyID(ctx)
	if err != nil {
		err = errors.New("ukm already closed")
		return nil, err
	}

	count := u.DB.ProfileUKM.Query().
		Where(
			profileukm.And(
				profileukm.HasOwnerUkmWith(ukmField.ID(ukmID)),
				profileukm.HasOwnerProfileWith(profile.ID(profileID)),
				),
		).
		CountX(ctx)
	if count > 0 {
		err = errors.New("already registered in this ukm")
		return nil, err
	}

	_, err = u.DB.ProfileUKM.Create().
		SetReason(input.Reason).
		SetOwnerProfileID(profileID).
		SetOwnerUkmID(ukmID).
		SetOwnerRoleID(2).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	profile := u.DB.Profile.Query().
		Where(profile.IDEQ(profileID)).
		WithUkms().
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
