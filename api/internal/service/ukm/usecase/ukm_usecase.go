package usecase

import (
	"context"

	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/service/ukm"
	"github.com/confus1on/UKM/internal/utils"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

// UKMUsecase hold value ukm repo
type UKMUsecase struct {
	UKMRepo ukm.RepositoryUKM
}

// NewUKMUsecase init ukm usecase
func NewUKMUsecase(UKMRepo ukm.RepositoryUKM) ukm.UsecaseUKM {
	return &UKMUsecase{UKMRepo: UKMRepo}
}

// GetAll converting response and custom when error
func (u *UKMUsecase) GetAll(ctx context.Context) (*model.ResponseGetAllUKM, *utils.Error) {
	ukms, err := u.UKMRepo.FetchAll(ctx)
	if err != nil {
		err := errors.Wrap(err, "failed get all ukm: ")

		return nil, utils.WrapErrorJson(err, fasthttp.StatusOK)
	}

	response := &model.ResponseGetAllUKM{
		StatusCode: 200,
		Status:     true,
		Result:     ukms,
	}

	return response, nil
}

// Register doing validation data and converting response
func (u *UKMUsecase) RegisterUKM(ctx context.Context, profileID int, input model.InputRegisterUKM) (*model.ResponseRegisterUKM, *utils.Error) {
	if input.Name == "" || &input.Name == nil {
		err := errors.New("name can't be empty")

		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	if profileID == 0 || &profileID == nil {
		err := errors.New("profile can't be empty")

		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	profile, err := u.UKMRepo.RegisterUKM(ctx, profileID, input)
	if err != nil {
		err = errors.Wrap(err, "failed register ukm: ")

		return nil, utils.WrapErrorJson(err, fasthttp.StatusOK)
	}

	response := &model.ResponseRegisterUKM{
		StatusCode: 200,
		Status:     true,
		Result:     profile,
	}

	return response, nil
}

// Update validate input and converting response
func (u *UKMUsecase) UpdateUKM(ctx context.Context, id int, input model.InputUpdateUKM) (*model.ResponseUpdateUKM, *utils.Error) {
	if input.Name == "" || &input.Name == nil {
		err := errors.New("name can't be empty")

		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	if id == 0 || &id == nil {
		err := errors.New("id can't be empty")

		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	ukm, err := u.UKMRepo.UpdateOne(ctx, id, input)
	if err != nil {
		err = errors.Wrap(err, "failed update ukm: ")

		return nil, utils.WrapErrorJson(err, fasthttp.StatusOK)
	}

	response := &model.ResponseUpdateUKM{
		StatusCode: 200,
		Status:     true,
		Result:     ukm,
	}

	return response, nil
}
