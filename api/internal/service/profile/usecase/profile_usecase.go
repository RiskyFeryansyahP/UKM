package usecase

import (
	"context"
	"errors"
	"github.com/confus1on/UKM/internal/utils"
	"github.com/valyala/fasthttp"

	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/profile"
)

// ProfileUsecase usecase each profile
type ProfileUsecase struct {
	ProfileRepo profile.RepositoryProfile
}

// NewProfileUsecase create new usecase of profile
func NewProfileUsecase(profileRepo profile.RepositoryProfile) profile.UsecaseProfile {
	return &ProfileUsecase{ProfileRepo: profileRepo}
}

// GetProfile get one profile and validation email
func (p *ProfileUsecase) GetDetailProfile(ctx context.Context, email string) (*model.ResponseGetProfile, *utils.Error) {
	if email == "" || &email == nil {
		err := errors.New("email can't be empty")
		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	profile, err := p.ProfileRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, utils.WrapErrorJson(err, fasthttp.StatusOK)
	}

	response := &model.ResponseGetProfile{
		StatusCode: 200,
		Status:     true,
		Result:     profile,
	}

	return response, nil
}

// UpdateProfile update new of data profile
func (p *ProfileUsecase) UpdateProfile(ctx context.Context, email string, input model.InputUpdateProfile) (*model.ResponseUpdateProfile, *utils.Error) {
	if input.LastName == "" || &input.LastName == nil {
		err := errors.New("last name can't be empty")
		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	if input.FirstName == "" || &input.FirstName == nil {
		err := errors.New("first name can't be empty")
		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	if input.Jk == "" || &input.Jk == nil {
		err := errors.New("jenis kelamin can't be empty")
		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	if len(input.Phone) < 11 {
		err := errors.New("wrong format phone number")
		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	profile, err := p.ProfileRepo.UpdateOne(ctx, email, input)
	if err != nil {
		return nil, utils.WrapErrorJson(err, fasthttp.StatusOK)
	}

	response := &model.ResponseUpdateProfile{
		StatusCode: 200,
		Status:     true,
		Result:     profile,
	}

	return response, nil
}
