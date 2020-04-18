package ukm

import (
	"context"

	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/utils"
)

// UsecaseUKM abstract to be implemented in ukm usecase
type UsecaseUKM interface {
	GetAll(ctx context.Context) (*model.ResponseGetAllUKM, *utils.Error)
	RegisterUKM(ctx context.Context, profileID int, input model.InputRegisterUKM) (*model.ResponseRegisterUKM, *utils.Error)
	UpdateUKM(ctx context.Context, id int, input model.InputUpdateUKM) (*model.ResponseUpdateUKM, *utils.Error)
}
