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

// Register doing validation data and converting response
func (u *UKMUsecase) Register(ctx context.Context, input model.InputRegisterUKM) (*model.ResponseRegisterUKM, *utils.Error) {
	if input.Name == "" || &input.Name == nil {
		err := errors.New("name can't be empty")

		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	if input.ProfileID == 0 || &input.ProfileID == nil {
		err := errors.New("phone can't be empty")

		return nil, utils.WrapErrorJson(err, fasthttp.StatusBadRequest)
	}

	profile, err := u.UKMRepo.RegisterUKM(ctx, input)
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
