package repository

import (
	"context"
	"time"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/ent/user"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/profile"
)

// ProfileRepository repository of profile to connect database
type ProfileRepository struct {
	DB *ent.Client
}

// NewProfileRepository create new repository of profile
func NewProfileRepository(DB *ent.Client) profile.RepositoryProfile {
	return &ProfileRepository{DB: DB}
}

// Update is a update new profile
func (p *ProfileRepository) UpdateOne(ctx context.Context, email string, input model.InputUpdateProfile) (*ent.Profile, error) {
	now := time.Now()

	userID, err := p.DB.User.Query().
		Where(user.Email(email)).
		QueryProfile().
		OnlyID(ctx)

	if err != nil {
		return nil, err
	}

	profile, err := p.DB.Profile.UpdateOneID(userID).
		SetFirstName(input.FirstName).
		SetLastName(input.LastName).
		SetJk(input.Jk).
		SetAddress(input.Address).
		SetBirthDate(input.BirthDate).
		SetYearGeneration(input.YearGeneration).
		SetPhone(input.Phone).
		SetStatus(true).
		SetUpdatedAt(now).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

// GetByEmail get detail profile by email
func (p *ProfileRepository) FindByEmail(ctx context.Context, email string) (*ent.Profile, error) {
	profile, err := p.DB.User.Query().
		Where(user.EmailEQ(email)).
		QueryProfile().
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
