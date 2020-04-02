package usecase

import (
	"context"
	"errors"
	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/profile"
)

type ProfileUsecase struct {
	ProfileRepo profile.RepositoryProfile
}

// NewProfileUsecase create new usecase of profile
func NewProfileUsecase(profileRepo profile.RepositoryProfile) profile.UsecaseProfile {
	return &ProfileUsecase{ProfileRepo: profileRepo}
}

// UpdateProfile update new of data profile
func (p *ProfileUsecase) UpdateProfile(ctx context.Context, id int, input model.InputUpdateProfile) (*model.ResponseUpdateProfile, error) {
	if input.LastName == "" || &input.LastName == nil {
		err := errors.New("last name can't be empty")
		return nil, err
	}

	if input.FirstName == "" || &input.FirstName == nil {
		err := errors.New("first name can't be empty")
		return nil, err
	}

	if len(input.Phone) < 11 {
		err := errors.New("wrong format phone number")
		return nil, err
	}

	profile, err := p.ProfileRepo.Update(ctx, id, input)
	if err != nil {
		return nil, err
	}

	response := &model.ResponseUpdateProfile{
		StatusCode: 200,
		Status:     true,
		Result:     profile,
	}

	return response, nil
}


