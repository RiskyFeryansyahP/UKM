package announcement

import (
	"context"

	"github.com/confus1on/UKM/internal/model"
	"github.com/confus1on/UKM/internal/utils"
)

// UsecaseAnnouncement hold interface value for usecase
type UsecaseAnnouncement interface {
	GetSpecificAnnouncement(ctx context.Context, ukmID int) (*model.ResultAnnouncement, *utils.Error)
	ValidationPostAnnouncement(ctx context.Context, ukmID int, input model.InputAnnouncement) (*model.SingleResultAnnouncement, *utils.Error)
	UpdateOneAnnouncement(ctx context.Context, announcementID int, input model.InputAnnouncement) (*model.SingleResultAnnouncement, *utils.Error)
	DeleteOneAnnouncement(ctx context.Context, announcementID int) (*model.SingleResultAnnouncement, *utils.Error)
}
