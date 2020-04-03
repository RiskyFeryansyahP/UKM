package profile

import (
	"context"

	"github.com/confus1on/UKM/internal/model"
)

// UsecaseProfile is a interface that will be implemented in profile usecase
type UsecaseProfile interface {
	GetProfile(ctx context.Context, email string) (*model.ResponseGetProfile, error)
	UpdateProfile(ctx context.Context, email string, input model.InputUpdateProfile) (*model.ResponseUpdateProfile, error)
}
