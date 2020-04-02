package repository

import (
	"context"
	"time"

	"github.com/confus1on/UKM/ent"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/profile"
)

type ProfileRepository struct {
	DB *ent.Client
}

// NewProfileRepository create new repository of profile
func NewProfileRepository(DB *ent.Client) profile.RepositoryProfile {
	return &ProfileRepository{DB: DB}
}

// Update is a update new profile
func (p *ProfileRepository) Update(ctx context.Context, id int, input model.InputUpdateProfile) (*ent.Profile, error) {
	now := time.Now()

	profile, err := p.DB.Profile.UpdateOneID(id).
		SetFirstName(input.FirstName).
		SetLastName(input.LastName).
		SetYearGeneration(input.YearGeneration).
		SetPhone(input.YearGeneration).
		SetStatus(true).
		SetUpdatedAt(now).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return profile, nil
}