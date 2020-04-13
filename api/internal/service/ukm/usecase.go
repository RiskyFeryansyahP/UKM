package ukm

import (
	"context"

	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/utils"
)

// UsecaseUKM abstract to be implemented in ukm usecase
type UsecaseUKM interface {
	Register(ctx context.Context, input model.InputRegisterUKM) (*model.ResponseRegisterUKM, *utils.Error)
}
