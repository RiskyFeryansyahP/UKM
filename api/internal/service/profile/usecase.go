package profile

import (
	"context"
	"github.com/confus1on/UKM/internal/utils"

	"github.com/confus1on/UKM/internal/model"
)

// UsecaseProfile is a interface that will be implemented in profile usecase
type UsecaseProfile interface {
	GetDetailProfile(ctx context.Context, email string) (*model.ResponseGetProfile, *utils.Error)
	UpdateProfile(ctx context.Context, email string, input model.InputUpdateProfile) (*model.ResponseUpdateProfile, *utils.Error)
}
