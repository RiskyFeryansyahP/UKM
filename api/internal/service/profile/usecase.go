package profile

import (
	"context"
	"github.com/confus1on/UKM/internal/model"
)

// UsecaseProfile is a interface that will be implemented in profile usecase
type UsecaseProfile interface {
	UpdateProfile(ctx context.Context, id int, input model.InputUpdateProfile) (*model.ResponseUpdateProfile, error)
}
